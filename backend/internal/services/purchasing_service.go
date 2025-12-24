package services

import (
	"log"
	"net/http"

	"github.com/DavidAfdal/purchasing-systeam/internal/dto"
	"github.com/DavidAfdal/purchasing-systeam/internal/models"
	"github.com/DavidAfdal/purchasing-systeam/internal/repositories"
	"github.com/DavidAfdal/purchasing-systeam/pkg/datetime"
	"github.com/DavidAfdal/purchasing-systeam/pkg/errors"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type purchasingService struct {
	db                   *gorm.DB
	purchasingRepo       repositories.PurchasingRepo
	itemRepo             repositories.ItemRepo
	purchasingDetailRepo repositories.PurchasingDetailRepo
	webhookService       WebhookService
}

type PurchasingService interface {
	GetPurchasings() ([]*dto.PurchasingResponse, error)
	GetPurchasingByUserID(userID string) ([]*dto.PurchasingResponse, error)
	CreatePurchasing(req *dto.CreatePurchasingRequest) error
}

func NewPurchasingService(
	db *gorm.DB,
	purchasingRepo repositories.PurchasingRepo,
	itemRepo repositories.ItemRepo,
	purchasingDetailRepo repositories.PurchasingDetailRepo,
	webhookService WebhookService,
) PurchasingService {
	return &purchasingService{
		db:                   db,
		purchasingRepo:       purchasingRepo,
		itemRepo:             itemRepo,
		purchasingDetailRepo: purchasingDetailRepo,
		webhookService:       webhookService,
	}
}

func (s *purchasingService) GetPurchasings() ([]*dto.PurchasingResponse, error) {
	purchasing, err := s.purchasingRepo.FindPurchasing()

	if err != nil {
		log.Printf("failed to find purchasing: %v", err)
		return nil, errors.NewAppError(http.StatusInternalServerError, "something went wrong, please try again")
	}
	var purchasingResponse []*dto.PurchasingResponse

	for _, purchasing := range purchasing {
		purchasingResponse = append(purchasingResponse, s.toPurchasingResponse(&purchasing))
	}

	return purchasingResponse, nil
}

func (s *purchasingService) GetPurchasingByUserID(userID string) ([]*dto.PurchasingResponse, error) {
	purchasing, err := s.purchasingRepo.FindPurchasingByUserID(userID)

	if err != nil {
		log.Printf("failed to find purchasing: %v", err)
		return nil, errors.NewAppError(http.StatusInternalServerError, "something went wrong, please try again")
	}

	var purchasingResponse []*dto.PurchasingResponse

	for _, purchasing := range purchasing {
		purchasingResponse = append(purchasingResponse, s.toPurchasingResponse(&purchasing))
	}

	return purchasingResponse, nil
}

func (s *purchasingService) CreatePurchasing(req *dto.CreatePurchasingRequest) error {
	tx := s.db.Begin()

	purchasingDate, err := datetime.ParseDateWIB(req.Date)

	if err != nil {
		log.Printf("failed to parse date: %v", err)
		return errors.NewAppError(http.StatusBadRequest, "invalid date format")
	}

	data := &models.Purchasing{
		UserID:     uuid.MustParse(req.UserID),
		Date:       purchasingDate,
		SupplierID: uuid.MustParse(req.SupplierID),
		GrandTotal: 0,
	}

	purchasing, err := s.purchasingRepo.CreatePurchasing(tx, data)

	if err != nil {
		tx.Rollback()
		log.Printf("failed to create purchasing: %v", err)
		return errors.NewAppError(http.StatusInternalServerError, "something went wrong, please try again")
	}

	var details []*models.PurchasingDetail
	var itemPayloads []dto.ItemPayload
	var grandTotal int64

	for _, detail := range req.Items {
		item, err := s.itemRepo.FindByIDForUpdate(tx, detail.ItemID)
		if err != nil {
			tx.Rollback()
			log.Printf("failed to find item by id %s: %v", detail.ItemID, err)
			return errors.NewAppError(http.StatusNotFound, "item not found")
		}

		if item.Stock < detail.Qty {
			tx.Rollback()
			log.Printf("stock not enough for item %s: have %d, need %d", item.ID, item.Stock, detail.Qty)
			return errors.NewAppError(http.StatusBadRequest, "stock not enough for item "+item.Name)
		}

		subTotal := detail.Qty * item.Price

		details = append(details, &models.PurchasingDetail{
			PurchasingID: purchasing.ID,
			ItemID:       item.ID,
			Qty:          detail.Qty,
			SubTotal:     subTotal,
		})

		itemPayloads = append(itemPayloads, dto.ItemPayload{
			ItemID:    item.ID,
			ItemName:  item.Name,
			ItemPrice: item.Price,
			Qty:       detail.Qty,
			SubTotal:  subTotal,
		})

		grandTotal += subTotal

		if err := s.itemRepo.DecramentStock(tx, item, int(detail.Qty)); err != nil {
			tx.Rollback()
			log.Printf("failed to decrement stock for item %s: %v", item.ID, err)
			return errors.NewAppError(http.StatusInternalServerError, "something went wrong, please try again")
		}
	}

	if err := s.purchasingDetailRepo.BulkCreate(tx, details); err != nil {
		tx.Rollback()
		log.Printf("failed to bulk create purchasing details: %v", err)
		return errors.NewAppError(http.StatusInternalServerError, "something went wrong, please try again")
	}

	if err := s.purchasingRepo.UpdateGrandTotal(tx, purchasing.ID, grandTotal); err != nil {
		tx.Rollback()
		log.Printf("failed to update grand total: %v", err)
		return errors.NewAppError(http.StatusInternalServerError, "something went wrong, please try again")
	}

	if err := tx.Commit().Error; err != nil {
		log.Printf("failed to commit transaction: %v", err)
		return errors.NewAppError(http.StatusInternalServerError, "something went wrong, please try again")
	}

	go func() {
		payload := dto.WebhookPayload{
			PurchasingId: purchasing.ID,
			UserID:       purchasing.UserID,
			SupplierID:   purchasing.SupplierID,
			GrandTotal:   grandTotal,
			Items:        itemPayloads,
		}

		if err := s.webhookService.SendWebhook(payload); err != nil {
			log.Println("Failed to send webhook:", err)
		}
	}()

	return nil
}

func (s *purchasingService) toPurchasingResponse(purchasing *models.Purchasing) *dto.PurchasingResponse {
	return &dto.PurchasingResponse{
		ID:         purchasing.ID.String(),
		Date:       purchasing.Date.Format("2006-01-02"),
		UserID:     purchasing.UserID.String(),
		SupplierID: purchasing.SupplierID.String(),
		GrandTotal: purchasing.GrandTotal,
		Items:      s.toPurchasingDetailResponses(purchasing.Details),
	}
}

func (s *purchasingService) toPurchasingDetailResponses(details []models.PurchasingDetail) []dto.PurchasingDetailResponse {
	var responses []dto.PurchasingDetailResponse
	for _, detail := range details {
		responses = append(responses, dto.PurchasingDetailResponse{
			ItemID:    detail.ItemID.String(),
			ItemName:  detail.Items.Name,
			ItemPrice: detail.Items.Price,
			Qty:       detail.Qty,
			SubTotal:  detail.SubTotal,
		})
	}
	return responses
}

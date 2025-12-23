package services

import (
	"errors"

	"github.com/DavidAfdal/purchasing-systeam/internal/dto"
	"github.com/DavidAfdal/purchasing-systeam/internal/models"
	"github.com/DavidAfdal/purchasing-systeam/internal/repositories"
	"github.com/DavidAfdal/purchasing-systeam/pkg/datetime"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type purchasingService struct {
	db                   *gorm.DB
	purchasingRepo       repositories.PurchasingRepo
	itemRepo             repositories.ItemRepo
	purchasingDetailRepo repositories.PurchasingDetailRepo
}

type PurchasingService interface {
	CreatePurchasing(req *dto.CreatePurchasingRequest) error
}

func NewPurchasingService(
	db *gorm.DB,
	purchasingRepo repositories.PurchasingRepo,
	itemRepo repositories.ItemRepo,
	purchasingDetailRepo repositories.PurchasingDetailRepo,
) PurchasingService {
	return &purchasingService{
		db:                   db,
		purchasingRepo:       purchasingRepo,
		itemRepo:             itemRepo,
		purchasingDetailRepo: purchasingDetailRepo,
	}
}

func (s *purchasingService) CreatePurchasing(req *dto.CreatePurchasingRequest) error {
	tx := s.db.Begin()

	purchasingDate, err := datetime.ParseDateWIB(req.Date)

	if err != nil {
		return errors.New("failed to parse date")
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
		return errors.New("")
	}

	for _, detail := range req.Items {

		item, err := s.itemRepo.FindByIDForUpdate(tx, uuid.MustParse(detail.ItemID))

		if err != nil {
			tx.Rollback()
			return errors.New("")
		}

		if item.Stock < detail.Qty {
			tx.Rollback()
			return errors.New("")
		}

		subTotal := detail.Qty * item.Price

		data := &models.PurchasingDetail{
			PurchasingID: purchasing.ID,
			ItemID:       item.ID,
			Qty:          detail.Qty,
			SubTotal:     subTotal,
		}

		if err := s.purchasingDetailRepo.Create(tx, data); err != nil {
			tx.Rollback()
			return errors.New("")
		}

		if err := s.itemRepo.DecramentStock(tx, item, int(detail.Qty)); err != nil {
			tx.Rollback()
			return errors.New("")
		}
	}

	return tx.Commit().Error
}

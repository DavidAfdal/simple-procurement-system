package repositories

import (
	"github.com/DavidAfdal/purchasing-systeam/internal/models"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type PurchasingRepo interface {
	FindPurchasing() ([]models.Purchasing, error)
	FindPurchasingByUserID(userID string) ([]models.Purchasing, error)
	CreatePurchasing(tx *gorm.DB, purchasing *models.Purchasing) (*models.Purchasing, error)
	UpdateGrandTotal(tx *gorm.DB, purchasingID uuid.UUID, total int64) error
}
type purchasingRepo struct {
	db *gorm.DB
}

func NewPurchasingRepo(db *gorm.DB) PurchasingRepo {
	return &purchasingRepo{db: db}
}

func (r *purchasingRepo) CreatePurchasing(tx *gorm.DB, purchasing *models.Purchasing) (*models.Purchasing, error) {
	if err := tx.Create(purchasing).Error; err != nil {
		return nil, err
	}
	return purchasing, nil
}

func (r *purchasingRepo) FindPurchasing() ([]models.Purchasing, error) {
	var purchasing []models.Purchasing
	if err := r.db.Preload("Details.Items").Find(&purchasing).Error; err != nil {
		return nil, err
	}
	return purchasing, nil
}

func (r *purchasingRepo) FindPurchasingByUserID(userID string) ([]models.Purchasing, error) {
	var purchasings []models.Purchasing
	if err := r.db.Where("user_id = ?", userID).Preload("Details.Items").Find(&purchasings).Error; err != nil {
		return nil, err
	}
	return purchasings, nil
}

func (r *purchasingRepo) UpdateGrandTotal(tx *gorm.DB, purchasingID uuid.UUID, total int64) error {
	return tx.Model(&models.Purchasing{}).
		Where("id = ?", purchasingID).
		Update("grand_total", total).Error
}

package repositories

import (
	"github.com/DavidAfdal/purchasing-systeam/internal/models"
	"gorm.io/gorm"
)

type PurchasingRepo interface {
	CreatePurchasing(tx *gorm.DB, purchasing *models.Purchasing) (*models.Purchasing, error)
	GetPurchasingByID(id string) (*models.Purchasing, error)
	GetPurchasingByUserID(userID string) ([]models.Purchasing, error)
}
type purchasingRepo struct {
	db *gorm.DB
}

func NewPurchasingRepo(db *gorm.DB) PurchasingRepo {
	return &purchasingRepo{db: db}
}

func (p *purchasingRepo) CreatePurchasing(tx *gorm.DB, purchasing *models.Purchasing) (*models.Purchasing, error) {
	if err := tx.Create(purchasing).Error; err != nil {
		return nil, err
	}
	return purchasing, nil
}

func (p *purchasingRepo) GetPurchasingByID(id string) (*models.Purchasing, error) {
	var purchasing models.Purchasing
	if err := p.db.Where("id = ?", id).First(&purchasing).Error; err != nil {
		return nil, err
	}
	return &purchasing, nil
}

func (p *purchasingRepo) GetPurchasingByUserID(userID string) ([]models.Purchasing, error) {
	var purchasings []models.Purchasing
	if err := p.db.Where("user_id = ?", userID).Preload("Details").Find(&purchasings).Error; err != nil {
		return nil, err
	}
	return purchasings, nil
}

package repositories

import (
	"github.com/DavidAfdal/purchasing-systeam/internal/models"
	"gorm.io/gorm"
)

type PurchasingDetailRepo interface {
	BulkCreate(tx *gorm.DB, details []*models.PurchasingDetail) error
}

type purchasingDetailRepo struct {
}

func NewPurchasingDetailRepo() PurchasingDetailRepo {
	return &purchasingDetailRepo{}
}

func (r *purchasingDetailRepo) BulkCreate(tx *gorm.DB, details []*models.PurchasingDetail) error {
	return tx.Create(&details).Error
}

package models

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type PurchasingDetail struct {
	ID           uuid.UUID `gorm:"type:char(36);primaryKey"`
	PurchasingID uuid.UUID `gorm:"type:char(36);not null"`
	ItemID       uuid.UUID `gorm:"type:char(36);not null"`
	Qty          int64     `gorm:"type:bigint;not null"`
	SubTotal     int64     `gorm:"type:bigint;not null"`

	Purchasing Purchasing `gorm:"foreignKey:PurchasingID"`
	Items      Item       `gorm:"foreignKey:ItemID"`
}

func (pd *PurchasingDetail) BeforeCreate(tx *gorm.DB) error {
	pd.ID = uuid.New()
	return nil
}

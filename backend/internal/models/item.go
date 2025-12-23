package models

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Item struct {
	ID    uuid.UUID `gorm:"type:char(36);primaryKey"`
	Name  string    `gorm:"type:varchar(150);not null"`
	Stock int64     `gorm:"type:bigint;not null"`
	Price int64     `gorm:"type:bigint;not null"`

	PurchasingDetails []PurchasingDetail `gorm:"foreignKey:ItemID"`
}

func (p *Item) BeforeCreate(tx *gorm.DB) error {
	p.ID = uuid.New()
	return nil
}

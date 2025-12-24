package models

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Supplier struct {
	ID      uuid.UUID `gorm:"type:char(36);primaryKey"`
	Name    string    `gorm:"type:varchar(150);not null"`
	Email   string    `gorm:"type:varchar(150);unique;not null"`
	Address string    `gorm:"type:varchar(255);unique;not null"`

	Purchasings []Purchasing `gorm:"foreignKey:SupplierID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE"`
}

func (s *Supplier) BeforeCreate(tx *gorm.DB) error {
	s.ID = uuid.New()
	return nil
}

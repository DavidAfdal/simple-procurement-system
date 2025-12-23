package models

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Supplier struct {
	ID      uuid.UUID `gorm:"type:char(36);primaryKey"`
	Name    string    `gorm:"type:varchar(150);not null"`
	Email   string    `gorm:"type:varchar(150)"`
	Address string    `gorm:"type:varchar(255)"`

	Purchasings []Purchasing `gorm:"foreignKey:SupplierID"`
}

func (s *Supplier) BeforeCreate(tx *gorm.DB) error {
	s.ID = uuid.New()
	return nil
}

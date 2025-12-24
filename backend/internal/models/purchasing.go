package models

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Purchasing struct {
	ID         uuid.UUID `gorm:"type:char(36);primaryKey"`
	Date       time.Time `gorm:"type:date;not null"`
	UserID     uuid.UUID `gorm:"type:char(36);not null"`
	SupplierID uuid.UUID `gorm:"type:char(36);not null"`
	GrandTotal int64     `gorm:"type:bigint;not null"`

	User     User               `gorm:"foreignKey:UserID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE"`
	Supplier Supplier           `gorm:"foreignKey:SupplierID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE"`
	Details  []PurchasingDetail `gorm:"foreignKey:PurchasingID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE"`
}

func (p *Purchasing) BeforeCreate(tx *gorm.DB) error {
	p.ID = uuid.New()
	return nil
}

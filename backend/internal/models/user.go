package models

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type User struct {
	ID       uuid.UUID `gorm:"type:char(36);primaryKey"`
	Username string    `gorm:"type:varchar(100);unique;not null"`
	Password string    `gorm:"type:varchar(255);not null"`
	Role     string    `gorm:"type:varchar(50);not null"`

	Purchasings []Purchasing `gorm:"foreignKey:UserID"`
}

func (u *User) BeforeCreate(tx *gorm.DB) error {
	u.ID = uuid.New()
	return nil
}

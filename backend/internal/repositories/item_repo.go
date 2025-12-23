package repositories

import (
	"github.com/DavidAfdal/purchasing-systeam/internal/models"
	"github.com/google/uuid"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

type ItemRepo interface {
	FindItems() ([]models.Item, error)
	FindItemByID(id string) (models.Item, error)
	FindByIDForUpdate(tx *gorm.DB, id uuid.UUID) (*models.Item, error)
	CreateItem(item *models.Item) error
	UpdateItem(item *models.Item) error
	DeleteItem(item *models.Item) error
	DecramentStock(tx *gorm.DB, item *models.Item, qty int) error
}
type itemRepo struct {
	db *gorm.DB
}

func NewItemRepo(db *gorm.DB) ItemRepo {
	return &itemRepo{db: db}
}

func (i *itemRepo) FindItems() ([]models.Item, error) {
	var items []models.Item
	if err := i.db.Find(&items).Error; err != nil {
		return nil, err
	}
	return items, nil
}

func (i *itemRepo) FindItemByID(id string) (models.Item, error) {
	var item models.Item
	if err := i.db.Where("id = ?", id).First(&item).Error; err != nil {
		return models.Item{}, err
	}
	return item, nil
}

func (r *itemRepo) FindByIDForUpdate(tx *gorm.DB, id uuid.UUID) (*models.Item, error) {
	var item models.Item
	err := tx.Clauses(clause.Locking{Strength: "UPDATE"}).
		First(&item, "id = ?", id).Error
	return &item, err
}

func (i *itemRepo) CreateItem(item *models.Item) error {
	if err := i.db.Create(item).Error; err != nil {
		return err
	}
	return nil
}

func (i *itemRepo) UpdateItem(item *models.Item) error {
	if err := i.db.Save(item).Error; err != nil {
		return err
	}

	return nil
}

func (i *itemRepo) DecramentStock(tx *gorm.DB, item *models.Item, qty int) error {
	if err := tx.Model(item).Update("stock", gorm.Expr("stock - ?", qty)).Error; err != nil {
		return err
	}
	return nil
}

func (i *itemRepo) DeleteItem(item *models.Item) error {
	if err := i.db.Delete(item).Error; err != nil {
		return err
	}
	return nil
}

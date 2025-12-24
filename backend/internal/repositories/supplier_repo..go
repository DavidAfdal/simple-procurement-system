package repositories

import (
	"github.com/DavidAfdal/purchasing-systeam/internal/models"
	"gorm.io/gorm"
)

type SupplierRepo interface {
	FindSuppliers() ([]models.Supplier, error)
	FindSupplierByID(id string) (*models.Supplier, error)
	CreateSupplier(supplier *models.Supplier) (*models.Supplier, error)
	UpdateSupplier(supplier *models.Supplier) error
	DeleteSupplier(supplier *models.Supplier) error
}

type supplierRepo struct {
	db *gorm.DB
}

func NewSupplierRepo(db *gorm.DB) SupplierRepo {
	return &supplierRepo{db: db}
}

func (s *supplierRepo) FindSuppliers() ([]models.Supplier, error) {
	var suppliers []models.Supplier
	if err := s.db.Find(&suppliers).Error; err != nil {
		return nil, err
	}
	return suppliers, nil
}

func (s *supplierRepo) FindSupplierByID(id string) (*models.Supplier, error) {
	var supplier models.Supplier
	if err := s.db.Where("id = ?", id).First(&supplier).Error; err != nil {
		return nil, err
	}
	return &supplier, nil
}

func (s *supplierRepo) CreateSupplier(supplier *models.Supplier) (*models.Supplier, error) {
	if err := s.db.Create(supplier).Error; err != nil {
		return nil, err
	}
	return supplier, nil
}

func (s *supplierRepo) UpdateSupplier(supplier *models.Supplier) error {
	if err := s.db.Save(supplier).Error; err != nil {
		return err
	}
	return nil
}

func (s *supplierRepo) DeleteSupplier(supplier *models.Supplier) error {
	if err := s.db.Delete(&supplier).Error; err != nil {
		return err
	}
	return nil
}

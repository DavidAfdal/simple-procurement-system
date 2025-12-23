package services

import (
	"github.com/DavidAfdal/purchasing-systeam/internal/dto"
	"github.com/DavidAfdal/purchasing-systeam/internal/models"
	"github.com/DavidAfdal/purchasing-systeam/internal/repositories"
)

type SupplierService interface {
	GetSuppliers() ([]*dto.SupplierResponse, error)
	GetSupplierByID(id string) (*dto.SupplierResponse, error)
	CreateSupplier(req *dto.CreateSupplierRequest) (*dto.SupplierResponse, error)
	UpdateSupplier(id string, req *dto.UpdateSupplierRequest) (*dto.SupplierResponse, error)
	DeleteSupplier(id string) error
}

type supplierService struct {
	supplierRepo repositories.SupplierRepo
}

func NewSupplierService(supplierRepo repositories.SupplierRepo) SupplierService {
	return &supplierService{
		supplierRepo: supplierRepo,
	}
}

func (s *supplierService) GetSuppliers() ([]*dto.SupplierResponse, error) {
	suppliers, err := s.supplierRepo.GetSuppliers()
	if err != nil {
		return nil, err
	}
	var supplierResponses []*dto.SupplierResponse
	for _, supplier := range suppliers {
		supplierResponses = append(supplierResponses, s.toSupplierResponse(&supplier))
	}
	return supplierResponses, nil
}

func (s *supplierService) GetSupplierByID(id string) (*dto.SupplierResponse, error) {
	supplier, err := s.supplierRepo.GetSupplierByID(id)
	if err != nil {
		return nil, err
	}
	return s.toSupplierResponse(supplier), nil
}

func (s *supplierService) CreateSupplier(req *dto.CreateSupplierRequest) (*dto.SupplierResponse, error) {
	supplier := &models.Supplier{
		Name:    req.Name,
		Email:   req.Email,
		Address: req.Address,
	}

	supplier, err := s.supplierRepo.CreateSupplier(supplier)

	if err != nil {
		return nil, err
	}
	return s.toSupplierResponse(supplier), nil
}

func (s *supplierService) UpdateSupplier(id string, req *dto.UpdateSupplierRequest) (*dto.SupplierResponse, error) {
	supplier, err := s.supplierRepo.GetSupplierByID(id)

	if err != nil {
		return nil, err
	}

	supplier.Name = req.Name
	supplier.Email = req.Email
	supplier.Address = req.Address

	if err := s.supplierRepo.UpdateSupplier(supplier); err != nil {
		return nil, err
	}

	return s.toSupplierResponse(supplier), nil
}

func (s *supplierService) DeleteSupplier(id string) error {
	supplier, err := s.supplierRepo.GetSupplierByID(id)

	if err != nil {
		return err
	}

	return s.supplierRepo.DeleteSupplier(supplier)
}

func (s *supplierService) toSupplierResponse(supplier *models.Supplier) *dto.SupplierResponse {
	return &dto.SupplierResponse{
		ID:      supplier.ID.String(),
		Email:   supplier.Email,
		Name:    supplier.Name,
		Address: supplier.Address,
	}
}

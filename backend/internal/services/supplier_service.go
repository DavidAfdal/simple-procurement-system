package services

import (
	"log"
	"net/http"

	"github.com/DavidAfdal/purchasing-systeam/internal/dto"
	"github.com/DavidAfdal/purchasing-systeam/internal/models"
	"github.com/DavidAfdal/purchasing-systeam/internal/repositories"
	"github.com/DavidAfdal/purchasing-systeam/pkg/errors"
	"github.com/go-sql-driver/mysql"
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
	suppliers, err := s.supplierRepo.FindSuppliers()
	if err != nil {
		log.Printf("failed to fetch suppliers: %v", err)
		return nil, errors.NewAppError(http.StatusInternalServerError, "something went wrong, please try again")
	}

	var supplierResponses []*dto.SupplierResponse
	for _, supplier := range suppliers {
		supplierResponses = append(supplierResponses, s.toSupplierResponse(&supplier))
	}
	return supplierResponses, nil
}

func (s *supplierService) CreateSupplier(req *dto.CreateSupplierRequest) (*dto.SupplierResponse, error) {
	supplier := &models.Supplier{
		Name:    req.Name,
		Email:   req.Email,
		Address: req.Address,
	}

	supplier, err := s.supplierRepo.CreateSupplier(supplier)
	if err != nil {
		if mysqlErr, ok := err.(*mysql.MySQLError); ok && mysqlErr.Number == 1062 {
			log.Printf("duplicate supplier email attempt: %s", req.Email)
			return nil, errors.NewAppError(http.StatusBadRequest, "supplier email already exists")
		}
		log.Printf("failed to create supplier %s: %v", req.Email, err)
		return nil, errors.NewAppError(http.StatusInternalServerError, "something went wrong, please try again")
	}

	return s.toSupplierResponse(supplier), nil
}

func (s *supplierService) UpdateSupplier(id string, req *dto.UpdateSupplierRequest) (*dto.SupplierResponse, error) {
	supplier, err := s.supplierRepo.FindSupplierByID(id)
	if err != nil {
		log.Printf("supplier not found with id %s: %v", id, err)
		return nil, errors.NewAppError(http.StatusNotFound, "supplier not found")
	}

	supplier.Name = req.Name
	supplier.Email = req.Email
	supplier.Address = req.Address

	if err := s.supplierRepo.UpdateSupplier(supplier); err != nil {
		log.Printf("failed to update supplier %s: %v", id, err)
		return nil, errors.NewAppError(http.StatusInternalServerError, "something went wrong, please try again")
	}

	return s.toSupplierResponse(supplier), nil
}

func (s *supplierService) DeleteSupplier(id string) error {
	supplier, err := s.supplierRepo.FindSupplierByID(id)
	if err != nil {
		log.Printf("supplier not found with id %s: %v", id, err)
		return errors.NewAppError(http.StatusNotFound, "supplier not found")
	}

	if err := s.supplierRepo.DeleteSupplier(supplier); err != nil {
		log.Printf("failed to delete supplier %s: %v", id, err)
		return errors.NewAppError(http.StatusInternalServerError, "something went wrong, please try again")
	}

	return nil
}

func (s *supplierService) GetSupplierByID(id string) (*dto.SupplierResponse, error) {
	supplier, err := s.supplierRepo.FindSupplierByID(id)
	if err != nil {
		log.Printf("supplier not found with id %s: %v", id, err)
		return nil, errors.NewAppError(http.StatusNotFound, "supplier not found")
	}
	return s.toSupplierResponse(supplier), nil
}

func (s *supplierService) toSupplierResponse(supplier *models.Supplier) *dto.SupplierResponse {
	return &dto.SupplierResponse{
		ID:      supplier.ID.String(),
		Email:   supplier.Email,
		Name:    supplier.Name,
		Address: supplier.Address,
	}
}

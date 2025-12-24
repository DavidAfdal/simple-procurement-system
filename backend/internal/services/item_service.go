package services

import (
	"log"
	"net/http"
	"strings"

	"github.com/DavidAfdal/purchasing-systeam/internal/dto"
	"github.com/DavidAfdal/purchasing-systeam/internal/models"
	"github.com/DavidAfdal/purchasing-systeam/internal/repositories"
	"github.com/DavidAfdal/purchasing-systeam/pkg/errors"
)

type ItemService interface {
	GetItems() ([]*dto.ItemResponse, error)
	GetItemByID(id string) (*dto.ItemResponse, error)
	CreateItem(req *dto.CreateItemRequest) (*dto.ItemResponse, error)
	UpdateItem(id string, req *dto.UpdateItemRequest) (*dto.ItemResponse, error)
	DeleteItem(id string) error
}

type itemService struct {
	itemRepo repositories.ItemRepo
}

func NewItemService(itemRepo repositories.ItemRepo) ItemService {
	return &itemService{
		itemRepo: itemRepo,
	}
}

func (s *itemService) GetItems() ([]*dto.ItemResponse, error) {
	items, err := s.itemRepo.FindItems()
	if err != nil {
		log.Printf("failed to fetch items: %v", err)
		return nil, errors.NewAppError(http.StatusInternalServerError, "something went wrong, please try again")
	}

	var itemResponses []*dto.ItemResponse
	for _, item := range items {
		itemResponses = append(itemResponses, s.toItemResponse(&item))
	}
	return itemResponses, nil
}

func (s *itemService) GetItemByID(id string) (*dto.ItemResponse, error) {
	item, err := s.itemRepo.FindItemByID(id)
	if err != nil {
		log.Printf("item not found with id %s: %v", id, err)
		return nil, errors.NewAppError(http.StatusNotFound, "item not found")
	}
	return s.toItemResponse(item), nil
}

func (s *itemService) CreateItem(req *dto.CreateItemRequest) (*dto.ItemResponse, error) {
	data := &models.Item{
		Name:  req.Name,
		Stock: req.Stock,
		Price: req.Price,
	}

	item, err := s.itemRepo.CreateItem(data)
	if err != nil {
		log.Printf("failed to create item: %v", err)
		if strings.Contains(err.Error(), "duplicate") {
			return nil, errors.NewAppError(http.StatusBadRequest, "item already exists")
		}
		return nil, errors.NewAppError(http.StatusInternalServerError, "something went wrong, please try again")
	}

	return s.toItemResponse(item), nil
}

func (s *itemService) UpdateItem(id string, req *dto.UpdateItemRequest) (*dto.ItemResponse, error) {
	item, err := s.itemRepo.FindItemByID(id)
	if err != nil {
		log.Printf("item not found with id %s: %v", id, err)
		return nil, errors.NewAppError(http.StatusNotFound, "item not found")
	}

	item.Name = req.Name
	item.Stock = req.Stock
	item.Price = req.Price

	if err := s.itemRepo.UpdateItem(item); err != nil {
		log.Printf("failed to update item %s: %v", id, err)
		return nil, errors.NewAppError(http.StatusInternalServerError, "something went wrong, please try again")
	}

	return s.toItemResponse(item), nil
}

func (s *itemService) DeleteItem(id string) error {
	item, err := s.itemRepo.FindItemByID(id)
	if err != nil {
		log.Printf("item not found with id %s: %v", id, err)
		return errors.NewAppError(http.StatusNotFound, "item not found")
	}

	if err := s.itemRepo.DeleteItem(item); err != nil {
		log.Printf("failed to delete item %s: %v", id, err)
		return errors.NewAppError(http.StatusInternalServerError, "something went wrong, please try again")
	}

	return nil
}

func (s *itemService) toItemResponse(item *models.Item) *dto.ItemResponse {
	return &dto.ItemResponse{
		ID:    item.ID.String(),
		Name:  item.Name,
		Stock: item.Stock,
		Price: item.Price,
	}
}

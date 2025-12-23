package services

import (
	"github.com/DavidAfdal/purchasing-systeam/internal/dto"
	"github.com/DavidAfdal/purchasing-systeam/internal/models"
	"github.com/DavidAfdal/purchasing-systeam/internal/repositories"
)

type ItemService interface {
	FindItems() ([]*dto.ItemResponse, error)
	FindItemByID(id string) (*dto.ItemResponse, error)
	CreateItem(req *dto.CreateItemRequest) error
	UpdateItem(id string, req *dto.UpdateItemRequest) error
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

func (s *itemService) FindItems() ([]*dto.ItemResponse, error) {
	items, err := s.itemRepo.FindItems()
	if err != nil {
		return nil, err
	}
	var itemResponses []*dto.ItemResponse
	for _, item := range items {
		itemResponses = append(itemResponses, s.toItemResponse(&item))
	}
	return itemResponses, nil
}

func (s *itemService) FindItemByID(id string) (*dto.ItemResponse, error) {
	item, err := s.itemRepo.FindItemByID(id)
	if err != nil {
		return nil, err
	}
	return s.toItemResponse(&item), nil
}

func (s *itemService) CreateItem(req *dto.CreateItemRequest) error {

	return s.itemRepo.CreateItem(&models.Item{
		Name:  req.Name,
		Stock: req.Stock,
		Price: req.Price,
	})
}

func (s *itemService) UpdateItem(id string, req *dto.UpdateItemRequest) error {
	item, err := s.itemRepo.FindItemByID(id)

	if err != nil {
		return err
	}
	item.Name = req.Name
	item.Stock = req.Stock
	item.Price = req.Price
	return s.itemRepo.UpdateItem(&item)
}

func (s *itemService) DeleteItem(id string) error {
	item, err := s.itemRepo.FindItemByID(id)
	if err != nil {
		return err
	}
	return s.itemRepo.DeleteItem(&item)
}

func (s *itemService) toItemResponse(item *models.Item) *dto.ItemResponse {
	return &dto.ItemResponse{
		ID:    item.ID.String(),
		Name:  item.Name,
		Stock: item.Stock,
		Price: item.Price,
	}
}

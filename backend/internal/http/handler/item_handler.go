package handler

import (
	"github.com/DavidAfdal/purchasing-systeam/internal/services"
	"github.com/gin-gonic/gin"
)

type ItemHandler struct {
	itemService services.ItemService
}

func NewItemHandler(itemService services.ItemService) *ItemHandler {
	return &ItemHandler{
		itemService: itemService,
	}
}

func (h *ItemHandler) FindItems(c *gin.Context) {

}

func (h *ItemHandler) FindItemByID(c *gin.Context) {

}

func (h *ItemHandler) CreateItem(c *gin.Context) {

}

func (h *ItemHandler) UpdateItem(c *gin.Context) {

}

func (h *ItemHandler) DeleteItem(c *gin.Context) {

}

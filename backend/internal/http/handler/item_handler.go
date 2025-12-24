package handler

import (
	"net/http"

	"github.com/DavidAfdal/purchasing-systeam/internal/dto"
	"github.com/DavidAfdal/purchasing-systeam/internal/services"
	"github.com/DavidAfdal/purchasing-systeam/pkg/response"
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

func (h *ItemHandler) GetItems(c *gin.Context) {
	items, err := h.itemService.GetItems()
	if err != nil {
		handleErrorService(c, err)
		return
	}
	response.SuccessResponse(c, http.StatusOK, "items retrieved successfully", items)
}

func (h *ItemHandler) GetItemByID(c *gin.Context) {
	itemId := c.Param("id")

	item, err := h.itemService.GetItemByID(itemId)
	if err != nil {
		handleErrorService(c, err)
		return
	}
	response.SuccessResponse(c, http.StatusOK, "item retrieved successfully", item)
}

func (h *ItemHandler) CreateItem(c *gin.Context) {
	var req dto.CreateItemRequest

	if err := c.ShouldBindJSON(&req); err != nil {
		response.ErrorResponse(c, http.StatusBadRequest, "invalid request")
		return
	}

	if errMsg, data := checkValidation(req); errMsg != "" {
		response.ErrorResponse(c, http.StatusBadRequest, errMsg, data)
		return
	}

	res, err := h.itemService.CreateItem(&req)
	if err != nil {
		handleErrorService(c, err)
		return
	}

	response.SuccessResponse(c, http.StatusCreated, "item created successfully", res)
}

func (h *ItemHandler) UpdateItem(c *gin.Context) {
	itemId := c.Param("id")
	var req dto.UpdateItemRequest

	if err := c.ShouldBindJSON(&req); err != nil {
		response.ErrorResponse(c, http.StatusBadRequest, "invalid request")
		return
	}

	if errMsg, data := checkValidation(req); errMsg != "" {
		response.ErrorResponse(c, http.StatusBadRequest, errMsg, data)
		return
	}

	res, err := h.itemService.UpdateItem(itemId, &req)
	if err != nil {
		handleErrorService(c, err)
		return
	}

	response.SuccessResponse(c, http.StatusOK, "item updated successfully", res)
}

func (h *ItemHandler) DeleteItem(c *gin.Context) {
	itemId := c.Param("id")

	if err := h.itemService.DeleteItem(itemId); err != nil {
		handleErrorService(c, err)
		return
	}

	response.SuccessResponse(c, http.StatusOK, "item deleted successfully", nil)
}

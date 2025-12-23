package handler

import (
	"github.com/DavidAfdal/purchasing-systeam/internal/dto"
	"github.com/DavidAfdal/purchasing-systeam/internal/services"
	"github.com/DavidAfdal/purchasing-systeam/pkg/response"
	"github.com/gin-gonic/gin"
)

type SupplierHandler struct {
	supplierService services.SupplierService
}

func NewSupplierHandler(supplierService services.SupplierService) *SupplierHandler {
	return &SupplierHandler{supplierService: supplierService}
}

func (h *SupplierHandler) GetSuppliers(c *gin.Context) {
	res, err := h.supplierService.GetSuppliers()
	if err != nil {
		response.ErrorResponse(c, 400, err.Error())
		return
	}
	response.SuccessResponse(c, 200, "success", res)
}

func (h *SupplierHandler) GetSupplier(c *gin.Context) {
	id := c.Param("id")
	res, err := h.supplierService.GetSupplierByID(id)
	if err != nil {
		response.ErrorResponse(c, 400, err.Error())
		return
	}
	response.SuccessResponse(c, 200, "success", res)
}

func (h *SupplierHandler) CreateSupplier(c *gin.Context) {
	var req dto.CreateSupplierRequest

	if err := c.ShouldBindJSON(&req); err != nil {
		response.ErrorResponse(c, 400, "invalid request")
		return
	}

	if errMsg, data := checkValidation(req); errMsg != "" {
		response.ErrorResponse(c, 400, errMsg, data)
		return
	}

	res, err := h.supplierService.CreateSupplier(&req)
	if err != nil {
		response.ErrorResponse(c, 400, err.Error())
		return
	}
	response.SuccessResponse(c, 200, "success", res)
}

func (h *SupplierHandler) UpdateSupplier(c *gin.Context) {
	var req dto.UpdateSupplierRequest

	id := c.Param("id")

	if err := c.ShouldBindJSON(&req); err != nil {
		response.ErrorResponse(c, 400, "invalid request")
		return
	}

	if errMsg, data := checkValidation(req); errMsg != "" {
		response.ErrorResponse(c, 400, errMsg, data)
		return
	}

	res, err := h.supplierService.UpdateSupplier(id, &req)
	if err != nil {
		response.ErrorResponse(c, 400, err.Error())
		return
	}
	response.SuccessResponse(c, 200, "success", res)
}

func (h *SupplierHandler) DeleteSupplier(c *gin.Context) {
	id := c.Param("id")

	if err := h.supplierService.DeleteSupplier(id); err != nil {
		response.ErrorResponse(c, 400, err.Error())
		return
	}

	response.SuccessResponse(c, 200, "success", nil)
}

package handler

import (
	"net/http"

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
		handleErrorService(c, err)
		return
	}

	response.SuccessResponse(c, http.StatusOK, "suppliers retrieved successfully", res)
}

func (h *SupplierHandler) GetSupplier(c *gin.Context) {
	id := c.Param("id")
	res, err := h.supplierService.GetSupplierByID(id)
	if err != nil {
		handleErrorService(c, err)
		return
	}
	response.SuccessResponse(c, http.StatusOK, "supplier retrieved successfully", res)
}

func (h *SupplierHandler) CreateSupplier(c *gin.Context) {
	var req dto.CreateSupplierRequest

	if err := c.ShouldBindJSON(&req); err != nil {
		response.ErrorResponse(c, http.StatusBadRequest, "invalid request")
		return
	}

	if errMsg, data := checkValidation(req); errMsg != "" {
		response.ErrorResponse(c, http.StatusBadRequest, errMsg, data)
		return
	}

	res, err := h.supplierService.CreateSupplier(&req)
	if err != nil {
		handleErrorService(c, err)
		return
	}
	response.SuccessResponse(c, http.StatusCreated, "supplier created successfully", res)
}

func (h *SupplierHandler) UpdateSupplier(c *gin.Context) {
	id := c.Param("id")
	var req dto.UpdateSupplierRequest

	if err := c.ShouldBindJSON(&req); err != nil {
		response.ErrorResponse(c, http.StatusBadRequest, "invalid request")
		return
	}

	if errMsg, data := checkValidation(req); errMsg != "" {
		response.ErrorResponse(c, http.StatusBadRequest, errMsg, data)
		return
	}

	res, err := h.supplierService.UpdateSupplier(id, &req)
	if err != nil {
		handleErrorService(c, err)
		return
	}
	response.SuccessResponse(c, http.StatusOK, "supplier updated successfully", res)
}

func (h *SupplierHandler) DeleteSupplier(c *gin.Context) {
	id := c.Param("id")

	if err := h.supplierService.DeleteSupplier(id); err != nil {
		handleErrorService(c, err)
		return
	}

	response.SuccessResponse(c, http.StatusOK, "supplier deleted successfully", nil)
}

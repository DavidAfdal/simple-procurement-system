package handler

import (
	"net/http"

	"github.com/DavidAfdal/purchasing-systeam/pkg/errors"
	"github.com/DavidAfdal/purchasing-systeam/pkg/response"
	"github.com/DavidAfdal/purchasing-systeam/pkg/validator"
	"github.com/gin-gonic/gin"
)

type Handler struct {
	UserHandler       *UserHandler
	SupplierHandler   *SupplierHandler
	ItemHandler       *ItemHandler
	PurchasingHandler *PurchasingHandler
}

func NewHandler(userHandler *UserHandler, supplierHandler *SupplierHandler, itemHandler *ItemHandler, purchasingHandler *PurchasingHandler) *Handler {
	return &Handler{
		UserHandler:       userHandler,
		SupplierHandler:   supplierHandler,
		ItemHandler:       itemHandler,
		PurchasingHandler: purchasingHandler,
	}
}

func checkValidation(input interface{}) (errorMessage string, data interface{}) {
	validationErrors := validator.Validate(input)
	if validationErrors != nil {
		if _, exists := validationErrors["error"]; exists {
			return "input validation failed", nil
		}
		return "input validation failed", validationErrors
	}
	return "", nil
}

func handleErrorService(c *gin.Context, err error) {
	if appErr, ok := err.(*errors.AppError); ok {
		response.ErrorResponse(c, appErr.Code, appErr.Message)
		return
	}
	response.ErrorResponse(c, http.StatusInternalServerError, "something went wrong, please try again")
}

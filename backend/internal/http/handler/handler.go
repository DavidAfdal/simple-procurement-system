package handler

import "github.com/DavidAfdal/purchasing-systeam/pkg/validator"

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
			return "validasi input gagal", nil
		}
		return "validasi input gagal", validationErrors
	}
	return "", nil
}

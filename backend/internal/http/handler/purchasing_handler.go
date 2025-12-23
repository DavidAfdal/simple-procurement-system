package handler

import (
	"github.com/DavidAfdal/purchasing-systeam/internal/dto"
	"github.com/DavidAfdal/purchasing-systeam/internal/services"
	"github.com/DavidAfdal/purchasing-systeam/pkg/response"
	"github.com/gin-gonic/gin"
)

type PurchasingHandler struct {
	purchasingService services.PurchasingService
}

func NewPurchasingHandler(purchasingService services.PurchasingService) *PurchasingHandler {
	return &PurchasingHandler{
		purchasingService: purchasingService,
	}
}

func (h *PurchasingHandler) CreatePurchasing(c *gin.Context) {
	var req dto.CreatePurchasingRequest

	if err := c.ShouldBindJSON(&req); err != nil {
		response.ErrorResponse(c, 400, "invalid request")
		return
	}

	if errMsg, data := checkValidation(req); errMsg != "" {
		response.ErrorResponse(c, 400, errMsg, data)
		return
	}

	req.UserID = c.MustGet("user_id").(string)

	if err := h.purchasingService.CreatePurchasing(&req); err != nil {
		response.ErrorResponse(c, 400, err.Error())
		return
	}

	response.SuccessResponse(c, 200, "success", nil)
}

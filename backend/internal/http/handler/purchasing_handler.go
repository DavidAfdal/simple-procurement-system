package handler

import (
	"net/http"

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
		response.ErrorResponse(c, http.StatusBadRequest, "invalid request")
		return
	}

	if errMsg, data := checkValidation(req); errMsg != "" {
		response.ErrorResponse(c, http.StatusBadRequest, errMsg, data)
		return
	}

	req.UserID = c.MustGet("user_id").(string)

	if err := h.purchasingService.CreatePurchasing(&req); err != nil {
		handleErrorService(c, err)
		return
	}

	response.SuccessResponse(c, http.StatusCreated, "purchasing created successfully", nil)
}

func (h *PurchasingHandler) GetMyPurchasings(c *gin.Context) {
	userID := c.MustGet("user_id").(string)

	res, err := h.purchasingService.GetPurchasingByUserID(userID)

	if err != nil {
		handleErrorService(c, err)
		return
	}
	response.SuccessResponse(c, http.StatusOK, "your purchasing retrieved successfully", res)
}

func (h *PurchasingHandler) GetPurchasing(c *gin.Context) {

	res, err := h.purchasingService.GetPurchasings()
	if err != nil {
		handleErrorService(c, err)
		return
	}
	response.SuccessResponse(c, http.StatusOK, "purchasing retrieved successfully", res)
}

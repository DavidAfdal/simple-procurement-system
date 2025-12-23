package handler

import (
	"github.com/DavidAfdal/purchasing-systeam/internal/dto"
	"github.com/DavidAfdal/purchasing-systeam/internal/services"
	"github.com/DavidAfdal/purchasing-systeam/pkg/response"
	"github.com/gin-gonic/gin"
)

type UserHandler struct {
	userService services.UserService
}

func NewUserHandler(userService services.UserService) *UserHandler {
	return &UserHandler{userService: userService}
}

func (h *UserHandler) Register(c *gin.Context) {
	var req dto.RegisterRequest

	if err := c.ShouldBindJSON(&req); err != nil {
		response.ErrorResponse(c, 400, "invalid request")
		return
	}

	if errMsg, data := checkValidation(req); errMsg != "" {
		response.ErrorResponse(c, 400, errMsg, data)
		return
	}

	res, err := h.userService.Register(&req)

	if err != nil {
		response.ErrorResponse(c, 400, err.Error())
		return
	}

	response.SuccessResponse(c, 201, "success", res)
}

func (h *UserHandler) Login(c *gin.Context) {
	var req dto.LoginRequest

	if err := c.ShouldBindJSON(&req); err != nil {
		response.ErrorResponse(c, 400, "invalid request")
		return
	}

	if errMsg, data := checkValidation(req); errMsg != "" {
		response.ErrorResponse(c, 400, errMsg, data)
		return
	}

	res, err := h.userService.Login(&req)

	if err != nil {
		response.ErrorResponse(c, 400, err.Error())
		return
	}

	response.SuccessResponse(c, 200, "success", res)
}

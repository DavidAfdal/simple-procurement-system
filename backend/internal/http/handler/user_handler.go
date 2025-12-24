package handler

import (
	"net/http"
	"strings"

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
		response.ErrorResponse(c, http.StatusBadRequest, "invalid request")
		return
	}

	if errMsg, data := checkValidation(req); errMsg != "" {
		response.ErrorResponse(c, http.StatusBadRequest, errMsg, data)
		return
	}

	res, err := h.userService.Register(&req)
	if err != nil {
		handleErrorService(c, err)
		return
	}

	response.SuccessResponse(c, http.StatusCreated, "user registered successfully", res)
}

func (h *UserHandler) Login(c *gin.Context) {
	var req dto.LoginRequest

	if err := c.ShouldBindJSON(&req); err != nil {
		response.ErrorResponse(c, http.StatusBadRequest, "invalid request")
		return
	}

	if errMsg, data := checkValidation(req); errMsg != "" {
		response.ErrorResponse(c, http.StatusBadRequest, errMsg, data)
		return
	}

	res, err := h.userService.Login(&req)
	if err != nil {
		handleErrorService(c, err)
		return
	}

	response.SuccessResponse(c, http.StatusOK, "login successful", res)
}

func (h *UserHandler) Logout(c *gin.Context) {
	tokenString := c.GetHeader("Authorization")
	if strings.HasPrefix(tokenString, "Bearer ") {
		tokenString = tokenString[7:]
	} else {
		response.ErrorResponse(c, http.StatusBadRequest, "invalid authorization header format")
		return
	}

	if err := h.userService.Logout(tokenString); err != nil {
		response.ErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	response.SuccessResponse(c, http.StatusOK, "logout successful", nil)
}

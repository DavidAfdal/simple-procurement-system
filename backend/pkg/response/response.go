package response

import "github.com/gin-gonic/gin"

type ApiResponse struct {
	Success bool        `json:"success"`
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
}

func SuccessResponse(c *gin.Context, code int, message string, data interface{}) {
	c.JSON(code, ApiResponse{
		Success: true,
		Message: message,
		Data:    data,
	})
}

func ErrorResponse(c *gin.Context, code int, message string, data ...interface{}) {
	var respData interface{} = nil

	if len(data) > 0 {
		respData = data[0]
	}

	c.JSON(code, ApiResponse{
		Success: false,
		Message: message,
		Data:    respData,
	})
}

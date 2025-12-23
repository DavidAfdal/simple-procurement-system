package middelwares

import (
	"net/http"

	"github.com/DavidAfdal/purchasing-systeam/pkg/response"
	"github.com/DavidAfdal/purchasing-systeam/pkg/token"
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
)

func UserContextMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		user, exists := c.Get("user")
		if !exists {
			response.ErrorResponse(c, http.StatusUnauthorized, "anda harus login untuk mengakses resource ini")
			c.Abort()
			return
		}

		claims := user.(*jwt.Token).Claims.(*token.JwtCustomClaims)

		c.Set("user_id", claims.ID)
		c.Set("user_email", claims.Username)

		c.Next()
	}
}

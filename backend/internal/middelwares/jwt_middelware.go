package middelwares

import (
	"net/http"

	"github.com/DavidAfdal/purchasing-systeam/pkg/response"
	"github.com/DavidAfdal/purchasing-systeam/pkg/token"
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
)

func JWTProtection(secretKey string) gin.HandlerFunc {
	return func(c *gin.Context) {
		auth := c.GetHeader("Authorization")
		if auth == "" {
			response.ErrorResponse(c, http.StatusUnauthorized, "anda harus login untuk mengakses resource ini")
			c.Abort()
			return
		}

		tokenStr := auth[len("Bearer "):]

		tokenData, err := jwt.ParseWithClaims(
			tokenStr,
			&token.JwtCustomClaims{},
			func(token *jwt.Token) (interface{}, error) {
				return []byte(secretKey), nil
			},
		)

		if err != nil || !tokenData.Valid {
			response.ErrorResponse(c, http.StatusUnauthorized, "anda harus login untuk mengakses resource ini")
			c.Abort()
			return
		}

		c.Set("user", tokenData)
		c.Next()
	}
}

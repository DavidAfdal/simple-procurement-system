package middelwares

import (
	"log"
	"net/http"
	"strings"

	"github.com/DavidAfdal/purchasing-systeam/pkg/response"
	"github.com/DavidAfdal/purchasing-systeam/pkg/token"
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
)

func JWTProtection(secretKey string) gin.HandlerFunc {
	return func(c *gin.Context) {
		auth := c.GetHeader("Authorization")
		if auth == "" {
			log.Println("Authorization header is missing")
			response.ErrorResponse(c, http.StatusUnauthorized, "You must be logged in to access this resource")
			c.Abort()
			return
		}

		if !strings.HasPrefix(auth, "Bearer ") {
			log.Println("Invalid Authorization header format:", auth)
			response.ErrorResponse(c, http.StatusUnauthorized, "You must be logged in to access this resource")
			c.Abort()
			return
		}

		tokenStr := strings.TrimPrefix(auth, "Bearer ")

		tokenData, err := jwt.ParseWithClaims(
			tokenStr,
			&token.JwtCustomClaims{},
			func(token *jwt.Token) (interface{}, error) {
				return []byte(secretKey), nil
			},
		)

		if err != nil {
			log.Println("Failed to parse JWT:", err)
			response.ErrorResponse(c, http.StatusUnauthorized, "You must be logged in to access this resource")
			c.Abort()
			return
		}

		if !tokenData.Valid {
			log.Println("Invalid JWT token:", tokenStr)
			response.ErrorResponse(c, http.StatusUnauthorized, "You must be logged in to access this resource")
			c.Abort()
			return
		}

		c.Set("user", tokenData)
		c.Next()
	}
}

package middelwares

import (
	"log"
	"net/http"
	"strings"

	"github.com/DavidAfdal/purchasing-systeam/pkg/response"
	"github.com/DavidAfdal/purchasing-systeam/pkg/token"
	"github.com/gin-gonic/gin"
)

func CheckBlacklistToken(tokenUse token.TokenUseCase) gin.HandlerFunc {
	return func(c *gin.Context) {
		tokenString := c.GetHeader("Authorization")
		if tokenString == "" {
			log.Println("Authorization header is missing")
			response.ErrorResponse(c, http.StatusUnauthorized, "Unauthorized access")
			c.Abort()
			return
		}

		tokenString = strings.TrimPrefix(tokenString, "Bearer ")

		if tokenUse.IsTokenBlacklisted(tokenString) {
			log.Println("Attempt to use blacklisted token:", tokenString)
			response.ErrorResponse(c, http.StatusUnauthorized, "Unauthorized access")
			c.Abort()
			return
		}

		c.Next()
	}
}

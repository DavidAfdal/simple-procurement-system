package middelwares

import (
	"log"
	"net/http"

	"github.com/DavidAfdal/purchasing-systeam/pkg/response"
	"github.com/DavidAfdal/purchasing-systeam/pkg/token"
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
)

func RBACMiddleware(roles ...string) gin.HandlerFunc {
	return func(c *gin.Context) {
		user, exists := c.Get("user")
		if !exists {
			log.Println("RBAC check failed: user not found in context")
			response.ErrorResponse(c, http.StatusUnauthorized, "You must be logged in to access this resource")
			c.Abort()
			return
		}

		claims, ok := user.(*jwt.Token).Claims.(*token.JwtCustomClaims)
		if !ok {
			log.Println("RBAC check failed: invalid token claims")
			response.ErrorResponse(c, http.StatusUnauthorized, "Invalid token")
			c.Abort()
			return
		}

		if !contains(roles, claims.Role) {
			log.Printf("RBAC check failed: user role '%s' not authorized, required roles: %v\n", claims.Role, roles)
			response.ErrorResponse(c, http.StatusForbidden, "You do not have access to this resource")
			c.Abort()
			return
		}

		c.Next()
	}
}

func contains(slice []string, val string) bool {
	for _, item := range slice {
		if item == val {
			return true
		}
	}
	return false
}

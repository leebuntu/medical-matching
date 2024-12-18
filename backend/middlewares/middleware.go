package middlewares

import (
	"medical-matching/constants"
	"medical-matching/db/providers"

	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
)

func AuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		authHeader := c.GetHeader("Authorization")
		if authHeader == "" {
			c.JSON(http.StatusUnauthorized, gin.H{"error": constants.Unauthorized})
			c.Abort()
			return
		}

		tokenString := strings.TrimPrefix(authHeader, "Bearer ")

		claims, err := ValidateJWT(tokenString)
		if err != nil {
			c.JSON(http.StatusUnauthorized, gin.H{"error": constants.Unauthorized})
			c.Abort()
			return
		}

		ap := providers.GetAuthProvider()
		isExist, err := ap.IsExistUser(claims.UserID)
		if err != nil || !isExist {
			c.JSON(http.StatusUnauthorized, gin.H{"error": constants.Unauthorized})
			c.Abort()
			return
		}

		c.Set("userID", claims.UserID)

		c.Next()
	}
}

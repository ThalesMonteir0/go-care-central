package middleware

import (
	"github.com/ThalesMonteir0/go-care-central/internal/infra/auth"
	"github.com/gin-gonic/gin"
	"net/http"
	"strings"
)

func AuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		header := c.GetHeader("Authorization")

		if header == "" || !strings.HasPrefix(header, "Bearer ") {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "token não é valido"})
			return
		}

		token := strings.TrimPrefix(header, "Bearer ")

		clinicID, err := auth.VerifyToken(token)
		if err != nil {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "token não é valido"})
			return
		}

		c.Set("clinic_id", clinicID)
		c.Next()
	}
}

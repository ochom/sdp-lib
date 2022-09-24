package middleware

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/ochom/sdp-lib/models"
	"github.com/ochom/sdp-lib/utils"
)

// Authenticated ...
func Authenticated() gin.HandlerFunc {
	return func(c *gin.Context) {
		authHeader := c.Request.Header.Get("Authorization")
		if authHeader == "" {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "no auth header"})
			return
		}

		token := authHeader[7:]
		if token == "" {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "no token"})
			return
		}

		claim, err := utils.ValidateToken(token)
		if err != nil {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": err.Error()})
			return
		}

		user := models.User{
			ID:             claim.ID,
			Mobile:         claim.Mobile,
			Email:          claim.Email,
			OrganizationID: claim.Organization.ID,
		}

		authContext := utils.AuthenticatedContext(c, &user)
		c.Request = c.Request.WithContext(authContext)

		c.Next()
	}
}

package middlewares

import (
	"net/http"

	"github.com/gin-gonic/contrib/sessions"
	"github.com/gin-gonic/gin"
)

// GuestRequired
func GuestRequired() gin.HandlerFunc {
	return func(c *gin.Context) {
		session := sessions.Default(c)
		user := session.Get("user")
		if user != nil {
			c.JSON(http.StatusForbidden, gin.H{"error": "You are forbidden"})
		} else {
			c.Next()
		}
	}
}

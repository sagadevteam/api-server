package middlewares

import (
	"net/http"

	"github.com/gin-gonic/contrib/sessions"
	"github.com/gin-gonic/gin"
)

// AuthRequired
func AuthRequired() gin.HandlerFunc {
	return func(c *gin.Context) {
		session := sessions.Default(c)
		user := session.Get("user")
		if user != nil {
			c.Next()
		} else {
			c.JSON(http.StatusForbidden, gin.H{"err": "You are not authenticated"})
			c.Abort()
			return
		}
	}
}

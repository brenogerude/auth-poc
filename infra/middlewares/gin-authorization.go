package middleware

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func Authorization() gin.HandlerFunc {
	return func(c *gin.Context) {
		authorized, err := authorization(c.GetHeader("Authorization"), c.Request)

		if !authorized {
			c.JSON(http.StatusUnauthorized, gin.H{
				"error": err,
			})
			c.Abort()
			return
		}
	}
}

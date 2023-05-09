package middleware

import (
	"log"

	"github.com/gin-gonic/gin"
)

func Authentication() gin.HandlerFunc {
	return func(c *gin.Context) {
		log.Println("authenticated")
	}
}

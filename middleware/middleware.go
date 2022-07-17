package middleware

import (
	"shape-api/auth"

	"github.com/gin-gonic/gin"
)

var authCheck = new(auth.Auth)

func AuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		authCheck.TokenValid(c)
		c.Next()
	}
}

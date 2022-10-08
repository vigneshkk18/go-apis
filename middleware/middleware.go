package middleware

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/vigneshkk18/go-apis/initializers"
)

func AuthRequired() gin.HandlerFunc {
	return func(c *gin.Context) {
		secretKey := c.Request.Header.Get("SECRET_KEY")
		if secretKey != initializers.SECRET_KEY {
			c.JSON(http.StatusForbidden, gin.H{
				"error": "Secret key is required",
			})
			c.Abort()
			return
		}
	}
}

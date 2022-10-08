package middleware

import (
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/vigneshkk18/go-apis/initializers"
)

func CORSMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {

		c.Header("Access-Control-Allow-Origin", "*")
		c.Header("Access-Control-Allow-Credentials", "true")
		c.Header("Access-Control-Allow-Headers", "Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization, accept, origin, Cache-Control, X-Requested-With")
		c.Header("Access-Control-Allow-Methods", "POST,HEAD,PATCH, OPTIONS, GET, PUT")

		if c.Request.Method == "OPTIONS" {
			c.JSON(http.StatusOK, gin.H{})
			return
		}

		c.Next()
	}
}

func AuthRequired() gin.HandlerFunc {
	return func(c *gin.Context) {
		if c.Request.Header["Authorization"] == nil {
			c.AbortWithStatusJSON(http.StatusForbidden, gin.H{
				"error": "Token Expected",
			})
			return
		}

		secretKey := strings.Split(c.Request.Header.Get("Authorization"), "Bearer ")[1]
		if secretKey != initializers.SECRET_KEY {
			c.JSON(http.StatusForbidden, gin.H{
				"error": "Secret key is required",
			})
			c.Abort()
			return
		}
		c.Next()
	}
}

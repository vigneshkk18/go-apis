package routes

import (
	"github.com/gin-gonic/gin"
)

func CORSMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {

		c.Header("Access-Control-Allow-Origin", "*")
		c.Header("Access-Control-Allow-Credentials", "true")
		c.Header("Access-Control-Allow-Headers", "Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization, accept, origin, Cache-Control, X-Requested-With, SECRET_KEY")
		c.Header("Access-Control-Allow-Methods", "POST,HEAD,PATCH, OPTIONS, GET, PUT")

		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(204)
			return
		}

		c.Next()
	}
}

func InitializeRoutes(r *gin.Engine) {
	r.Use(CORSMiddleware())

	// Non - Project Specific Routes
	gpGroup := r.Group("/generate-paragraph")
	{
		gpGroup.GET("/difficulty-options", GetDifficultyOptions)
		gpGroup.GET("/:difficulty", GetRandomParagraph)
	}

	// Typing Practice route
	tpGroup := r.Group("/typing-practice")
	{
		tpGroup.GET("/user-activity/:emailId", GetUserActivity)
		tpGroup.POST("/user-activity", CreateUserActivity)
	}
}

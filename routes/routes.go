package routes

import "github.com/gin-gonic/gin"

func InitializeRoutes(r *gin.Engine) {
	// Non - Project Specific Routes
	gpGroup := r.Group("/generate-paragraph")
	{
		gpGroup.GET("/difficulty-options", GetDifficultyOptions)
		gpGroup.GET("/:difficulty", GetRandomParagraph)
	}

	// Typing Practice route
	tpGroup := r.Group("/typing-practice")
	{
		tpGroup.GET("/user-activity/:emailId/stats", GetUserActivity)
		tpGroup.POST("/user-activity", CreateUserActivity)
	}
}

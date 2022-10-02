package routes

import (
	"github.com/gin-gonic/gin"
)

func InitializeRoutes(r *gin.Engine) {
	// Non - Project Specific Routes
	r.GET("/generate-paragraph/:difficulty", GetRandomParagraph)

	// Typing Practice route
	tpGroup := r.Group("/typing-practice")
	{
		tpGroup.GET("/user-activity/:emailId/stats", GetUserActivity)
		tpGroup.POST("/user-activity", CreateUserActivity)
	}
}

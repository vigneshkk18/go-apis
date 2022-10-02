package routes

import (
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/vigneshkk18/go-apis/controllers"
	"github.com/vigneshkk18/go-apis/models"
)

func CreateUserActivity(c *gin.Context) {
	// Get User Activity data from body
	var userActivity models.UserActivityRecord
	_ = c.BindJSON(&userActivity)
	userActivity.CompletedAt = time.Now().UTC().String()

	// Check if user activity is valid
	if isValid := userActivity.IsValid(); !isValid {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Some of the Required fields are missing or invalid..",
		})
		return
	}

	// insert user activity record
	if err := controllers.TP_CreateUserActivity(userActivity); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err,
		})
		return
	}

	c.JSON(http.StatusCreated, gin.H{
		"message": "User Activity Saved Successfully",
	})
}

func GetUserActivity(c *gin.Context) {
	// get emailid from params
	emailId := c.Param("emailId")

	if emailId == "" {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Invalid emailId",
		})
		return
	}

	// get user activity from received emailid
	userActivities, err := controllers.TP_GetUserActivity(emailId)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err,
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"data": userActivities,
	})
}

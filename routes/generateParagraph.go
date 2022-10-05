package routes

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/vigneshkk18/go-apis/controllers"
	"github.com/vigneshkk18/go-apis/utils"
)

func GetDifficultyOptions(c *gin.Context) {
	difficultyOptions, _ := controllers.GetDifficultyOptions()
	c.JSON(http.StatusOK, gin.H{
		"data": difficultyOptions,
	})
}

func GetRandomParagraph(c *gin.Context) {
	// get difficulty level from params
	difficulty := c.Param("difficulty")

	val, ok := utils.DifficultyMap[difficulty]
	if !ok {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Invalid difficulty",
		})
		return
	}

	// Get paragraph according to difficulty level
	para, err := controllers.GetRandomParagraph(val)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err,
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": para})
}

package routes

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/vigneshkk18/go-apis/controllers"
)

var difficultyMap = map[string]uint{
	"easy":   10,
	"medium": 15,
	"hard":   20,
	"expert": 25,
}

func GetRandomParagraph(c *gin.Context) {
	// get difficulty level from params
	difficulty := c.Param("difficulty")

	val, ok := difficultyMap[difficulty]
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

	c.JSON(http.StatusAccepted, gin.H{"data": para})
}

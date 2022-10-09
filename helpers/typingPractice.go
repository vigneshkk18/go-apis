package helpers

import (
	"github.com/vigneshkk18/go-apis/models"
	"github.com/vigneshkk18/go-apis/utils"
)

func TransformActivity(userActivities []models.UserActivityRecord) []models.Activity {
	activities := []models.Activity{}

	for _, userActivity := range userActivities {
		timeTaken := utils.TimeTaken(userActivity.CompletedIn)
		activity := models.Activity{Accuracy: userActivity.Accuracy, WPM: userActivity.WPM, Difficulty: userActivity.Difficulty, TimeTaken: timeTaken}
		activities = append(activities, activity)
	}
	return activities
}

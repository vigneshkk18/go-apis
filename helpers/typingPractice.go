package helpers

import (
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/vigneshkk18/go-apis/models"
	"github.com/vigneshkk18/go-apis/utils"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func GetUserActivityQueryParams(c *gin.Context) (int, int, int, int) {
	var queries = [4]string{"year", "month", "week", "day"}
	var values = [4]int{}

	for i := 0; i < len(queries); i++ {
		if val := c.Query(queries[i]); val != "" {
			value, _ := strconv.Atoi(val)
			values[i] = value
		} else {
			values[i] = -1
		}
	}

	return values[0], values[1], values[2], values[3]
}

func GetUserActivityFilterQuery(email string, year int, month int, week int, day int, groupBy string) primitive.M {
	filter := bson.M{"email": email}

	if groupBy == "ALL" {
		return filter
	}

	var (
		startDate time.Time
		endDate   time.Time
	)

	if groupBy == "GROUP_BY_DAY" {
		startDate = utils.Date(year, month, day)
		endDate = startDate.AddDate(0, 0, 1)
	} else if groupBy == "GROUP_BY_WEEK" {
		startDate = utils.Date(year, month, utils.GetStartDateOfWeek(week))
		endDate = startDate.AddDate(0, 0, 7)
	} else if groupBy == "GROUP_BY_MONTH" {
		startDate = utils.Date(year, month, 1)
		endDate = startDate.AddDate(0, 1, 0)
	} else if groupBy == "GROUP_BY_YEAR" {
		startDate = utils.Date(year, 1, 1)
		endDate = startDate.AddDate(1, 0, 0)
	}

	parsedStartDate := primitive.NewDateTimeFromTime(startDate)
	parsedEndDate := primitive.NewDateTimeFromTime(endDate)

	filter = bson.M{
		"email": email,
		"completed_at": bson.M{
			"$gt": parsedStartDate,
			"$lt": parsedEndDate,
		},
	}

	return filter
}

// group activities by year -> month -> week -> day -> activities
func GroupActivities(userActivities []models.UserActivityRecord, groupBy string) models.GroupedUserActivityMap {
	if groupBy == "GROUP_BY_YEAR" {
		return GroupActivitiesByMonth(userActivities)
	} else if groupBy == "GROUP_BY_MONTH" {
		return GroupActivitiesByWeek(userActivities)
	} else if groupBy == "GROUP_BY_WEEK" {
		return GroupActivitiesByDay(userActivities)
	} else if groupBy == "GROUP_BY_DAY" {
		return GroupActivitiesByIndex(userActivities)
	}

	return GroupActivitiesByYear(userActivities)
}

func GroupActivitiesByYear(userActivities []models.UserActivityRecord) models.GroupedUserActivityMap {
	groupedActivities := models.GroupedUserActivityMap{}

	for _, activity := range userActivities {
		userActivity := models.Activity{CompletedIn: utils.TimeTaken(activity.CompletedIn), Accuracy: activity.Accuracy, WPM: activity.Accuracy, Difficulty: activity.Difficulty}

		year := activity.CompletedAt.Time().Year()
		if activityByYear, ok := groupedActivities[year]; ok {
			activityByYear.Hit = activityByYear.Hit + 1
			activityByYear.Activities = append(activityByYear.Activities, userActivity)
			groupedActivities[year] = activityByYear
		} else {
			groupedActivities[year] = models.GroupedUserActivity{Value: year, Hit: 1, Activities: []models.Activity{userActivity}}
		}
	}

	return groupedActivities
}

func GroupActivitiesByMonth(userActivities []models.UserActivityRecord) models.GroupedUserActivityMap {
	groupedActivities := models.GroupedUserActivityMap{}

	for _, activity := range userActivities {
		userActivity := models.Activity{CompletedIn: utils.TimeTaken(activity.CompletedIn), Accuracy: activity.Accuracy, WPM: activity.Accuracy, Difficulty: activity.Difficulty}

		month := utils.MonthInt[activity.CompletedAt.Time().Month().String()]
		if activityByMonth, ok := groupedActivities[month]; ok {
			activityByMonth.Hit = activityByMonth.Hit + 1
			activityByMonth.Activities = append(activityByMonth.Activities, userActivity)
			groupedActivities[month] = activityByMonth
		} else {
			groupedActivities[month] = models.GroupedUserActivity{Value: month, Hit: 1, Activities: []models.Activity{userActivity}}
		}
	}

	return groupedActivities
}

func GroupActivitiesByWeek(userActivities []models.UserActivityRecord) models.GroupedUserActivityMap {
	groupedActivities := models.GroupedUserActivityMap{}

	for _, activity := range userActivities {
		userActivity := models.Activity{CompletedIn: utils.TimeTaken(activity.CompletedIn), Accuracy: activity.Accuracy, WPM: activity.Accuracy, Difficulty: activity.Difficulty}

		week := activity.CompletedAt.Time().Day() % 7
		if activityByWeek, ok := groupedActivities[week]; ok {
			activityByWeek.Hit = activityByWeek.Hit + 1
			activityByWeek.Activities = append(activityByWeek.Activities, userActivity)
			groupedActivities[week] = activityByWeek
		} else {
			groupedActivities[week] = models.GroupedUserActivity{Value: week, Hit: 1, Activities: []models.Activity{userActivity}}
		}
	}

	return groupedActivities
}

func GroupActivitiesByDay(userActivities []models.UserActivityRecord) models.GroupedUserActivityMap {
	groupedActivities := models.GroupedUserActivityMap{}

	for _, activity := range userActivities {
		userActivity := models.Activity{CompletedIn: utils.TimeTaken(activity.CompletedIn), Accuracy: activity.Accuracy, WPM: activity.Accuracy, Difficulty: activity.Difficulty}

		day := activity.CompletedAt.Time().Day()
		if activityByDay, ok := groupedActivities[day]; ok {
			activityByDay.Hit = activityByDay.Hit + 1
			activityByDay.Activities = append(activityByDay.Activities, userActivity)
			groupedActivities[day] = activityByDay
		} else {
			groupedActivities[day] = models.GroupedUserActivity{Value: day, Hit: 1, Activities: []models.Activity{userActivity}}
		}
	}

	return groupedActivities
}

func GroupActivitiesByIndex(userActivities []models.UserActivityRecord) models.GroupedUserActivityMap {
	groupedActivities := models.GroupedUserActivityMap{}

	for idx, activity := range userActivities {
		userActivity := models.Activity{CompletedIn: utils.TimeTaken(activity.CompletedIn), Accuracy: activity.Accuracy, WPM: activity.Accuracy, Difficulty: activity.Difficulty}

		if activityByIdx, ok := groupedActivities[idx]; ok {
			activityByIdx.Hit = activityByIdx.Hit + 1
			activityByIdx.Activities = append(activityByIdx.Activities, userActivity)
			groupedActivities[idx] = activityByIdx
		} else {
			groupedActivities[idx] = models.GroupedUserActivity{Value: idx, Hit: 1, Activities: []models.Activity{userActivity}}
		}
	}

	return groupedActivities
}

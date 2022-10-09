package helpers

import (
	"time"

	"github.com/vigneshkk18/go-apis/models"
	"github.com/vigneshkk18/go-apis/utils"
)

// initialize all days in a week
func initializeDaysInWeek(week int) map[int]models.Day {
	days := map[int]models.Day{}

	for i := 1; i <= 7; i++ {
		currDay := (week * 7) + i
		days[currDay] = models.Day{Day: currDay, Hit: 0, Activities: []models.Activity{}}
	}

	return days
}

// initialize all weeks in a month, internally calls initializeDaysInWeek.
func initializeWeeksInMonth() map[int]models.Week {
	weeks := map[int]models.Week{}

	for i := 0; i < 4; i++ {
		weeks[i] = models.Week{Week: i, Hit: 0, Days: initializeDaysInWeek(i)}
	}

	return weeks
}

// initialize all months in a year, internally calls initializeWeeksInMonth.
func initializeMonthsInYear() map[string]models.Month {
	months := map[string]models.Month{}

	for i := 1; i <= 12; i++ {
		monthName := time.Month(i).String()
		months[monthName] = models.Month{Month: monthName, Hit: 0, Weeks: initializeWeeksInMonth()}
	}

	return months
}

// group activities by year -> month -> week -> day -> activities
func GroupActivities(userActivities []models.UserActivityRecord) models.GroupedUserActivity {
	groupedUserActivities := models.GroupedUserActivity{}

	for _, userActivity := range userActivities {
		Time := userActivity.CompletedAt.Time()
		year := Time.Year()
		month := Time.Month().String()
		day := Time.Day()
		week := day % 7
		timeTaken := utils.TimeTaken(userActivity.CompletedIn)
		activity := models.Activity{TimeTaken: timeTaken, WPM: userActivity.WPM, Accuracy: userActivity.Accuracy, Difficulty: userActivity.Difficulty}

		activityByYear, ok := groupedUserActivities[year]
		if ok {
			activityByYear.Hit = activityByYear.Hit + 1

			activityByMonth := activityByYear.Months[month]
			activityByWeek := activityByMonth.Weeks[week]
			activityByDay := activityByWeek.Days[day]

			// increment user activity in day, activities
			activityByDay.Hit = activityByDay.Hit + 1
			activityByDay.Activities = append(activityByDay.Activities, activity)
			activityByWeek.Days[day] = activityByDay

			// increment user activity in week, month.
			activityByWeek.Hit = activityByWeek.Hit + 1
			activityByMonth.Weeks[week] = activityByWeek

			activityByMonth.Hit = activityByMonth.Hit + 1
			activityByYear.Months[month] = activityByMonth
		} else {
			// initialize months
			months := initializeMonthsInYear()
			monthObj := months[month]
			weekObj := monthObj.Weeks[week]
			dayObj := weekObj.Days[day]

			// increment user activity in day, activities
			dayObj.Hit = dayObj.Hit + 1
			dayObj.Activities = append(dayObj.Activities, activity)
			weekObj.Days[day] = dayObj // set updated day -> week

			monthObj.Hit = monthObj.Hit + 1 // increment user activity in month
			monthObj.Weeks[week] = weekObj  // set updated week -> month

			activityByYear = models.Year{Year: year, Hit: 1, Months: months}
		}

		groupedUserActivities[year] = activityByYear
	}

	return groupedUserActivities
}

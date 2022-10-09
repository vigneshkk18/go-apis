package models

import (
	"reflect"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type UserActivityRecord struct {
	Email       string             `json:"email" bson:"email"`
	CompletedIn string             `json:"completedIn" bson:"completed_in"`
	CompletedAt primitive.DateTime `json:"completedAt" bson:"completed_at"`
	Accuracy    float64            `json:"accuracy" bson:"accuracy"`
	WPM         float64            `json:"wpm" bson:"wpm"`
	Difficulty  string             `json:"difficulty" bson:"difficulty"`
}

func (u UserActivityRecord) IsValid() bool {
	v := reflect.ValueOf(u)
	isValid := true
	for i := 0; i < v.NumField(); i++ {
		if v.Field(i).Type().Name() == "ObjectID" {
			continue
		}

		if v.Field(i).IsZero() {
			isValid = false
			break
		}
	}
	return isValid
}

type Year struct {
	Year   int              `json:"year" bson:"year"`
	Hit    int              `json:"hit" bson:"hit"`
	Months map[string]Month `json:"months" bson:"months"`
}

type Month struct {
	Month string       `json:"month" bson:"month"`
	Hit   int          `json:"hit" bson:"hit"`
	Weeks map[int]Week `json:"weeks" bson:"weeks"`
}

type Week struct {
	Week int         `json:"week" bson:"week"`
	Hit  int         `json:"hit" bson:"hit"`
	Days map[int]Day `json:"days" bson:"days"`
}

type Day struct {
	Day        int        `json:"day" bson:"day"`
	Hit        int        `json:"hit" bson:"hit"`
	Activities []Activity `json:"activities" bson:"activities"`
}

type Activity struct {
	TimeTaken  float64 `json:"timeTaken" bson:"timeTaken"`
	WPM        float64 `json:"wpm" bson:"wpm"`
	Accuracy   float64 `json:"accuracy" bson:"accuracy"`
	Difficulty string  `json:"difficulty" bson:"difficulty"`
}

type GroupedUserActivity map[int]Year

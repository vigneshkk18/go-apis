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

type Activity struct {
	TimeTaken  float64 `json:"timeTaken" bson:"completed_in"`
	Accuracy   float64 `json:"accuracy" bson:"accuracy"`
	WPM        float64 `json:"wpm" bson:"wpm"`
	Difficulty string  `json:"difficulty" bson:"difficulty"`
}

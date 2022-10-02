package controllers

import (
	"context"
	"errors"

	"github.com/vigneshkk18/go-apis/initializers"
	"github.com/vigneshkk18/go-apis/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

func getCollection() *mongo.Collection {
	DB := initializers.Client.Database("typing-practice")
	return DB.Collection("user-activity")
}

func TP_CreateUserActivity(userActivity models.UserActivityRecord) error {
	collection := getCollection()

	_, err := collection.InsertOne(context.Background(), userActivity)
	return err
}

func TP_GetUserActivity(emailId string) ([]models.UserActivityRecord, error) {
	collection := getCollection()
	filter := bson.M{"email": emailId}

	// Get all user activity records
	cursor, err := collection.Find(context.Background(), filter)
	if err != nil {
		return nil, errors.New("unable to find user activity")
	}

	var userActivities = []models.UserActivityRecord{}
	// Construct user activity records from the DB cursor result
	for cursor.Next(context.Background()) {
		var userActivity models.UserActivityRecord
		_ = cursor.Decode(&userActivity)
		userActivities = append(userActivities, userActivity)
	}

	return userActivities, nil
}

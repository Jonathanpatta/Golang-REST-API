package auth

import (
	"context"
	"task1/database"
	"time"

	"go.mongodb.org/mongo-driver/bson"
)

func Authenticate(token string) bool {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	collection := database.GetCollection("AuthToken")

	result := collection.FindOne(ctx, bson.M{"Token": token})

	if result != nil {
		return true
	}
	return false
}

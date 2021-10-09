package auth

import (
	"context"
	"fmt"
	"task1/database"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"golang.org/x/crypto/bcrypt"
)

func CreateToken(username string, password string) string {
	var token = ""
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	collection := database.GetCollection("Users")
	hashedPassword, _ := bcrypt.GenerateFromPassword([]byte(password), 14)
	result := collection.FindOne(ctx, bson.M{"Password": string(hashedPassword)})

	if result != nil {
		ctx_, cancel := context.WithTimeout(context.Background(), 10*time.Second)
		defer cancel()
		collection := database.GetCollection("AuthToken")

		token = primitive.NewObjectIDFromTimestamp(time.Now()).Hex()
		result, tokenInsertionError := collection.InsertOne(ctx_, bson.D{
			{Key: "Id", Value: primitive.NewObjectIDFromTimestamp(time.Now()).Hex()},
			{Key: "Token", Value: token},
		})

		if tokenInsertionError != nil {
			fmt.Println(result, tokenInsertionError)
		}
	}

	return token
}

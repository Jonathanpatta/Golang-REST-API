package database

import (
	"context"
	"log"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func GetClient() *mongo.Client {
	clientOptions := options.Client().
		ApplyURI(`mongodb+srv://jonathan:Jonu%40123@cluster0.ucwar.mongodb.net/instagram-db?retryWrites=true&w=majority`)

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	client, err := mongo.Connect(ctx, clientOptions)

	if err != nil {
		log.Fatal(err)
	}
	return client
}

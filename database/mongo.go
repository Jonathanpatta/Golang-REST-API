package database

import (
	"context"
	"log"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var Client *mongo.Client

var dbname = "instagram-db"

func CreateClient() {
	clientOptions := options.Client().
		ApplyURI(`mongodb+srv://jonathan:Jonu%40123@cluster0.ucwar.mongodb.net/instagram-db?retryWrites=true&w=majority`)

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	client, err := mongo.Connect(ctx, clientOptions)
	client.Database("instagram-db").Collection("")

	Client = client

	if err != nil {
		log.Fatal(err)
	}
}

func GetCollection(collection string) *mongo.Collection {
	var col = Client.Database(dbname).Collection(collection)

	return col
}

func GetClient() *mongo.Client {

	return Client
}

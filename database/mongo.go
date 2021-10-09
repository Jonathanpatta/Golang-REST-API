package database

import (
	"context"
	"fmt"
	"log"
	"task1/config"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"net/url"
)

var Client *mongo.Client

var dbname = "instagram-db"

func CreateClient() {

	url.QueryEscape(config.DB_USERNAME)

	var host = config.DB_HOST
	var username = url.QueryEscape(config.DB_USERNAME)
	var password = url.QueryEscape(config.DB_PASSWORD)

	var dbUrl = `mongodb+srv://` + username + `:` + password + `@` + host + `?retryWrites=true&w=majority`

	fmt.Println(host, password, username)
	clientOptions := options.Client().
		ApplyURI(dbUrl)

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

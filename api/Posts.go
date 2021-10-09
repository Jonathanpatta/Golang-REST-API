package api

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"strings"
	"task1/database"
	"task1/structures"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	// "go.mongodb.org/mongo-driver/bson/primitive"
)

func HandlePosts(res http.ResponseWriter, req *http.Request) {

	var url = req.URL.String()
	var splitUrl = strings.Split(url, "/")

	if len(splitUrl) == 3 {
		if req.Method == "POST" {
			CreatePost(res, req)
			return
		}
	}
	if len(splitUrl) == 3 {
		var postId = splitUrl[2]
		if req.Method == "GET" {
			fmt.Println(postId)
			GetPost(res, req, postId)

			return
		}
	}

	if len(splitUrl) == 4 {
		if splitUrl[2] == "user" {
			var userId = splitUrl[3]
			if req.Method == "GET" {
				GetUserPosts(res, req, userId)
				return
			}
		}
	}

	http.Error(res, "invalid Url", http.StatusNotFound)
}

func GetUserPosts(res http.ResponseWriter, req *http.Request, userId string) {

	var users []bson.M

	id, idError := primitive.ObjectIDFromHex(userId)

	if idError != nil {
		http.Error(res, "Invalid Id", http.StatusBadRequest)
		return
	}

	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	collection := database.GetCollection("Posts")
	cursor, dbError := collection.Find(ctx, bson.M{"UserId": id})

	fmt.Println(userId)

	// decodingError := cursor.Decode(&users)

	for cursor.Next(ctx) {
		var user bson.M
		decodingError := cursor.Decode(&user)

		users = append(users, user)

		if decodingError != nil {
			fmt.Println(decodingError)
			http.Error(res, "couldn't get user", http.StatusBadGateway)
			return
		}

		fmt.Println(user)
	}

	if dbError != nil {
		fmt.Println(dbError)
		http.Error(res, "couldn't get user due to db error", http.StatusBadGateway)
		return
	}

	json.NewEncoder(res).Encode(users)
}

func CreatePost(res http.ResponseWriter, req *http.Request) {

	//var client = database.GetClient()
	//fmt.Println(client)

	//fmt.Println(req.Body)

	// id, error := primitive.ObjectIDFromHex(userId)
	// fmt.Print(id)
	// if error != nil {
	// 	http.Error(res, "invalid user Id", http.StatusNotFound)
	// 	return
	// }

	var post structures.Post

	var decoder = json.NewDecoder(req.Body)
	decoder.DisallowUnknownFields()
	decoder.Decode(&post)

	req.ParseForm()
	caption := req.PostForm.Get("Caption")
	fmt.Println("caption:", caption)

	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	collection := database.GetCollection("Posts")

	post.Id = primitive.NewObjectID()
	post.PostedTimeStamp = primitive.NewDateTimeFromTime(time.Now())
	result, error := collection.InsertOne(ctx, post)

	fmt.Println(result)
	fmt.Println(error)
	fmt.Print("create Post")

	fmt.Println(post)

	json.NewEncoder(res).Encode(post)

}

func GetPost(res http.ResponseWriter, req *http.Request, Id string) {

	var post structures.Post

	id, idError := primitive.ObjectIDFromHex(Id)

	if idError != nil {
		http.Error(res, "Invalid Id", http.StatusBadRequest)
		return
	}

	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	collection := database.GetCollection("Posts")
	result := collection.FindOne(ctx, bson.M{"_id": id})

	decodingError := result.Decode(&post)

	if decodingError != nil {
		http.Error(res, "couldn't get Post", http.StatusBadGateway)
		return
	}

	json.NewEncoder(res).Encode(post)

}

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
		var userId = splitUrl[2]
		if req.Method == "GET" {
			fmt.Println(userId)
			GetPost(res, req, userId)

			return
		}
	}

	http.Error(res, "invalid Url", http.StatusNotFound)
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

	fmt.Print("get post")

	json.NewEncoder(res).Encode(post)

}

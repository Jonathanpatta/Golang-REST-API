package api

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
	"strings"
	"task1/database"
	"task1/structures"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"task1/auth"
)

//Post Controller

func HandlePosts(res http.ResponseWriter, req *http.Request) {

	var url = req.URL.String()
	var pureurl = strings.Split(url, "?")[0]
	var splitUrl = strings.Split(pureurl, "/")

	var tokenString = req.URL.Query().Get("token")

	if !auth.Authenticate(tokenString) {
		http.Error(res, "invalid api auth token", http.StatusMethodNotAllowed)
	}
	//Manual Routing

	// /posts/ method="POST"
	if len(splitUrl) == 3 {
		if req.Method == "POST" {
			CreatePost(res, req)
			return
		}
	}

	// /posts/postid method="get"
	if len(splitUrl) == 3 {
		var postId = splitUrl[2]
		if req.Method == "GET" {
			fmt.Println(postId)
			GetPost(res, req, postId)

			return
		}
	}

	// /posts/user/userid method="get"
	if len(splitUrl) == 4 {
		if splitUrl[2] == "user" {
			var userId = splitUrl[3]
			if req.Method == "GET" {
				GetUserPosts(res, req, userId)
				return
			}
		}
	}

	//if condition is not met then url returns invalid url
	http.Error(res, "invalid Url", http.StatusNotFound)
}

func paginate(x []bson.M, skip int, size int) []bson.M {

	// pagination function for list of objects
	limit := func() int {
		if skip+size > len(x) {
			return len(x)
		} else {
			return skip + size
		}

	}

	start := func() int {
		if skip > len(x) {
			return len(x)
		} else {
			return skip
		}

	}
	return x[start():limit()]
}

func GetUserPosts(res http.ResponseWriter, req *http.Request, userId string) {

	// /posts/user/id handler

	var users []bson.M

	id, idError := primitive.ObjectIDFromHex(userId)

	//invalid Id
	if idError != nil {
		http.Error(res, "Invalid Id", http.StatusBadRequest)
		return
	}

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	collection := database.GetCollection("Posts")
	cursor, dbError := collection.Find(ctx, bson.M{"UserId": id})

	//Database Error handling

	if dbError != nil {
		fmt.Println(dbError)
		http.Error(res, "Couldn't get user due to db error", http.StatusBadGateway)
		return
	}

	//loop through every element and decode it
	for cursor.Next(ctx) {
		var user bson.M
		decodingError := cursor.Decode(&user)

		users = append(users, user)

		if decodingError != nil {
			fmt.Println(decodingError)
			http.Error(res, "Couldn't get user", http.StatusBadGateway)
			return
		}
	}

	//parsing the page Query String parameter to get the page number

	var pageString = req.URL.Query().Get("page")
	var page = 0

	page, _ = strconv.Atoi(pageString)

	users = paginate(users, page, 2)

	json.NewEncoder(res).Encode(users)
}

func CreatePost(res http.ResponseWriter, req *http.Request) {

	//Create Post Handler

	var post structures.Post

	var decoder = json.NewDecoder(req.Body)
	decoder.DisallowUnknownFields()
	decoder.Decode(&post)

	req.ParseForm()
	caption := req.PostForm.Get("Caption")
	fmt.Println("caption:", caption)

	//Getting Collection from database
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	collection := database.GetCollection("Posts")

	//Assigning a newobject ID
	post.Id = primitive.NewObjectID()
	post.PostedTimeStamp = primitive.NewDateTimeFromTime(time.Now())
	result, postInsertionError := collection.InsertOne(ctx, post)

	//insertion error handling
	if postInsertionError != nil {
		fmt.Println(result)
		http.Error(res, "Couldn't insert Post", http.StatusBadGateway)

	}

	json.NewEncoder(res).Encode(post)

}

func GetPost(res http.ResponseWriter, req *http.Request, Id string) {

	// /posts/id Handler

	var post structures.Post

	id, idError := primitive.ObjectIDFromHex(Id)

	if idError != nil {
		http.Error(res, "Invalid Id", http.StatusBadRequest)
		return
	}

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	collection := database.GetCollection("Posts")
	result := collection.FindOne(ctx, bson.M{"_id": id})

	decodingError := result.Decode(&post)

	if decodingError != nil {
		http.Error(res, "couldn't get Post", http.StatusBadGateway)
		return
	}

	json.NewEncoder(res).Encode(post)

}

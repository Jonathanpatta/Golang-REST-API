package api

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"strings"
	"time"

	"net/http"

	"task1/database"
	"task1/structures"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"golang.org/x/crypto/bcrypt"
	"task1/auth"
)

//Users Controller
func HandleUsers(res http.ResponseWriter, req *http.Request) {
	var url = req.URL.String()
	var splitUrl = strings.Split(url, "/")

	var tokenString = req.URL.Query().Get("token")

	if !auth.Authenticate(tokenString) {
		http.Error(res, "invalid api auth token", http.StatusMethodNotAllowed)
	}

	// /users method = "POST"
	if len(splitUrl) == 3 {
		if req.Method == "POST" {
			CreateUser(res, req)
			return
		}
	}
	// /users/userId method = "GET"
	if len(splitUrl) == 3 {
		var userId = splitUrl[2]
		if req.Method == "GET" {
			GetUser(res, req, userId)
			return
		}
	}

	http.Error(res, "invalid Url", http.StatusNotFound)

}

func CreateUser(res http.ResponseWriter, req *http.Request) {

	//Create User Handler

	var user structures.User

	var decoder = json.NewDecoder(req.Body)
	decoder.DisallowUnknownFields()
	decoder.Decode(&user)

	//Creating New ObjectId for User
	user.Id = primitive.NewObjectID()

	var plainTextPassword = user.Password
	//Hashing password
	hashedPassword, hashingError := bcrypt.GenerateFromPassword([]byte(plainTextPassword), 14)

	if hashingError != nil {
		http.Error(res, "hasing error", http.StatusBadGateway)
		return
	}

	user.Password = string(hashedPassword)

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	collection := database.GetCollection("Users")

	result, insertError := collection.InsertOne(ctx, user)
	if insertError != nil {
		fmt.Println(result)
		http.Error(res, "insert error", http.StatusBadGateway)
		log.Fatal(insertError)
		return
	}

	json.NewEncoder(res).Encode(user)

}

func GetUser(res http.ResponseWriter, req *http.Request, Id string) {

	var user structures.User

	id, idError := primitive.ObjectIDFromHex(Id)

	if idError != nil {
		http.Error(res, "Invalid Id", http.StatusBadRequest)
		return
	}

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	collection := database.GetCollection("Users")
	result := collection.FindOne(ctx, bson.M{"_id": id})

	decodingError := result.Decode(&user)

	if decodingError != nil {
		http.Error(res, "couldn't get user", http.StatusBadGateway)
		return
	}

	//omitting password when displaying users

	user.Password = ""

	json.NewEncoder(res).Encode(user)

}

package api

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"strings"
	"time"

	//"fmt"
	"net/http"

	"task1/database"
	"task1/structures"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"golang.org/x/crypto/bcrypt"
)

func HandleUsers(res http.ResponseWriter, req *http.Request) {
	//var id = primitive.NewObjectID()
	//var user = structures.User{Id: id, Name: "Jonathan", Email: "jonathan@gmail.com", Password: "lsajfsdafjl"}
	var url = req.URL.String()
	var splitUrl = strings.Split(url, "/")

	fmt.Println(splitUrl)

	fmt.Println(len(splitUrl))

	fmt.Println(req.Method)

	if len(splitUrl) == 3 {
		if req.Method == "POST" {
			CreateUser(res, req)
			return
		}
	}
	if len(splitUrl) == 3 {
		var userId = splitUrl[2]
		if req.Method == "GET" {
			fmt.Println(userId)
			GetUser(res, req, userId)

			return
		}
	}

	http.Error(res, "invalid Url", http.StatusNotFound)

}

func CreateUser(res http.ResponseWriter, req *http.Request) {

	var user structures.User

	var decoder = json.NewDecoder(req.Body)
	decoder.DisallowUnknownFields()
	decoder.Decode(&user)

	user.Id = primitive.NewObjectID()

	var plainTextPassword = user.Password
	hashedPassword, hashingError := bcrypt.GenerateFromPassword([]byte(plainTextPassword), 14)

	if hashingError != nil {
		http.Error(res, "hasing error", http.StatusBadGateway)
		return
	}

	user.Password = string(hashedPassword)

	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
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

	fmt.Print("get user")

	var user structures.User

	id, idError := primitive.ObjectIDFromHex(Id)

	if idError != nil {
		http.Error(res, "Invalid Id", http.StatusBadRequest)
		return
	}

	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	collection := database.GetCollection("Users")
	result := collection.FindOne(ctx, bson.M{"_id": id})

	decodingError := result.Decode(&user)

	if decodingError != nil {
		http.Error(res, "couldn't get user", http.StatusBadGateway)
		return
	}

	user.Password = ""

	json.NewEncoder(res).Encode(user)

}

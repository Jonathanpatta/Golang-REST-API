package api

import (
	"encoding/json"
	"fmt"
	"strings"

	//"fmt"
	"net/http"

	"task1/structures"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

func HandleUsers(res http.ResponseWriter, req *http.Request) {
	//var id = primitive.NewObjectID()
	//var user = structures.User{Id: id, Name: "Jonathan", Email: "jonathan@gmail.com", Password: "lsajfsdafjl"}
	var url = req.URL.String()
	var splitUrl = strings.Split(url, "/")

	fmt.Println(splitUrl)

	fmt.Println(len(splitUrl))

	if len(splitUrl) == 2 {
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

	fmt.Fprint(res, "Invalid Url")

}

func CreateUser(res http.ResponseWriter, req *http.Request) {

	var userId = "ljkasdjf"

	id, error := primitive.ObjectIDFromHex(userId)
	if error != nil {
		fmt.Fprint(res, "invalid user Id")
		return
	}

	fmt.Println(req.Body)

	var user = structures.User{Id: id}
	fmt.Print("create user")

	json.NewEncoder(res).Encode(user)

}

func GetUser(res http.ResponseWriter, req *http.Request, Id string) {

	var user structures.User

	fmt.Print("get user")

	json.NewEncoder(res).Encode(user)

}

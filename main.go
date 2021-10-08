package main

import (
	//"fmt"
	"log"
	"net/http"
	"task1/api"
	db "task1/database"
)

func handleRequests() {
	http.HandleFunc("/users/", api.HandleUsers)
	http.HandleFunc("/posts/", api.HandlePosts)

	log.Fatal(http.ListenAndServe(":9000", nil))

}

func main() {
	db.CreateClient()

	handleRequests()
}

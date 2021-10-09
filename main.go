package main

import (
	"log"
	"net/http"
	"task1/api"
	"task1/auth"
	db "task1/database"
)

func handleRequests() {
	http.HandleFunc("/authenticate", auth.HandleAuth)
	http.HandleFunc("/users/", api.HandleUsers)
	http.HandleFunc("/posts/", api.HandlePosts)

	log.Fatal(http.ListenAndServe(":9000", nil))

}

func main() {
	db.CreateClient()

	handleRequests()
}

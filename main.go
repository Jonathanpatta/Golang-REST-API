package main

import (
	"fmt"
	"log"
	"net/http"
	db "task1/database"
)

var client = db.GetClient()

func homePage(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "hello there")
}

func handleRequests() {
	http.HandleFunc("/", homePage)

	log.Fatal(http.ListenAndServe(":9000", nil))

}

func main() {
	handleRequests()
}

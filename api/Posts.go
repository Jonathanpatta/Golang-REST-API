package api

import (
	"fmt"
	"net/http"
	"task1/database"
)

func HandlePosts(res http.ResponseWriter, req *http.Request) {
	var client = database.GetClient()
	fmt.Print(client)
	fmt.Fprint(res, "hello there from posts")
}

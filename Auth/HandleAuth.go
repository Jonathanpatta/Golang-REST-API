package auth

import (
	"fmt"
	"net/http"
)

func HandleAuth(res http.ResponseWriter, req *http.Request) {
	req.ParseForm()
	var username = req.PostForm.Get("Username")
	var password = req.PostForm.Get("Password")

	var token = CreateToken(username, password)
	if token == "" {
		fmt.Fprint(res, "invalid auth")
		return
	}

	fmt.Fprint(res, token)
}

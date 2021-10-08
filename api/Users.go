package api

import (
	"fmt"
	"net/http"
)

func AddUser(req http.ResponseWriter, res *http.Response) {
	fmt.Fprint(req, "hello there")
}

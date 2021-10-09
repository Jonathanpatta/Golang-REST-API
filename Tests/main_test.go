package main

import (
	"bytes"
	"encoding/json"
	"net/http"
	"testing"
)

func TestGetUserEndpoint(t *testing.T) {
	var url = "http://localhost:9000/users/61616dacc2c77d22124abc2c"
	res, err := http.Get(url)

	if res.StatusCode != 200 || err != nil {
		t.Error()
	}
}

func TestGetPostEndpoint(t *testing.T) {
	var url = "http://localhost:9000/posts/6161383af3bc8281d4c74ad2"
	res, err := http.Get(url)

	if res.StatusCode != 200 || err != nil {
		t.Error()
	}
}

func TestGetUserPostsEndpoint(t *testing.T) {

	var url = "http://localhost:9000/posts/user/61616dacc2c77d22124abc2b?page=1"
	res, err := http.Get(url)

	if res.StatusCode != 200 || err != nil {
		t.Error()
	}
}

func TestCreatePostEndpoint(t *testing.T) {

	var url = "http://localhost:9000/posts/"

	values := map[string]string{"Caption": "how are you so good!! aasdfsdf", "ImageUrl": "http://skadjfs.com/image", "UserId": "61616dacc2c77d22124abc2b"}

	jsonValue, _ := json.Marshal(values)

	res, err := http.Post(url, "application/json", bytes.NewBuffer(jsonValue))

	if res.StatusCode != 200 || err != nil {
		t.Error(res.StatusCode, res.Body)
	}
}

func TestCreateUserEndpoint(t *testing.T) {

	var url = "http://localhost:9000/users/"

	values := map[string]string{"Name": "Raj", "Email": "raj@mail.com", "Password": "pass@123"}

	jsonValue, _ := json.Marshal(values)

	res, err := http.Post(url, "application/json", bytes.NewBuffer(jsonValue))

	if res.StatusCode != 200 || err != nil {
		t.Error(res.StatusCode, res.Body)
	}
}

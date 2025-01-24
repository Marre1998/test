package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

var task string

type requestBody struct {
	Message string `json:"message"`
}

func PostHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" {
		var req requestBody
		err := json.NewDecoder(r.Body).Decode(&req)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
		task = req.Message
	}
}
func GetHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == "GET" {
		fmt.Fprintln(w, "Hello,", task)
	}
}

func main() {
	http.HandleFunc("/", GetHandler)
	http.HandleFunc("/task", PostHandler)

	http.ListenAndServe(":8080", nil)
}

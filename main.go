package main

import (
	"fmt"
	"encoding/json"
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/hello", GetHello) // set router

	err := http.ListenAndServe(":8080", nil) // set listen port
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
	log.Print("Listening on 8080")
}

func GetHello(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	name := r.URL.Query().Get("name")
	if name == "" {
		name = "Jane Doe"
	}

	ret := &Hello{
		Name: name,
		Message: fmt.Sprintf("Hello %s!", name),
	}

	json.NewEncoder(w).Encode(ret)
}

type Hello struct {
	Name string
	Message string
}

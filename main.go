package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

func main() {
	mux := http.NewServeMux()

	mux.HandleFunc("GET /home", helloHandler)
	mux.HandleFunc("GET /user", GetUserHandler)
	http.ListenAndServe(":8080", mux)

}

func helloHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "hello")

}

func GetUserHandler(w http.ResponseWriter, r *http.Request) {
	if err := json.NewEncoder(w).Encode(NewUser("Petya", 18)); err != nil {
		http.Error(w, "Json не поддерживается", http.StatusBadRequest)
	}

}

type User struct {
	Name string `json:"name"`
	Age  int    `json:"age"`
	
}

func NewUser(name string, age int) *User {
	return &User{
		Name: name,
		Age:  age,
	}
}



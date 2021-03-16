package main

import (
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

type User struct {
	UserId    int    `json:userid`
	Name      string `json:name`
	Lastname  string `json:lastname`
	Birthdate string `json:birthdate`
}

var users []User

func getUsers() {
	fmt.Println("getUsers")
}

func acceptUsers() {
	fmt.Println("acceptUsers json")
}

func editUser() {
	fmt.Println("editUser")
}

func deleteUser() {
	fmt.Println("deleteUser")
}

func main() {
	router := mux.NewRouter()
	router.HandleFunc("/users", getUsers).Methods("GET")
	router.HandleFunc("/users", acceptUsers).Methods("POST")
	router.HandleFunc("/users/{id}", editUser).Methods("PUT")
	router.HandleFunc("/users/{id}", deleteUser).Methods("DELETE")
	http.Handle("/", router)
}

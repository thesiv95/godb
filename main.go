package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

type User struct {
	UserId    int    `json:userid`
	Name      string `json:name`
	Lastname  string `json:lastname`
	Birthdate string `json:birthdate`
}

//var users []User

func getUsers(w http.ResponseWriter, r *http.Request) {
	fmt.Println("getUsers")
}

func acceptUsers(w http.ResponseWriter, r *http.Request) {
	fmt.Println("acceptUsers json")
}

func editUser(w http.ResponseWriter, r *http.Request) {
	fmt.Println("editUser")
}

func deleteUser(w http.ResponseWriter, r *http.Request) {
	fmt.Println("deleteUser")
}

func main() {
	router := mux.NewRouter()
	router.HandleFunc("/users", getUsers).Methods("GET")
	router.HandleFunc("/users", acceptUsers).Methods("POST")
	router.HandleFunc("/users/{id}", editUser).Methods("PUT")
	router.HandleFunc("/users/{id}", deleteUser).Methods("DELETE")
	log.Fatal(http.ListenAndServe(":8000", router))
}

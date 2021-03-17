package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

type User struct {
	UserId    int    `json:userid`
	Name      string `json:name`
	Lastname  string `json:lastname`
	Birthdate string `json:birthdate`
}

var users []User

func getUsers(w http.ResponseWriter, r *http.Request) {
	json.NewEncoder(w).Encode(users)
}

func acceptUsers(w http.ResponseWriter, r *http.Request) {
	fmt.Println("acceptUsers json")
}

func editUser(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)

	i, _ := strconv.Atoi(params["id"]) // var i

	for _, user := range users {
		if user.UserId == i {
			fmt.Printf("edit %d", i)
		}
	}
}

func deleteUser(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)

	i, _ := strconv.Atoi(params["id"])

	for _, user := range users {
		if user.UserId == i {
			fmt.Printf("delete %d", i)
		}
	}
}

func main() {
	router := mux.NewRouter()
	users = append(users,
		User{UserId: 1, Name: "John", Lastname: "Smith", Birthdate: "11/03/1986"},
		User{UserId: 2, Name: "Moshe", Lastname: "Dayan", Birthdate: "10/02/1993"},
		User{UserId: 3, Name: "Sarah", Lastname: "Connor", Birthdate: "05/11/1997"},
		User{UserId: 4, Name: "Nor", Lastname: "Levinov", Birthdate: "11/12/2000"},
		User{UserId: 5, Name: "Tal", Lastname: "Manov", Birthdate: "03/09/1991"})
	router.HandleFunc("/users", getUsers).Methods("GET")
	router.HandleFunc("/users", acceptUsers).Methods("POST")
	router.HandleFunc("/users/{id}", editUser).Methods("PUT")
	router.HandleFunc("/users/{id}", deleteUser).Methods("DELETE")
	log.Fatal(http.ListenAndServe(":8000", router))
}

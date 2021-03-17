package main

import (
	"log"
	"net/http"
	"os"

	"github.com/gorilla/mux"
	"github.com/lib/pq"
	"github.com/subosito/gotenv"
)

type User struct {
	UserId    int    `json:userid`
	Name      string `json:name`
	Lastname  string `json:lastname`
	Birthdate string `json:birthdate`
}

var users []User

func getUsers(w http.ResponseWriter, r *http.Request) {
}

func addUser(w http.ResponseWriter, r *http.Request) {
}

func editUser(w http.ResponseWriter, r *http.Request) {
}

func deleteUser(w http.ResponseWriter, r *http.Request) {
}

func init() {
	gotenv.Load()
}

func logFatal(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

func main() {

	pgUrl, err := pq.ParseURL(os.Getenv("ELEPHANTSQL_URL"))
	logFatal(err)

	log.Println(pgUrl)

	router := mux.NewRouter()

	router.HandleFunc("/users", getUsers).Methods("GET")
	router.HandleFunc("/users", addUser).Methods("POST")
	router.HandleFunc("/users/{id}", editUser).Methods("PUT")
	router.HandleFunc("/users/{id}", deleteUser).Methods("DELETE")
	log.Fatal(http.ListenAndServe(":8000", router))
}

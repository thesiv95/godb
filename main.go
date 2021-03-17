package main

import (
	"database/sql"
	"encoding/json"
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
var db *sql.DB

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

	db, err = sql.Open("postgres", pgUrl)
	logFatal(err)

	err = db.Ping()
	logFatal(err) // output will be shown only when error occurs

	router := mux.NewRouter()

	router.HandleFunc("/users", getUsers).Methods("GET")
	router.HandleFunc("/users", addUser).Methods("POST")
	router.HandleFunc("/users/{id}", editUser).Methods("PUT")
	router.HandleFunc("/users/{id}", deleteUser).Methods("DELETE")
	log.Fatal(http.ListenAndServe(":8000", router))
}

func getUsers(w http.ResponseWriter, r *http.Request) {
	var user User
	users = []User{}

	rows, err := db.Query("select * from users")
	logFatal(err)

	defer rows.Close()

	for rows.Next() {
		err := rows.Scan(&user.UserId, &user.Name, &user.Lastname, &user.Birthdate)
		logFatal(err) // output will be shown only when error occurs

		users = append(users, user)
	}
	json.NewEncoder(w).Encode(users)
}

func addUser(w http.ResponseWriter, r *http.Request) {
	var user User
	var userId int

	json.NewDecoder(r.Body).Decode(&user)
	err := db.QueryRow("insert into users (name, lastname, birthdate) values($1, $2, $3) RETURNING userid;", user.Name, user.Lastname, user.Birthdate).Scan(&userId)
	logFatal(err)

	json.NewEncoder(w).Encode(userId)
}

func editUser(w http.ResponseWriter, r *http.Request) {
	var user User
	json.NewDecoder(r.Body).Decode(&user)
	params := mux.Vars(r)

	result, err := db.Exec("update users set name=$1, lastname=$2, birthdate=$3 where userid=$4 RETURNING userid;", user.Name, user.Lastname, user.Birthdate, params["id"])
	logFatal(err)

	rowsUpdated, err := result.RowsAffected() // how many rows were updated?
	logFatal(err)

	json.NewEncoder(w).Encode(rowsUpdated)
}

func deleteUser(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)

	result, err := db.Exec("delete from users where userid=$1", params["id"])
	logFatal(err)

	rowsDeleted, err := result.RowsAffected()
	logFatal(err)

	json.NewEncoder(w).Encode(rowsDeleted)
}

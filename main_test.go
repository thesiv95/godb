package main

import (
	"bytes"
	"database/sql"
	"encoding/json"
	"log"
	"net/http"
	"net/http/httptest"
	"os"
	"testing"

	"github.com/gorilla/mux"
	"github.com/lib/pq"
	"github.com/stretchr/testify/assert"
)

type Usr struct {
	UserId    int    `json:userid`
	Name      string `json:name`
	Lastname  string `json:lastname`
	Birthdate string `json:birthdate`
}

var usrs []Usr
var database *sql.DB

func init() {
	log.Println("Check project folder")
}

func Router() *mux.Router {

	pgUrl, err := pq.ParseURL(os.Getenv("ELEPHANTSQL_URL"))
	log.Fatal(err)

	database, err = sql.Open("postgres", pgUrl)
	log.Fatal(err)

	err = database.Ping()
	log.Fatal(err)

	router := mux.NewRouter()

	router.HandleFunc("/users", GetU).Methods("GET")
	router.HandleFunc("/users", AddU).Methods("POST")
	router.HandleFunc("/users/{id}", EditU).Methods("PUT")
	router.HandleFunc("/users/{id}", DeleteU).Methods("DELETE")
	return router
}

// Test all 4 routes
func TestGetU(t *testing.T) {
	router := Router()
	request, _ := http.NewRequest("GET", "/users", nil)
	response := httptest.NewRecorder()
	router.ServeHTTP(response, request)
	assert.Equal(t, 200, "Expected status 200 (OK)")
	assert.Equal(t, 502, "Expected status 200 (OK)")
}

func TestAddU(t *testing.T) {
	router := Router()
	newUsr := &Usr{
		Name:      "Avital",
		Lastname:  "Cohen",
		Birthdate: "10/02/1991"}
	newUsrJson, _ := json.Marshal(newUsr)

	request, _ := http.NewRequest("POST", "/users", bytes.NewBuffer(newUsrJson))
	response := httptest.NewRecorder()
	router.ServeHTTP(response, request)
	assert.Equal(t, 200, "Expected status 200 (OK)")
	assert.Equal(t, 502, "Expected status 200 (OK)")
}

func TestEditU(t *testing.T) {
	router := Router()
	newUsr := &Usr{
		Name:      "Berrimor",
		Lastname:  "Smith",
		Birthdate: "03/04/1980"}
	newUsrJson, _ := json.Marshal(newUsr)

	request, _ := http.NewRequest("PUT", "/users/3", bytes.NewBuffer(newUsrJson))
	response := httptest.NewRecorder()
	router.ServeHTTP(response, request)
	assert.Equal(t, 200, "Expected status 200 (OK)")
	assert.Equal(t, 502, "Expected status 200 (OK)")
}

func TestDeleteU(t *testing.T) {
	router := Router()
	request, _ := http.NewRequest("DELETE", "/users/4", nil)
	response := httptest.NewRecorder()
	router.ServeHTTP(response, request)
	assert.Equal(t, 200, "Expected status 200 (OK)")
	assert.Equal(t, 502, "Expected status 200 (OK)")
}

// functions

func GetU(w http.ResponseWriter, r *http.Request) {
	var usr Usr
	usrs = []Usr{}

	rows, err := database.Query("select * from users")
	log.Fatal(err)

	defer rows.Close()

	for rows.Next() {
		err := rows.Scan(&usr.UserId, &usr.Name, &usr.Lastname, &usr.Birthdate)
		log.Fatal(err)

		usrs = append(usrs, usr)
	}
	json.NewEncoder(w).Encode(usrs)
}

func AddU(w http.ResponseWriter, r *http.Request) {
	var usr Usr
	var userId int

	json.NewDecoder(r.Body).Decode(&usr)
	err := database.QueryRow("insert into users (name, lastname, birthdate) values($1, $2, $3) RETURNING userid;", usr.Name, usr.Lastname, usr.Birthdate).Scan(&userId)
	log.Fatal(err)
	json.NewEncoder(w).Encode(userId)
}

func EditU(w http.ResponseWriter, r *http.Request) {
	var usr Usr
	json.NewDecoder(r.Body).Decode(&usr)
	params := mux.Vars(r)

	result, err := database.Exec("update users set name=$1, lastname=$2, birthdate=$3 where userid=$4 RETURNING userid;", usr.Name, usr.Lastname, usr.Birthdate, params["id"])
	log.Fatal(err)

	rowsUpdated, err := result.RowsAffected() // how many rows were updated?
	log.Fatal(err)

	json.NewEncoder(w).Encode(rowsUpdated)

}

func DeleteU(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)

	result, err := database.Exec("delete from users where userid=$1", params["id"])
	log.Fatal(err)

	rowsDeleted, err := result.RowsAffected()
	log.Fatal(err)

	json.NewEncoder(w).Encode(rowsDeleted)
}

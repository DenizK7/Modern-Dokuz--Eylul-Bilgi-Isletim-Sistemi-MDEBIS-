package main

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"

	_ "github.com/go-sql-driver/mysql"
)

// DB CONSTANTS COME HERE
var DB *sql.DB
var ACTIVE_USERS = make(map[string]user)

func main() {
	//Connect to the DB
	fmt.Println(string(hashPassword("354152")))
	var err error
	DB, err = sql.Open("mysql", "root:354152@tcp(127.0.0.1:3306)/mdebis")
	if DB != nil {

	}
	//trying functions
	//try any back-end function here

	//trying functions
	if err != nil {
		panic(err.Error())
	}

	//start to listen to port and response to the requests
	r := Router()
	fmt.Println("Starting server on the port 8080...")
	log.Fatal(http.ListenAndServe(":8080", r))
}

func enableCors(w *http.ResponseWriter) {
	(*w).Header().Set("Access-Control-Allow-Origin", "*")
	(*w).Header().Set("Access-Control-Allow-Headers", "Content-Type")

}

func getUser(sessionHash string) *user {
	user, found := ACTIVE_USERS[sessionHash]
	if found == false {
		return nil
	}
	return &user

}

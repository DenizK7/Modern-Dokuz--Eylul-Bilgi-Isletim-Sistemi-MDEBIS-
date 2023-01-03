package main

import (
	"database/sql"
	"fmt"
	"log"
	"math/rand"
	"net/http"
	"strconv"
	"strings"
	"time"

	_ "github.com/go-sql-driver/mysql"
)

// DB CONSTANTS COME HERE
var DB *sql.DB
var ACTIVE_USERS = make(map[string]*user)
var GRADES = [9]string{"AA", "BA", "BB", "CB", "CC", "DC", "DD", "FD", "FF"}

/*
PLEASE CHANGE BELOW SETTINGS ACCORDING TO YOUR LOCAL CONFIGURATIONS

# PLEASE USE DUMP FILE TO IMPORT THE DB TO YOUR LOCAL

# OTHERWISE, YOU WILL FACE AN ERROR

# IN CASE AN ERROR ALTHOUGH YOU ARE SURE YOU FOLLOWED THE STEPS RIGHT

PLEASE CONTACT US.

HAVE A GOOD DAY :)
*/
func main() {
	var (
		password = "354152"
		err      error
	)
	//Connect to the DB
	DB, err = sql.Open("mysql", "root:"+password+"@tcp(127.0.0.1:3306)/mdebis")
	if DB == nil || err != nil {
		fmt.Println("having a problem when trying to connect to db")
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
	return user

}

/*
This function returns randomly created hash
to hold the logged user's records
to be able to serve them later faster without a need to log in everytime
*/
func generateRandomSession() string {
	s1 := rand.NewSource(time.Now().UnixNano())
	r1 := rand.New(s1)
	hash := string(hashPassword(strconv.Itoa(r1.Intn(100000))))
	hash = strings.Replace(hash, "/", "", -1)
	return hash

}

func isGradeLegal(grade string) bool {
	for _, oneGrade := range GRADES {
		if oneGrade == grade {
			return true
		}
	}
	return false
}

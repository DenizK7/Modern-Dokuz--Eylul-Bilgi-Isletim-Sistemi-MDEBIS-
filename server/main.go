package main

import (
	"database/sql"
	"fmt"
	"log"
	"math/rand"
	"net/http"
	"strconv"
	"time"

	_ "github.com/go-sql-driver/mysql"
)

// DB CONSTANTS COME HERE
var DB *sql.DB
var ACTIVE_USERS = make(map[string]*user)
var GRADES = [9]string{"AA", "BA", "BB", "CB", "CC", "DC", "DD", "FD", "FF"}

func main() {
	//Connect to the DB
	var err error
	DB, err = sql.Open("mysql", "root:354152@tcp(127.0.0.1:3306)/mdebis")
	if DB != nil {

	}
	//trying functions
	fmt.Println(getRealPasswordStudent(2015501167))
	sessionHash := generateRandomSession()
	fmt.Println(sessionHash)
	var user user
	ACTIVE_USERS[sessionHash] = &user
	user.Lecturer = getLecturer(2000576383)

	//try any back-end function here
	if err != nil {
		panic(err.Error())
	}
	addAnnouncement(2000506140, 288, "SECOND ANNOUNCEMENT OF THE COURSE", "PLEASE before the class, read the chapter shared with you in the resources page of the class.")
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
	return string(hashPassword(strconv.Itoa(r1.Intn(100000))))

}

func isGradeLegal(grade string) bool {
	for _, oneGrade := range GRADES {
		if oneGrade == grade {
			return true
		}
	}
	return false
}

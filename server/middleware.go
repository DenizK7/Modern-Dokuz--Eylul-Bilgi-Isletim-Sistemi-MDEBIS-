package main

import (
	json "encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"golang.org/x/crypto/bcrypt"
	"math/rand"
	"net/http"
	"time"
)

/*
This function encodes all the GENERAL ANNOUNCEMENTS as a response
*/
func responseGetGeneralAnnouncements(w http.ResponseWriter, r *http.Request) {
	enableCors(&w)
	announcements := getGeneralAnnouncements()
	json.NewEncoder(w).Encode(announcements)
}

func responseStudentLogIn(w http.ResponseWriter, r *http.Request) {
	enableCors(&w)
	params := mux.Vars(r)
	id := params["username"]
	typedPassword := params["password"]
	encoder := json.NewEncoder(w)
	err, realPassword := getRealPasswordStudent(id)
	if err == false {
		encoder.Encode(false)
		return
	}
	if bcrypt.CompareHashAndPassword([]byte(realPassword), []byte(typedPassword)) != nil {
		// If the two passwords don't match, return a 401 status
		fmt.Println("password error")
		encoder.Encode("false")
		return
	}
	//create a session for the new user, type of student
	sessionHash := generateRandomSession()
	newUser := new(user)
	newUser.Student = getStudent(id)
	ACTIVE_USERS[sessionHash] = *newUser
	encoder.Encode(sessionHash)
	return
}

/*
this function encodes the courses as a response
*/
func responseGetCourses(w http.ResponseWriter, r *http.Request) {
	enableCors(&w)
	encoder := json.NewEncoder(w)
	params := mux.Vars(r)
	sessionHash := params["sessionHash"]
	user := getUser(sessionHash)
	if user == nil || user.Student == nil {
		fmt.Println("! ! !first you MUST log in! ! !")
		encoder.Encode(false)
		return
	}
	courses := getCourses(user.Student)
	json.NewEncoder(w).Encode(courses)
}

/*
This function responses the request by encoding the timetable in json format
!ATTENTION! - STUDENT MUST ALREADY LOGGED IN - !ATTENTION!
*/

func responseGetTimeTable(w http.ResponseWriter, r *http.Request) {
	enableCors(&w)
	params := mux.Vars(r)
	sessionHash := params["sessionHash"]
	user := getUser(sessionHash)
	if user == nil || user.Student == nil {
		fmt.Println("! ! !first you MUST log in! ! !")
		json.NewEncoder(w).Encode(false)
	}
	timeTable := getCoursesTimeTable(user.Student)
	json.NewEncoder(w).Encode(timeTable)
}

/*
This function returns randomly created hash
to hold the logged user's records
to be able to serve them later faster without a need to log in everytime
*/
func generateRandomSession() string {
	s1 := rand.NewSource(time.Now().UnixNano())
	r1 := rand.New(s1)
	return string(hashPassword(string(r1.Intn(100000))))

}

/*
This function encodes the logging lecturer if there is a match in the DB with the given id-password pair
*/
func responseLecturerLogIn(w http.ResponseWriter, r *http.Request) {
	enableCors(&w)
	encoder := json.NewEncoder(w)
	params := mux.Vars(r)
	id := params["username"]
	typedPassword := params["password"]
	isFound, realPassword := getRealPasswordLecturer(id)
	if isFound == false {
		encoder.Encode(false)
		return
	}
	if bcrypt.CompareHashAndPassword([]byte(realPassword), []byte(typedPassword)) != nil {
		// If the two passwords do not match, return a 401 status
		fmt.Println("password error")
		encoder.Encode("false")
	}
	//create a session for the new user, type of lecturer
	sessionHash := generateRandomSession()
	newUser := new(user)
	newUser.Lecturer = getLecturer(id)
	ACTIVE_USERS[sessionHash] = *newUser
	encoder.Encode(sessionHash)
	return
}

/*
This function encodes the logging manager if there is a match in the DB with the given id-password pair
*/
func responseAdminLogIn(w http.ResponseWriter, r *http.Request) {
	enableCors(&w)
	encoder := json.NewEncoder(w)
	params := mux.Vars(r)
	id := params["username"]
	typedPassword := params["password"]
	isFound, realPassword := getRealPasswordManager(id)
	if isFound == false {
		fmt.Println("no such a student")
		encoder.Encode(false)
	}
	if bcrypt.CompareHashAndPassword([]byte(realPassword), []byte(typedPassword)) != nil {
		// If the two passwords don't match, return a 401 status
		fmt.Println("password error")
		encoder.Encode("WRONG PASSWORD!")
		return
	}
	//create a session for the new user, type of lecturer
	sessionHash := generateRandomSession()
	newUser := new(user)
	newUser.Manager = getManager(id)
	ACTIVE_USERS[sessionHash] = *newUser
	encoder.Encode(sessionHash)
}

/*
This function hashes the given string
*/
func hashPassword(password string) []byte {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), 8)
	if err != nil {
		fmt.Printf("error occurred when hashing")
		return nil
	}
	return hashedPassword
}

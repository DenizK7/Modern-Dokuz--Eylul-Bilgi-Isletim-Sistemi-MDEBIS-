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
	rows, err := DB.Query("SELECT * FROM mdebis.general_announcement")
	if err != nil {
		fmt.Println("error occurred when getting general announcements")
	}

	var announcements []general_announcement

	// Loop through rows, using Scan to assign column data to struct fields.
	for rows.Next() {
		var announcement general_announcement
		if err := rows.Scan(&announcement.AnnouncementId, &announcement.Title, &announcement.Content, &announcement.Link); err != nil {
			fmt.Println("error occurred when getting general announcements")
		}
		announcements = append(announcements, announcement)
	}
	err = json.NewEncoder(w).Encode(announcements)
	if err != nil {
		return
	}
}

/*
This function encodes the logging student if there is a match in the DB with the given id-password pair
*/

func responseStudentLogIn(w http.ResponseWriter, r *http.Request) {
	enableCors(&w)
	params := mux.Vars(r)
	id := params["username"]
	typedPassword := params["password"]
	encoder := json.NewEncoder(w)
	err, realPassword := getRealPasswordStudent(id)
	if err == true {
		encoder.Encode(false)
		return
	}
	if bcrypt.CompareHashAndPassword([]byte(realPassword), []byte(typedPassword)) != nil {
		// If the two passwords don't match, return a 401 status
		fmt.Println("password error")
		encoder.Encode("false")
		return
	}
	//create student struct and return its information
	isAchieved, student := getStudent(id)
	if isAchieved == false {
		fmt.Println("error occured when logging")
		encoder.Encode(false)
		return
	}
	getCourses(student)
	encoder.Encode(student)
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
	if user == nil {
		fmt.Println("! ! !first you MUST log in! ! !")
		encoder.Encode(false)
		return
	}
	getCourses(user.Student)
	json.NewEncoder(w).Encode(user.Student.Courses)
}

/*
T
his function responses the request by encoding the timetable in json format
!ATTENTION! - STUDENT MUST ALREADY LOGGED IN - !ATTENTION!
*/

func responseGetTimeTable(w http.ResponseWriter, r *http.Request) {
	enableCors(&w)
	params := mux.Vars(r)
	sessionHash := params["sessionHash"]
	user := getUser(sessionHash)
	if user == nil {
		fmt.Println("! ! !first you MUST log in! ! !")
	}
	getCoursesTimeTable(user.Student)
	json.NewEncoder(w).Encode(user.Student.TimeTable)
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
	//create student struct and return its information
	isFoundLecturer, lecturer := getLecturer(id)
	if isFoundLecturer == false {
		encoder.Encode(false)
		return
	}
	encoder.Encode(lecturer)
}

/*
This function encodes the logging manager if there is a match in the DB with the given id-password pair
*/
func responseManagerLogIn(w http.ResponseWriter, r *http.Request) {
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
	//create student struct and return its information
	isCreatedManager, manager := getManager(id)
	if isCreatedManager == false {
		encoder.Encode(false)
		fmt.Println("problem with finding the manager in the DB")
		return
	}
	encoder.Encode(manager)
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

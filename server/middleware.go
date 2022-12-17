package main

import (
	json "encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"golang.org/x/crypto/bcrypt"
	"net/http"
	"strconv"
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
	encoder := json.NewEncoder(w)
	params := mux.Vars(r)
	id, err := strconv.Atoi(params["username"])
	if err != nil {
		fmt.Println("error wen converting id to int ")
		encoder.Encode(false)
		return
	}
	typedPassword := params["password"]
	isFound, realPassword := getRealPasswordStudent(id)
	if isFound == false {
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
	ACTIVE_USERS[sessionHash] = newUser
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
	if user == nil {
		fmt.Println("! ! !first you MUST log in! ! !")
		encoder.Encode(false)
		return
	}
	if user.Student != nil {
		courses := getCoursesOfAStudent(user.Student.Id)
		json.NewEncoder(w).Encode(courses)
	}
	if user.Lecturer != nil {
		courses := getCoursesOfALecturer(user.Lecturer)
		json.NewEncoder(w).Encode(courses)
	}

}

func responseAddGrade(w http.ResponseWriter, r *http.Request) {
	enableCors(&w)
	encoder := json.NewEncoder(w)
	params := mux.Vars(r)
	sessionHash := params["sessionHash"]
	grade := params["grade"]
	courseid, _ := strconv.Atoi(params["courseId"])
	studentId, _ := strconv.Atoi(params["studentId"])

	user := getUser(sessionHash)
	if user == nil {
		fmt.Println("! ! !first you MUST log in! ! !")
		encoder.Encode(false)
		return
	}
	if user.Student != nil {
		json.NewEncoder(w).Encode(false)
		return
	}
	if user.Lecturer != nil {
		if isGradeLegal(grade) == false {
			json.NewEncoder(w).Encode(false)
			return
		}
		json.NewEncoder(w).Encode(addGrade(user.Lecturer.Id, courseid, studentId, grade))
	}
}

func responseAddAnnouncement(w http.ResponseWriter, r *http.Request) {
	enableCors(&w)
	encoder := json.NewEncoder(w)
	params := mux.Vars(r)
	sessionHash := params["sessionHash"]
	title := params["title"]
	content := params["content"]
	courseId, _ := strconv.Atoi(params["courseId"])

	user := getUser(sessionHash)
	if user == nil {
		fmt.Println("! ! !first you MUST log in! ! !")
		encoder.Encode(false)
		return
	}
	if user.Student != nil {
		json.NewEncoder(w).Encode(false)
		return
	}
	if user.Lecturer != nil {
		json.NewEncoder(w).Encode(addAnnouncement(user.Lecturer.Id, courseId, title, content))
		return
	}
}

/*
This function responses the request by encoding the timetable in json format
!ATTENTION! - STUDENT MUST ALREADY LOGGED IN - !ATTENTION!
*/

func responseChangeActiveOfCourse(w http.ResponseWriter, r *http.Request) {
	enableCors(&w)
	encoder := json.NewEncoder(w)
	params := mux.Vars(r)
	sessionHash := params["sessionHash"]
	courseId, err := strconv.Atoi(params["courseId"])
	if err != nil {
		encoder.Encode(false)
		return
	}
	assignedStatus := params["assignedStatus"]
	user := getUser(sessionHash)
	if user == nil || user.Student != nil {
		encoder.Encode(false)
		return
	}
	//!CHECK THIS COURSE IS OWNED BY THIS LECTURER!
	var isOwned = checkACourseOwned(user, courseId)
	if isOwned == false {
		fmt.Println("course does not belong this user")
		encoder.Encode("course does not belong this user")
		return
	}

	//Find what user wants the course to be
	var isActive bool
	switch assignedStatus {
	case "true":
		isActive = true
	case "false":
		isActive = false
	default:
		encoder.Encode(false)
		return
	}
	//Make the course what user wants the course to be
	changeStatusOfCourse(courseId, isActive)

	/*todo
	DB'e check statement eklenebilir.

	Front end tarafına değişikliğin yapılamayacağını (zaten çoktan active veya inactive) olduğu bilgisi de DÖNDÜRÜLMELİ.

	Ki böylede front end kullanıcıyı uyarabilsin!
	*/

}

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

func responseGetDepartmentOfStudent(w http.ResponseWriter, r *http.Request) {
	enableCors(&w)
	encoder := json.NewEncoder(w)
	params := mux.Vars(r)
	sessionHash := params["sessionHash"]
	user := getUser(sessionHash)
	if user == nil || user.Student == nil {
		encoder.Encode(false)
		return
	}
	id := user.Student.Id
	user.Student = getStudent(id)
	encoder.Encode(getDepartmentOfStudent(id))
}

/*
This function encodes the logging lecturer if there is a match in the DB with the given id-password pair
*/
func responseLecturerLogIn(w http.ResponseWriter, r *http.Request) {
	enableCors(&w)
	encoder := json.NewEncoder(w)
	params := mux.Vars(r)
	id, _ := strconv.Atoi(params["username"])
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
	ACTIVE_USERS[sessionHash] = newUser
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
	id, _ := strconv.Atoi(params["username"])
	typedPassword := params["password"]
	isFound, realPassword := getRealPasswordAdmin(id)
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
	newUser.Manager = getAdmin(id)
	ACTIVE_USERS[sessionHash] = newUser
	encoder.Encode(sessionHash)
}

/*
This function hashes the given string
*/
func hashPassword(password string) []byte {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), 8)
	if err != nil {
		fmt.Println(err.Error())
		return nil
	}
	return hashedPassword
}

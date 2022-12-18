package main

import (
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"golang.org/x/crypto/bcrypt"
	"net/http"
	"strconv"
)

/*
This function encodes all the GENERAL ANNOUNCEMENTS as a response
*/
func responseGetGeneralAnnouncements(w http.ResponseWriter, _ *http.Request) {
	enableCors(&w)
	announcements := getGeneralAnnouncements()
	err := json.NewEncoder(w).Encode(announcements)
	if err != nil {
		fmt.Println(err.Error())
		return
	}
}

func responseStudentLogIn(w http.ResponseWriter, r *http.Request) {
	enableCors(&w)
	encoder := json.NewEncoder(w)
	params := mux.Vars(r)
	id, err := strconv.Atoi(params["username"])
	if err != nil {
		fmt.Println("error wen converting id to int ")
		err := encoder.Encode(false)
		if err != nil {
			fmt.Println(err.Error())
			return
		}
		return
	}
	typedPassword := params["password"]
	isFound, realPassword := getRealPasswordStudent(id)
	if isFound == false {
		err := encoder.Encode(false)
		if err != nil {
			fmt.Println(err.Error())

			return
		}
		return
	}
	if bcrypt.CompareHashAndPassword([]byte(realPassword), []byte(typedPassword)) != nil {
		// If the two passwords don't match, return a 401 status
		fmt.Println("password error")
		err := encoder.Encode("false")
		if err != nil {
			fmt.Println(err.Error())
			return
		}
		return
	}
	//create a session for the new user, type of student
	sessionHash := generateRandomSession()
	newUser := new(user)
	newUser.Student = getStudent(id)
	ACTIVE_USERS[sessionHash] = newUser
	err = encoder.Encode(sessionHash)
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	return
}

/*
this function encodes the courses as a response
*/

func responseDeleteStudent(w http.ResponseWriter, r *http.Request) {
	enableCors(&w)
	encoder := json.NewEncoder(w)
	params := mux.Vars(r)
	sessionHash := params["sessionHash"]
	studentId, _ := strconv.Atoi(params["studentId"])
	user := getUser(sessionHash)
	if isUserRight(user, 3) == false {
		err := encoder.Encode(false)
		if err != nil {
			fmt.Println(err.Error())
			return
		}
		return
	}
	err := encoder.Encode(deleteStudent(studentId))
	if err != nil {
		fmt.Println(err.Error())
		return
	}
}

func responseDeleteLecturer(w http.ResponseWriter, r *http.Request) {
	enableCors(&w)
	encoder := json.NewEncoder(w)
	params := mux.Vars(r)
	sessionHash := params["sessionHash"]
	lecturerId, _ := strconv.Atoi(params["lecturerId"])
	user := getUser(sessionHash)
	if isUserRight(user, 3) == false {
		err := encoder.Encode(false)
		if err != nil {
			fmt.Println(err.Error())
			return
		}
		return
	}
	err := encoder.Encode(deleteLecturer(lecturerId))
	if err != nil {
		fmt.Println(err.Error())
		return
	}
}

func responseGetStudentsOfCourse(w http.ResponseWriter, r *http.Request) {
	enableCors(&w)
	encoder := json.NewEncoder(w)
	params := mux.Vars(r)
	sessionHash := params["sessionHash"]
	courseId, _ := strconv.Atoi(params["courseId"])
	user := getUser(sessionHash)
	if isUserRight(user, 2) == false {
		err := encoder.Encode(false)
		if err != nil {
			fmt.Println(err.Error())
			return
		}
		return
	}
	err := encoder.Encode(getStudentsOfCourse(user.Lecturer.Id, courseId))
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	return
}
func responseGetAnnouncementOfCourse(w http.ResponseWriter, r *http.Request) {
	enableCors(&w)
	encoder := json.NewEncoder(w)
	params := mux.Vars(r)
	sessionHash := params["sessionHash"]
	courseId, _ := strconv.Atoi(params["courseId"])
	user := getUser(sessionHash)
	if !(isUserRight(user, 1) || isUserRight(user, 2)) {
		err := encoder.Encode(false)
		if err != nil {
			fmt.Println(err.Error())
			return
		}
		return
	}
	err := encoder.Encode(getAnnouncementOfCourse(courseId))
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	return
}

func responseAddGrade(w http.ResponseWriter, r *http.Request) {
	enableCors(&w)
	encoder := json.NewEncoder(w)
	params := mux.Vars(r)
	sessionHash := params["sessionHash"]
	grade := params["grade"]
	courseId, _ := strconv.Atoi(params["courseId"])
	studentId, _ := strconv.Atoi(params["studentId"])

	user := getUser(sessionHash)
	isUserRight := isUserRight(user, 2)
	if isUserRight == false {
		fmt.Println("! ! !first you MUST log in! ! !")
		err := encoder.Encode(false)
		if err != nil {
			fmt.Println(err.Error())
			return
		}
		return
	}
	if isGradeLegal(grade) == false {
		err := json.NewEncoder(w).Encode(false)
		if err != nil {
			fmt.Println(err.Error())
			return
		}
		return
	}
	err := json.NewEncoder(w).Encode(addGrade(user.Lecturer.Id, courseId, studentId, grade))
	if err != nil {
		fmt.Println(err.Error())
		return
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
		err := encoder.Encode(false)
		if err != nil {
			fmt.Println(err.Error())
			return
		}
		return
	}
	if user.Student != nil {
		err := json.NewEncoder(w).Encode(false)
		if err != nil {
			fmt.Println(err.Error())
			return
		}
		return
	}
	if user.Lecturer != nil {
		err := json.NewEncoder(w).Encode(addAnnouncement(user.Lecturer.Id, courseId, title, content))
		if err != nil {
			fmt.Println(err.Error())
			return
		}
		return
	}
}

func responseGetCourses(w http.ResponseWriter, r *http.Request) {
	enableCors(&w)
	encoder := json.NewEncoder(w)
	params := mux.Vars(r)
	sessionHash := params["sessionHash"]
	user := getUser(sessionHash)
	if user == nil {
		fmt.Println("! ! !first you MUST log in! ! !")
		err := encoder.Encode(false)
		if err != nil {
			fmt.Println(err.Error())
			return
		}
		return
	}
	if user.Student != nil {
		courses := getCoursesOfAStudent(user.Student.Id)
		err := json.NewEncoder(w).Encode(courses)
		if err != nil {
			fmt.Println(err.Error())
			return
		}
	}
	if user.Lecturer != nil {
		courses := getCoursesOfALecturer(user.Lecturer)
		err := json.NewEncoder(w).Encode(courses)
		if err != nil {
			fmt.Println(err.Error())
			return
		}
	}

}

func responseGetPastCoursesOfStudent(w http.ResponseWriter, r *http.Request) {
	enableCors(&w)
	encoder := json.NewEncoder(w)
	params := mux.Vars(r)
	sessionHash := params["sessionToken"]
	user := getUser(sessionHash)
	if user == nil || user.Student == nil {
		err := encoder.Encode(false)
		if err != nil {
			fmt.Println(err.Error())
			return
		}
		return
	}
	id := user.Student.Id
	err := encoder.Encode(getPastCoursesOfStudent(id))
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	return
}

func isUserRight(user *user, whichUser int) bool {
	//whichUser
	//1 --> student
	//2 --> lecturer
	//3 --> manager

	if user == nil {
		return false
	}
	if user.Student == nil && whichUser == 1 {
		return false
	}

	if user.Lecturer == nil && whichUser == 2 {
		return false
	}

	if user.Manager == nil && whichUser == 3 {
		return false
	}
	return true
}

func responseChangeNonAttendance(w http.ResponseWriter, r *http.Request) {
	enableCors(&w)
	encoder := json.NewEncoder(w)
	params := mux.Vars(r)
	sessionHash := params["sessionHash"]
	courseId, _ := strconv.Atoi(params["courseId"])
	studentId, _ := strconv.Atoi(params["studentId"])
	nonAttendance, _ := strconv.Atoi(params["nonAttendance"])
	user := getUser(sessionHash)
	if !isUserRight(user, 2) {
		err := encoder.Encode(false)
		if err != nil {
			fmt.Println(err.Error())
			return
		}
		return
	}
	err := encoder.Encode(changeNonAttendance(user.Student.Id, courseId, studentId, nonAttendance))
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	return
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
		err := encoder.Encode(false)
		if err != nil {
			fmt.Println(err.Error())
			return
		}
		return
	}
	assignedStatus := params["assignedStatus"]
	user := getUser(sessionHash)
	if !isUserRight(user, 2) {
		err := encoder.Encode(false)
		if err != nil {
			fmt.Println(err.Error())
			return
		}
		return
	}
	//!CHECK THIS COURSE IS OWNED BY THIS LECTURER!
	var isOwned = isLecturerOwnTheCourse(courseId, user.Lecturer.Id)
	if isOwned == false {
		fmt.Println("course does not belong this user")
		err := encoder.Encode("course does not belong this user")
		if err != nil {
			fmt.Println(err.Error())
			return
		}
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
		err := encoder.Encode(false)
		if err != nil {
			fmt.Println(err.Error())
			return
		}
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
	if !isUserRight(user, 1) {
		fmt.Println("! ! !first you MUST log in! ! !")
		err := json.NewEncoder(w).Encode(false)
		if err != nil {
			fmt.Println(err.Error())
			return
		}
	}
	timeTable := getCoursesTimeTable(user.Student)
	err := json.NewEncoder(w).Encode(timeTable)
	if err != nil {
		fmt.Println(err.Error())

		return
	}
}

func responseGetDepartmentOfStudent(w http.ResponseWriter, r *http.Request) {
	enableCors(&w)
	encoder := json.NewEncoder(w)
	params := mux.Vars(r)
	sessionHash := params["sessionToken"]
	user := getUser(sessionHash)
	if !isUserRight(user, 1) {
		err := encoder.Encode(false)
		if err != nil {
			fmt.Println(err.Error())
			return
		}
		return
	}
	id := user.Student.Id
	user.Student = getStudent(id)
	err := encoder.Encode(getDepartmentOfStudent(id))
	if err != nil {
		fmt.Println(err.Error())
		return
	}
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
		err := encoder.Encode(false)
		if err != nil {
			fmt.Println(err.Error())
			return
		}
		return
	}
	if bcrypt.CompareHashAndPassword([]byte(realPassword), []byte(typedPassword)) != nil {
		// If the two passwords do not match, return a 401 status
		fmt.Println("password error")
		err := encoder.Encode("false")
		if err != nil {
			fmt.Println(err.Error())
			return
		}
	}
	//create a session for the new user, type of lecturer
	sessionHash := generateRandomSession()
	newUser := new(user)
	newUser.Lecturer = getLecturer(id)
	ACTIVE_USERS[sessionHash] = newUser
	err := encoder.Encode(sessionHash)
	if err != nil {
		fmt.Println(err.Error())
		return
	}
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
		err := encoder.Encode(false)
		if err != nil {
			fmt.Println(err.Error())
			return
		}
	}
	if bcrypt.CompareHashAndPassword([]byte(realPassword), []byte(typedPassword)) != nil {
		// If the two passwords don't match, return a 401 status
		fmt.Println("password error")
		err := encoder.Encode("WRONG PASSWORD!")
		if err != nil {
			fmt.Println(err.Error())
			return
		}
		return
	}
	//create a session for the new user, type of lecturer
	sessionHash := generateRandomSession()
	newUser := new(user)
	newUser.Manager = getAdmin(id)
	ACTIVE_USERS[sessionHash] = newUser
	err := encoder.Encode(sessionHash)
	if err != nil {
		fmt.Println(err.Error())
		return
	}
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

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
	} else {
		addLog("Admin", user.Manager.Id, "Delete", "student", "DELETE STUDENT WITH ID "+strconv.Itoa(studentId))
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
	} else {
		addLog("Admin", user.Manager.Id, "Delete", "lecturer", "DELETE LECTURER WITH ID "+strconv.Itoa(lecturerId))
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

func responseAddGrade(w http.ResponseWriter, r *http.Request) {
	enableCors(&w)
	encoder := json.NewEncoder(w)
	params := mux.Vars(r)
	sessionHash := params["sessionHash"]
	grade := params["grade"]
	courseId, _ := strconv.Atoi(params["courseId"])
	studentId, _ := strconv.Atoi(params["studentId"])
	fmt.Println(courseId)
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
	} else {
		addLog("Lecturer", user.Lecturer.Id, "Update", "course_has_student", "SET GRADE TO "+grade+" OF STUDENT WITH ID "+strconv.Itoa(studentId)+" FOR COURSE ID WITH "+strconv.Itoa(courseId))
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
		} else {
			addLog("Lecturer", user.Lecturer.Id, "Insert", "course_has_announcement", "INSERT NEW ANNOUNCEMENT FOR THE COURSE ID "+strconv.Itoa(courseId)+" WITH THE TITLE "+title)
		}
		return
	}
}

func convertHomePageEntryLecturer(courses []course) []homePageEntryLecturer {
	var homePageEntries []homePageEntryLecturer
	for _, course := range courses {
		var homePageEntry homePageEntryLecturer
		homePageEntry.CourseName = course.Name
		homePageEntry.Announcements = course.Announcements
		homePageEntry.Credit = course.Credit
		homePageEntry.AttendanceLimit = course.AttandenceLimit
		homePageEntry.TimeInfo = course.Time_Inf
		homePageEntry.DepName = getDepartmentName(course.Dep_Id)
		homePageEntry.LecName = getLecturerNamesOfCourse(course.Id)
		homePageEntry.CourseId = course.Id
		homePageEntries = append(homePageEntries, homePageEntry)
	}
	return homePageEntries
}
func convertHomePageEntryStudent(courses []course, studentId int) []homePageEntryStudent {
	var homePageEntries []homePageEntryStudent
	for _, course := range courses {
		var homePageEntry homePageEntryStudent
		homePageEntry.CourseName = course.Name
		homePageEntry.Announcements = course.Announcements
		homePageEntry.Credit = course.Credit
		homePageEntry.AttendanceLimit = course.AttandenceLimit
		homePageEntry.TimeInfo = course.Time_Inf
		homePageEntry.DepName = getDepartmentName(course.Dep_Id)
		homePageEntry.LecName = getLecturerNamesOfCourse(course.Id)
		homePageEntry.CurrentNonAttendance = getNonAttendanceOfStudent(studentId, course.Id)
		homePageEntry.CourseId = course.Id
		homePageEntries = append(homePageEntries, homePageEntry)
	}
	return homePageEntries
}

func responseAddCourse(w http.ResponseWriter, r *http.Request) {
	//lecturer *lecturer, courseName string, attendanceLimit int, credit int
	enableCors(&w)
	encoder := json.NewEncoder(w)
	params := mux.Vars(r)
	sessionHash := params["sessionHash"]
	courseName := params["courseName"]
	attendanceLimit, _ := strconv.Atoi(params["attendanceLimit"])
	credit, _ := strconv.Atoi(params["credit"])
	user := getUser(sessionHash)
	if !isUserRight(user, 2) {
		err := encoder.Encode(false)
		if err != nil {
			fmt.Println(err.Error())
			return
		}
		return
	}
	lecturer := user.Lecturer
	err := encoder.Encode(addCourse(lecturer, courseName, attendanceLimit, credit))
	if err != nil {
		return
	} else {
		addLog("Lecturer", user.Lecturer.Id, "Insert", "course", "ADD NEW COURSE WITH THE NAME "+courseName)
	}

	return
}

func responseGetHomeEntry(w http.ResponseWriter, r *http.Request) {
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
		err := json.NewEncoder(w).Encode(convertHomePageEntryStudent(courses, user.Student.Id))
		if err != nil {
			fmt.Println(err.Error())
			return
		}
	}
	if user.Lecturer != nil {
		courses := getCoursesOfALecturer(user.Lecturer)
		err := json.NewEncoder(w).Encode(convertHomePageEntryLecturer(courses))
		if err != nil {
			fmt.Println(err.Error())
			return
		}
	}

}
func responseLogOut(w http.ResponseWriter, r *http.Request) {
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
	delete(ACTIVE_USERS, sessionHash)
	err := encoder.Encode(true)
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	return
}
func responseGetStudents(w http.ResponseWriter, r *http.Request) {
	enableCors(&w)
	encoder := json.NewEncoder(w)
	params := mux.Vars(r)
	sessionHash := params["sessionHash"]
	user := getUser(sessionHash)
	if isUserRight(user, 3) == false {
		err := encoder.Encode(false)
		if err != nil {
			fmt.Println(err.Error())
			return
		}
		return
	}
	encoder.Encode(getAllStudents())
}
func responseGetLecturers(w http.ResponseWriter, r *http.Request) {
	enableCors(&w)
	encoder := json.NewEncoder(w)
	params := mux.Vars(r)
	sessionHash := params["sessionHash"]
	user := getUser(sessionHash)
	if isUserRight(user, 3) == false {
		err := encoder.Encode(false)
		if err != nil {
			fmt.Println(err.Error())
			return
		}
		return
	}
	encoder.Encode(getAllLecturers())
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

func responseCreateLecturer(w http.ResponseWriter, r *http.Request) {
	enableCors(&w)
	encoder := json.NewEncoder(w)
	params := mux.Vars(r)
	sessionHash := params["sessionHash"]
	user := getUser(sessionHash)
	if isUserRight(user, 3) == false {
		err := encoder.Encode(false)
		if err != nil {
			fmt.Println(err.Error())
			return
		}
		return
	}
	//id int, password string, title string, name string, surname string, departmentName string
	id, err := strconv.Atoi(params["id"])
	if err != nil {
		fmt.Println("error occured when casting string to int")
		encoder.Encode(false)
		return
	}
	password := params["password"]
	password = string(hashPassword(password))
	title := params["title"]
	name := params["name"]
	surname := params["surname"]
	departmentName := params["departmentName"]
	err = encoder.Encode(createLecturer(id, password, title, name, surname, departmentName))
	if err != nil {
		return
	} else {
		addLog("Admin", user.Manager.Id, "Insert", "lecturer", "CREATE LECTURER WITH NAME "+title+" "+name+" "+surname+" IN THE DEPARTMENT WITH NAME "+departmentName)
	}
}
func responseCreateStudent(w http.ResponseWriter, r *http.Request) {
	enableCors(&w)
	encoder := json.NewEncoder(w)
	params := mux.Vars(r)
	sessionHash := params["sessionHash"]
	user := getUser(sessionHash)
	if isUserRight(user, 3) == false {
		err := encoder.Encode(false)
		if err != nil {
			fmt.Println(err.Error())
			return
		}
		return
	}
	//id int, password string, title string, name string, surname string, departmentName string
	id, err := strconv.Atoi(params["id"])
	if err != nil {
		fmt.Println("error occured when casting string to int")
		encoder.Encode(false)
		return
	}
	password := params["password"]
	password = string(hashPassword(password))
	name := params["name"]
	surname := params["surname"]
	departmentName := params["departmentName"]
	err = encoder.Encode(createStudent(id, password, name, surname, departmentName))
	if err != nil {
		return
	} else {
		addLog("Admin", user.Manager.Id, "Insert", "Student", "INSERT NEW STUDENT WITH ID "+strconv.Itoa(id)+" AND THE NAME "+name+" "+surname)
	}
}
func responseGetAllDepartmentNames(w http.ResponseWriter, r *http.Request) {
	enableCors(&w)
	encoder := json.NewEncoder(w)
	params := mux.Vars(r)
	sessionHash := params["sessionHash"]
	user := getUser(sessionHash)
	if isUserRight(user, 3) == false {
		err := encoder.Encode(false)
		if err != nil {
			fmt.Println(err.Error())
			return
		}
		return
	}
	encoder.Encode(getAllDepartmentNames())
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
	err := encoder.Encode(changeNonAttendance(user.Lecturer.Id, courseId, studentId, nonAttendance))
	if err != nil {
		fmt.Println(err.Error())
		return
	} else {
		addLog("Lecturer", user.Lecturer.Id, "Update", "course_has_student", "SET NON_ATTENDANCE TO "+strconv.Itoa(nonAttendance)+" OF STUDENT WITH ID "+strconv.Itoa(studentId)+" FOR COURSE ID WITH "+strconv.Itoa(courseId))
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
	//Make the course what user wants the course to be
	encoder.Encode(changeStatusOfCourse(courseId))

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

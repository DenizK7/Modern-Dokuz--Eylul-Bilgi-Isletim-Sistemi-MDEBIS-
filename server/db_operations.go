package main

import (
	"database/sql"
	"fmt"
	"strings"
)

/*
Returns the password taken from DB if there is match for the given id of a student in the DB
*/
func getRealPasswordStudent(id int) (bool, string) {
	var realPassword string
	query := "CALL student_get_password(?)"
	if err := DB.QueryRow(query, id).Scan(&realPassword); err != nil {
		fmt.Println(err.Error())
		return false, ""
	}
	return true, realPassword
}

/*
Returns the password taken from DB if there is match for the given id of a lecturer in the DB
*/
func getRealPasswordLecturer(id int) (bool, string) {
	var realPassword string
	query := "CALL lecturer_get_password(?)"

	if err := DB.QueryRow(query, id).Scan(&realPassword); err != nil {
		fmt.Println(err.Error())
	}
	return true, realPassword
}

func getDepartmentName(depId int) string {
	query := "select Name from department WHERE Department_Id=?"
	var depName string
	if err := DB.QueryRow(query, depId).Scan(&depName); err != nil {
		fmt.Println(err.Error())
	}
	return depName
}
func getLecturerNamesOfCourse(courseId int) string {
	query := "select Title,Name,Surname from lecturer where Lecturer_Id IN " +
		"(select Lecturer_Lecturer_Id from course_has_lecturer where Course_Course_Id=?)"
	rows, err := DB.Query(query, courseId)
	var names string
	if err != nil {
		fmt.Println(err.Error())
		return names
	}
	for rows.Next() {
		var title string
		var name string
		var surname string
		rows.Scan(&title, &name, &surname)
		names += title + ";" + name + ";" + surname + ";"
	}
	return names
}

func getAllDepartmentNames() []string {
	var departmentNames []string
	query := "SELECT Name FROM department;"
	row, err := DB.Query(query)
	if err != nil {
		fmt.Println(err.Error())
		return departmentNames
	}
	for row.Next() {
		var depName string
		row.Scan(&depName)
		departmentNames = append(departmentNames, depName)
	}
	return departmentNames
}

func getNonAttendanceOfStudent(studentId int, courseId int) int {
	query := "select Non_Attendance from department WHERE Course_Id=? and Student_Id=?"
	var nonAttendance int
	if err := DB.QueryRow(query, courseId, studentId).Scan(&nonAttendance); err != nil {
		fmt.Println(err.Error())
	}
	return nonAttendance
}

func getAllStudents() []student {
	var students []student
	query := "SELECT Student_Id,Name,Surname,Year,Department_Id,Mail,GPA,Photo_Path  FROM STUDENT"
	rows, err := DB.Query(query)
	if err != nil {
		fmt.Println(err.Error())
		return students
	}

	for i := 0; i < 100; i++ {
		rows.Next()
		var student student
		rows.Scan(&student.Id, &student.Name, &student.Surname, &student.Year, &student.DepId, &student.EMail, &student.GPA, &student.PhotoPath)
		students = append(students, student)
	}
	return students
}

func addCourseHasLecturer(lecId int, courseId int) bool {
	query := "INSERT INTO course_has_lecturer (Course_Course_Id, Lecturer_Lecturer_Id) VALUES (?, ?)"
	_, err := DB.Query(query, courseId, lecId)
	if err != nil {
		fmt.Println(err.Error())
		return false
	}
	return true
}

func addCourse(lecturer *lecturer, courseName string, attendanceLimit int, credit int) bool {
	//Check to see the course exists and is active
	queryToCheck := "select Course_Id from course where Name=? and Active=1"
	var isCourseExistId int
	DB.QueryRow(queryToCheck, courseName).Scan(&isCourseExistId)
	if isCourseExistId == 0 {
		//means the course should be created
		queryAdd := "INSERT INTO course (Name, Departmend_Ids, Attandence_Limit, Credit) VALUES (?,?,?);"
		_, err := DB.Exec(queryAdd, lecturer.DepId, attendanceLimit, credit)
		if err != nil {
			fmt.Println(err)
			return false
		}
		var courseID int
		queryToGetId := "select Course_Id from course where Name=? and Departmend_Ids=?"
		DB.QueryRow(queryToGetId, courseName, lecturer.DepId).Scan(&courseID)
		return addCourseHasLecturer(lecturer.Id, courseID)
	} else {
		fmt.Println("course already exist!")
		return false
	}

}

func createLecturer(id int, password string, title string, name string, surname string, departmentName string) bool {
	query := "INSERT INTO lecturer ('Lecturer_Id', 'Password', 'Name', 'Surname', 'Mail', 'Department_Id', 'Title') VALUES (?, ?, ?,?, ?, ?, ?)"
	mail := name + "." + surname + "@deu.edu.tr"
	success, depId := getDepIdByName(departmentName)
	if success != true {
		fmt.Println("error occured when finding the department in createLecturer function")
		return false
	}
	_, err := DB.Exec(query, id, password, name, surname, mail, depId, title)
	if err != nil {
		fmt.Println(err.Error())
		return false
	}
	return true
}
func createStudent(id int, password string, name string, surname string, departmentName string) bool {
	query := "INSERT INTO student ('Student_Id', 'Password', 'Name', 'Surname', 'Year','Mail','GPA', 'Department_Id') VALUES (?,?,?,?,?,?,?,?)"
	mail := name + "." + surname + "@ogr.deu.edu.tr"
	success, depId := getDepIdByName(departmentName)
	if success != true {
		fmt.Println("error occured when finding the department in createLecturer function")
		return false
	}
	_, err := DB.Exec(query, id, password, name, surname, 1, mail, 0, depId)
	if err != nil {
		fmt.Println(err.Error())
		return false
	}
	return true
}
func getDepIdByName(name string) (bool, int) {
	var depID int
	query := "select Department_Id from department where name=?"
	if err := DB.QueryRow(query, name).Scan(&depID); err != nil {
		fmt.Println(err.Error())
		return false, depID
	}
	return true, depID

}

func getAllLecturers() []lecturer {
	var lecturers []lecturer
	query := "SELECT Lecturer_Id,Name,Surname,Mail,Department_Id,Title,Photo_Path from FROM lecturer"
	rows, err := DB.Query(query)
	if err != nil {
		fmt.Println(err.Error())
		return lecturers
	}
	for rows.Next() {
		var lecturer lecturer
		rows.Scan(&lecturer.Id, &lecturer.Name, &lecturer.Surname, &lecturer.EMail, &lecturer.DepId, &lecturer.Title, &lecturer.PhotoPath)
		lecturers = append(lecturers, lecturer)
	}
	return lecturers
}

func getDepartmentOfStudent(id int) *department {
	var department department
	query := "SELECT Department_Id,Name,Head_Lecturer_Id FROM mdebis.student_department where Student_Id=(?)"

	if err := DB.QueryRow(query, id).Scan(&department.Id, &department.Name, &department.HeadLectId); err != nil {
		fmt.Println(err.Error())
		return nil
	}
	return &department
}

func getRealPasswordAdmin(id int) (bool, string) {
	var realPassword string
	query := "CALL manager_get_password(?)"
	if err := DB.QueryRow(query, id).Scan(&realPassword); err != nil {
		if err != nil {
			fmt.Println(err.Error())
			if err != nil {
				return false, ""
			}
			return false, ""
		}
		return false, ""
	}
	return true, realPassword
}

func deleteStudent(idStudent int) bool {
	queryDeleteCourses := "DELETE FROM course_has_student WHERE Student_Id=?"
	_, err := DB.Exec(queryDeleteCourses, idStudent)
	if err != nil {
		fmt.Println(err.Error())
		return false
	}

	query := "DELETE FROM student WHERE Student_Id=?"
	_, err = DB.Exec(query, idStudent)
	if err != nil {
		fmt.Println(err.Error())
		return false
	}
	return true
}

func deleteLecturer(idLecturer int) bool {
	query := "DELETE FROM lecturer WHERE Lecturer_Id=?"
	_, err := DB.Exec(query, idLecturer)
	if err != nil {
		fmt.Println(err.Error())
		return false
	}
	return true
}

/*
Returns the student struct with its session hash, usually called after a successful login
Also saves the student to the ACTIVE_USERS map
*/
func getStudent(id int) *student {
	var student student
	if err := DB.QueryRow("SELECT Student_Id,Name,Surname,Year,Department_Id,Mail,GPA,Photo_Path from mdebis.student where Student_Id=?", id).Scan(&student.Id, &student.Name, &student.Surname, &student.Year, &student.DepId, &student.EMail, &student.GPA, &student.PhotoPath); err != nil {
		fmt.Println(err.Error())
		return nil
	}
	return &student
}

/*
Returns the lecturer struct with its session hash, usually called after a successful login
Also saves the lecturer to the ACTIVE_USERS map
*/
func getGeneralAnnouncements() []general_announcement {
	rows, err := DB.Query("SELECT * FROM mdebis.general_announcement")
	if err != nil {
		fmt.Println(err.Error())
		return nil
	}

	var announcements []general_announcement

	// Loop through rows, using Scan to assign column data to struct fields.
	for rows.Next() {
		var announcement general_announcement
		if err := rows.Scan(&announcement.AnnouncementId, &announcement.Title, &announcement.Content, &announcement.Link); err != nil {
			fmt.Println(err.Error())
			return nil
		}
		announcements = append(announcements, announcement)
	}
	return announcements
}

func getLecturer(id int) *lecturer {
	var lecturer lecturer
	if err := DB.QueryRow("SELECT Lecturer_Id,Name,Surname,Mail,Department_Id,Title,Photo_Path from mdebis.lecturer where Lecturer_Id=?", id).Scan(&lecturer.Id, &lecturer.Name, &lecturer.Surname, &lecturer.EMail, &lecturer.DepId, &lecturer.Title, &lecturer.PhotoPath); err != nil {
		fmt.Println(err.Error())
		return nil
	}
	return &lecturer
}

func getAdmin(id int) *manager {
	var manager manager
	if err := DB.QueryRow("SELECT Manager_Id,Name,Surname,Photo_Path from mdebis.manager where Manager_Id=?", id).Scan(&manager.Id, &manager.Name, &manager.Surname, &manager.Photo_Path); err != nil {
		fmt.Println(err.Error())
		return nil
	}

	return &manager
}

/*
creates a timetable for a given student by querying the DB for the courses
*/
func getCoursesTimeTable(student *student) *[40]time_table_entry {

	courses := getCoursesOfAStudent(student.Id)
	var timeTable [40]time_table_entry
	for _, course := range courses {
		courseTime := course.Time_Inf
		var entry time_table_entry
		entry.AttandenceLimit = course.AttandenceLimit
		entry.Course_name = course.Name
		depId := course.Dep_Id
		if err := DB.QueryRow("SELECT name from mdebis.department where Department_Id=?", depId).Scan(&entry.Department); err != nil {
			fmt.Println(err.Error())
			return nil
		}
		lecturerInfos := getLecturerOfCourse(&course)
		for _, lecInfo := range lecturerInfos {
			entry.Lecturer_name = entry.Lecturer_name + " " + lecInfo + ";"
		}
		for _, time := range courseTime {
			rightIndex := ((time.Hour - 1) * 5) + time.Day
			timeTable[rightIndex-1] = entry
		}
	}
	return &timeTable
}

func canCourseClosed(courseId int) bool {
	query := "select * from course_has_student where Course_Id=? and Situtation='Current'"
	rows, err := DB.Query(query, courseId)
	if err != nil {
		fmt.Println(err.Error())
		return false
	}
	for rows.Next() {
		return false
	}
	return true

}

func changeStatusOfCourse(courseId int, isActive bool) bool {

	//check whether there is a student taking the course
	if !canCourseClosed(courseId) {
		return false
	}
	query := "UPDATE mdebis.course SET Active = ? WHERE (Course_Id = ?);"
	_, err := DB.Exec(query, isActive, courseId)
	if err != nil {
		fmt.Println(err.Error())
		return false
	}
	return true
}

/*
gets the courses that student enrolled in the current semester
*/
func getCoursesOfALecturer(lecturer *lecturer) []course {
	rowsCourses, err := DB.Query("select * from course where Course_Id IN (SELECT Course_Course_Id FROM mdebis.course_has_lecturer where Lecturer_Lecturer_Id=? and Active=1)", lecturer.Id)
	if err != nil {
		fmt.Println(err.Error())
		return nil
	}

	var courses []course

	for rowsCourses.Next() {
		//create course struct because they will also send to general course map (not created yet)
		var course course
		if err := rowsCourses.Scan(&course.Id, &course.Name, &course.Dep_Id, &course.AttandenceLimit, &course.Credit, &course.IsActive); err != nil {
			fmt.Println(err.Error())
			return nil
		}
		addTimeInfo(&course)
		course.Announcements = getAnnouncementOfCourse(course.Id)
		courses = append(courses, course)
	}
	if err = rowsCourses.Err(); err != nil {
		fmt.Println(err)
		return nil
	}
	return courses

}

func getCoursesOfAStudent(studentId int) []course {
	//GETTING COURSE IDS THAT STUDENT IS TAKING

	rowsCourses, err := DB.Query("select * from course where Course_Id IN (SELECT Course_Id FROM mdebis.course_has_student where Student_Id=? and Situtation='Current')", studentId)
	if err != nil {
		fmt.Println(err.Error())
		return nil
	}

	var courses []course

	for rowsCourses.Next() {
		//create course struct because they will also send to general course map (not created yet)
		var course course
		if err := rowsCourses.Scan(&course.Id, &course.Name, &course.Dep_Id, &course.AttandenceLimit, &course.Credit, &course.IsActive); err != nil {
			fmt.Println(err.Error())
			return nil
		}
		addTimeInfo(&course)
		course.Announcements = getAnnouncementOfCourse(course.Id)
		courses = append(courses, course)
	}
	if err = rowsCourses.Err(); err != nil {
		fmt.Println(err)
		return nil
	}
	//TODO HOMEENTRY BURAYA
	return courses
}

/*
gets the time information of a given course by querying the DB
*/
func addTimeInfo(course *course) {
	rows, err := DB.Query("SELECT Day,Hour FROM mdebis.course_time where Course_Id=?", course.Id)
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	for rows.Next() {
		var courseTime course_time
		if err := rows.Scan(&courseTime.Day, &courseTime.Hour); err != nil {
			fmt.Println(err.Error())
			return
		}
		course.Time_Inf = append(course.Time_Inf, courseTime)
	}

}

/*
gets the lecturer(s) information of a given course by querying the DB
*/
func getAnnouncementOfCourse(courseId int) []announcement {
	var announcements []announcement
	//TODO NEW TABLE CHECK
	queryGetsAnnouncements := "SELECT * FROM course_has_announcement WHERE Course_Id=?;"
	rows, err := DB.Query(queryGetsAnnouncements, courseId)
	if err != nil {
		fmt.Println(err.Error())
		return nil
	}
	for rows.Next() {
		var announcement announcement
		err := rows.Scan(&announcement.AnnouncementId, &announcement.CourseId, &announcement.Title, &announcement.Content, &announcement.LecturerId)
		if err != nil {
			return announcements
		}
		announcements = append(announcements, announcement)
	}
	return announcements
}
func getStudentsOfCourse(lecturerID, courseId int) []student {
	//CHECK IF lecturer owns the course
	if isLecturerOwnTheCourse(courseId, lecturerID) == false {
		return nil
	}
	//Check that whether the course is active or not
	queryCheckCourse := "select * from course where Course_Id=? and Active=1;"
	row, err := DB.Query(queryCheckCourse, courseId)
	if err != nil || row.Next() == false {
		print(err.Error())
		return nil
	}

	var students []student
	queryGetStudent := "SELECT Student_Id,Name,Surname,Year,Department_Id,Mail,GPA,Photo_Path  FROM student WHERE Student_Id IN" +
		"(SELECT Student_Id FROM course_has_student where Course_Id=? and Situtation='Current');"
	rowStudents, _ := DB.Query(queryGetStudent, courseId)
	for rowStudents.Next() {
		var student student
		err := rowStudents.Scan(&student.Id, &student.Name, &student.Surname, &student.Year, &student.DepId, &student.EMail, &student.GPA, &student.PhotoPath)
		if err != nil {
			return students
		}
		students = append(students, student)
	}
	return students
}

func changeNonAttendance(lecturerId int, courseId int, studentId int, nonAttendance int) bool {
	//check the lecturer owns the course
	if isLecturerOwnTheCourse(courseId, lecturerId) == false {
		return false
	}
	//NO NEED TO CHECK WHETHER THE STUDENT IS TAKING THE COURSE OR NOT
	//BECAUSE IF THE STUDENT IS NOT TAKING, THE DB WILL RETURN AN ERROR
	//AND SO THIS ERROR WILL ALSO BE RETURNED BY THIS FUNCTION :)
	queryToChange := "UPDATE course_has_student SET Non_Attendance=? where Course_Id=? and Student_Id=?;"
	_, err := DB.Exec(queryToChange, nonAttendance, courseId, studentId)
	if err != nil {
		fmt.Println(err.Error())
		return false
	}
	return true
}

func getPastCoursesOfStudent(studentId int) []course {
	//GETTING COURSE IDS THAT STUDENT IS TAKING
	rows, err := DB.Query("SELECT Course_Id FROM mdebis.course_has_student where Student_Id=? and (Situtation='Passed' or Situtation='Failed')", studentId)
	if err != nil {
		fmt.Println(err.Error())
		return nil
	}
	var courseIds []int
	for rows.Next() {
		//create course struct because they will also send to general course map (not created yet)
		var course int
		if err := rows.Scan(&course); err != nil {
			fmt.Println(err.Error())
			return nil
		}
		courseIds = append(courseIds, course)
	}
	//GETTING COURSES WITH THE GIVEN IDS
	params := make([]interface{}, 0)
	query := []string{"SELECT * FROM mdebis.course where"}
	if len(courseIds) > 0 {
		query = append(query,
			fmt.Sprintf(
				"Course_Id IN (%s)",
				strings.Join(strings.Split(strings.Repeat("?", len(courseIds)), ""), ", "),
			),
		)
	}
	for _, courseId := range courseIds {
		params = append(params, courseId)
	}
	rowsCourses, err := DB.Query(strings.Join(query, " ")+";", params...)
	if err != nil {
		fmt.Println(err.Error())
		return nil
	}

	var courses []course

	for rowsCourses.Next() {
		//create course struct because they will also send to general course map (not created yet)
		var course course
		if err := rowsCourses.Scan(&course.Id, &course.Name, &course.Dep_Id, &course.AttandenceLimit, &course.Credit, &course.IsActive); err != nil {
			fmt.Println(err.Error())
			return nil
		}
		addTimeInfo(&course)
		courses = append(courses, course)
	}
	if err = rowsCourses.Err(); err != nil {
		fmt.Println(err)
		return nil
	}
	return courses
}

func getLecturerOfCourse(course *course) []string {
	rows, err := DB.Query("SELECT Lecturer_Lecturer_Id FROM mdebis.course_has_lecturer where Course_Course_Id=?", course.Id)
	if err != nil {
		fmt.Println(err.Error())
		return nil
	}
	lecturerIds := make([]int, 0)
	for rows.Next() {
		var id int
		if err := rows.Scan(&id); err != nil {
			fmt.Println(err.Error())
			return nil
		}
		lecturerIds = append(lecturerIds, id)
	}
	lecturerInfos := make([]string, 0)
	for _, lecId := range lecturerIds {
		var title string
		var name string
		var surname string
		if err := DB.QueryRow("SELECT Title,Name,Surname from mdebis.lecturer where Lecturer_Id=?", lecId).Scan(&title, &name, &surname); err != nil {
			if err == sql.ErrNoRows {
				fmt.Println(err.Error())
				return nil
			}
			fmt.Println(err.Error())
			return nil
		}
		var lecturerInfo = title + " " + name + " " + surname
		lecturerInfos = append(lecturerInfos, lecturerInfo)
	}
	return lecturerInfos
}

func addGrade(lecturerId int, courseId int, studentId int, grade string) bool {
	//Checking whether this lecturer has this course or not
	if isLecturerOwnTheCourse(courseId, lecturerId) == false {
		return false
	}

	queryMakeUpdate := "UPDATE course_has_student SET Grade=? where Course_Id=? and Student_Id=? and Situtation='Current';"
	_, err := DB.Exec(queryMakeUpdate, grade, courseId, studentId)
	if err != nil {
		return false
	}
	return true
}

func isLecturerOwnTheCourse(courseId int, lecturerId int) bool {
	//Checking whether this lecturer has this course or not
	queryCheckingLecturerOwns := "select * from course_has_lecturer where Course_Course_Id=? and Lecturer_Lecturer_Id=?;"
	res, err := DB.Query(queryCheckingLecturerOwns, courseId, lecturerId)
	if err != nil {
		fmt.Println(err.Error())
		return false
	}
	if res.Next() == false {
		return false
	}
	return true
}

func addAnnouncement(lecturerId int, courseId int, title string, content string) bool {
	if isLecturerOwnTheCourse(courseId, lecturerId) == false {
		return false
	}
	query := "INSERT INTO course_has_announcement VALUES (0,?,?,?,?);"
	_, err := DB.Exec(query, courseId, title, content, lecturerId)
	if err != nil {
		fmt.Println(err.Error())
		return false
	}
	return true
}

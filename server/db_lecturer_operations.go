package main

import (
	"database/sql"
	"fmt"
)

/*
Returns the password taken from DB if there is match for the given id of a lecturer in the DB
*/
func getRealPasswordLecturer(id int) string {
	var realPassword string
	query := "CALL lecturer_get_password(?)"

	if err := DB.QueryRow(query, id).Scan(&realPassword); err != nil {
		fmt.Println(err.Error())
	}
	return realPassword
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
		_, err := DB.Exec(queryAdd, courseName, lecturer.DepId, attendanceLimit, credit)
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
func getLecturer(id int) *lecturer {
	var lecturer lecturer
	if err := DB.QueryRow("SELECT Lecturer_Id,Name,Surname,Mail,Department_Id,Title,Photo_Path from mdebis.lecturer where Lecturer_Id=?", id).Scan(&lecturer.Id, &lecturer.Name, &lecturer.Surname, &lecturer.EMail, &lecturer.DepId, &lecturer.Title, &lecturer.PhotoPath); err != nil {
		fmt.Println(err.Error())
		return nil
	}
	return &lecturer
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
func changeStatusOfCourse(courseId int) bool {

	//check whether there is a student taking the course
	if !canCourseClosed(courseId) {
		return false
	}
	query := "UPDATE mdebis.course SET Active = 0 WHERE (Course_Id = ?);"
	_, err := DB.Exec(query, courseId)
	if err != nil {
		fmt.Println(err.Error())
		return false
	}
	return true
}
func getStudentsOfCourse(lecturerID, courseId int) []studentOfCourse {
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

	var studentOfCourses []studentOfCourse
	queryGetStudent := "SELECT Student_Id,Name,Surname,Year,Department_Id,Mail,GPA,Photo_Path  FROM student WHERE Student_Id IN" +
		"(SELECT Student_Id FROM course_has_student where Course_Id=? and Situtation='Current');"
	rowStudents, _ := DB.Query(queryGetStudent, courseId)
	for rowStudents.Next() {
		var studentOfCourse studentOfCourse
		var student student
		err := rowStudents.Scan(&student.Id, &student.Name, &student.Surname, &student.Year, &student.DepId, &student.EMail, &student.GPA, &student.PhotoPath)
		if err != nil {
			return studentOfCourses
		}
		var nonAttendance = getNonAttendanceOfStudent(student.Id, courseId)
		if nonAttendance == -1 {
			fmt.Println("error when finding non attendance of the student!")
			return studentOfCourses
		}
		studentOfCourse.Student = student
		studentOfCourse.NonAttendance = nonAttendance
		studentOfCourses = append(studentOfCourses, studentOfCourse)
	}
	return studentOfCourses
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

	queryMakeUpdate := "UPDATE course_has_student SET Grade=?,Situtation='Passed' where Course_Id=? and Student_Id=? and Situtation='Current';"
	if grade == "FF" {
		queryMakeUpdate = "UPDATE course_has_student SET Grade=?, Situtation='Failed' where Course_Id=? and Student_Id=? and Situtation='Current';"

	}
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

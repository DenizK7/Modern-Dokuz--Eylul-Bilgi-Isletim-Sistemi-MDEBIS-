package main

import "fmt"

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
func getNonAttendanceOfStudent(studentId int, courseId int) int {
	query := "select Non_Attendance from course_has_student WHERE Course_Id=? and Student_Id=?"
	var nonAttendance int
	if err := DB.QueryRow(query, courseId, studentId).Scan(&nonAttendance); err != nil {
		fmt.Println(err.Error())
	}
	return nonAttendance
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
	return courses
}
func getPastCoursesOfStudent(studentId int) []PastCourse {
	//GETTING COURSES WITH THE GIVEN IDS
	query := "SELECT * FROM mdebis.course where " +
		"Course_Id In (SELECT Course_Id FROM mdebis.course_has_student where Student_Id=? and (Situtation='Passed' or Situtation='Failed'))"
	rowsCourses, err := DB.Query(query, studentId)
	if err != nil {
		fmt.Println(err.Error())
		return nil
	}

	var pastCourses []PastCourse

	for rowsCourses.Next() {
		//create course struct because they will also send to general course map (not created yet)
		var course course
		if err := rowsCourses.Scan(&course.Id, &course.Name, &course.Dep_Id, &course.AttandenceLimit, &course.Credit, &course.IsActive); err != nil {
			fmt.Println(err.Error())
			return nil
		}
		addTimeInfo(&course)
		var pastCourse PastCourse
		pastCourse.Course = course
		pastCourse.Grade = getGrade(studentId, course.Id)
		pastCourses = append(pastCourses, pastCourse)
	}
	if err = rowsCourses.Err(); err != nil {
		fmt.Println(err)
		return pastCourses
	}
	return pastCourses
}
func getGrade(idStudent int, idCourse int) string {
	query := "select Grade from course_has_student where Student_Id=? and Course_Id=?"
	var grade string
	if err := DB.QueryRow(query, idStudent, idCourse).Scan(&grade); err != nil {
		return "N/A"
	}
	if grade == "FF" || grade == "" {
		grade = "FF | Failed"
	}
	return grade

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

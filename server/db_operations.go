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

	courses := getCoursesOfAStudent(student)
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

func changeStatusOfCourse(courseId int, isActive bool) {
	query := "UPDATE mdebis.course SET Active = ? WHERE (Course_Id = ?);"
	_, err := DB.Exec(query, isActive, courseId)
	if err != nil {
		fmt.Println(err.Error())
		return
	}
}

/*
gets the courses that student enrolled in the current semester
*/
func getCoursesOfALecturer(lecturer *lecturer) []course {
	query := "SELECT Course_Course_Id FROM mdebis.course_has_lecturer where Lecturer_Lecturer_Id=? "
	rows, err := DB.Query(query, lecturer.Id)
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
	query_2 := []string{"SELECT * FROM mdebis.course where"}
	if len(courseIds) > 0 {
		query_2 = append(query_2,
			fmt.Sprintf(
				"Course_Id IN (%s)",
				strings.Join(strings.Split(strings.Repeat("?", len(courseIds)), ""), ", "),
			),
		)
	}
	for _, courseId := range courseIds {
		params = append(params, courseId)
	}
	rowsCourses, err := DB.Query(strings.Join(query_2, " ")+";", params...)
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

func getCoursesOfAStudent(student *student) []course {
	//GETTING COURSE IDS THAT STUDENT IS TAKING
	rows, err := DB.Query("SELECT Course_Id FROM mdebis.course_has_student where Student_Id=? and Situtation='Current'", student.Id)
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
		if err := rowsCourses.Scan(&course.Id, &course.Name, &course.Dep_Id, &course.AttandenceLimit, &course.Credit); err != nil {
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

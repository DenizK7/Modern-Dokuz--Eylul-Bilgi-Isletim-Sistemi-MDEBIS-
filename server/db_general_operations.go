package main

import "fmt"

func getDepartmentName(depId int) string {
	query := "select Name from department WHERE Department_Id=?"
	var depName string
	if err := DB.QueryRow(query, depId).Scan(&depName); err != nil {
		fmt.Println(err.Error())
	}
	return depName
}

func getLecturerNamesOfCourse(courseId int) []string {
	query := "select Title,Name,Surname from lecturer where Lecturer_Id IN " +
		"(select Lecturer_Lecturer_Id from course_has_lecturer where Course_Course_Id=?)"
	rows, err := DB.Query(query, courseId)
	var names []string
	if err != nil {
		fmt.Println(err.Error())
		return names
	}
	for rows.Next() {
		var title string
		var name string
		var surname string
		rows.Scan(&title, &name, &surname)
		names = append(names, title+" "+name+" "+surname)
	}
	return names
}

func addLog(whoTypeDid string, whoDidId int, operation string, whichTable string, values string) bool {
	SQL := "INSERT LOG VALUES(0,?,?,?,?,?)"
	_, err := DB.Exec(SQL, whoTypeDid, whoDidId, operation, whichTable, values)
	if err != nil {
		fmt.Println(err.Error())
		return false
	}
	return true
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

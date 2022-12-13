package main

import "database/sql"

type course_status int

const (
	Failed course_status = iota
	Current
	Passed
)

type user struct {
	Student  *student
	Lecturer *lecturer
	Manager  *manager
}

type student struct {
	Id        int
	Name      string
	Surname   string
	Year      int
	DepId     int
	GPA       float32
	EMail     string
	PhotoPath sql.NullString
}

type lecturer struct {
	Id        string
	Name      string
	Surname   string
	EMail     string
	DepId     int
	Title     string
	PhotoPath sql.NullString
}
type general_announcement struct {
	AnnouncementId int
	Title          string
	Content        string
	Link           string
}
type announcement struct {
	AnnouncementId int
	CourseId       int
	Title          string
	Content        string
}

type course struct {
	Id              int
	Name            string
	Dep_Id          int
	AttandenceLimit int
	Time_Inf        []course_time
	Credit          int
	Announcements   []announcement
}

type department struct {
	Id         int
	HeadLectId int
	Name       string
}

type manager struct {
	Id           int
	Name         string
	Surname      string
	Photo_Path   sql.NullString
	SessionToken string
}
type time_table_entry struct {
	Department      string
	Course_name     string
	Lecturer_name   string
	AttandenceLimit int
}
type course_time struct {
	Day  int
	Hour int
}

package main

import "github.com/gorilla/mux"

func Router() *mux.Router {

	router := mux.NewRouter()
	router.HandleFunc("/get_gen_announcements", responseGetGeneralAnnouncements) //returns all general announcements

	router.HandleFunc("/log_student/{username}/{password}", responseStudentLogIn)   //returns session hash if successful, false otherwise
	router.HandleFunc("/log_lecturer/{username}/{password}", responseLecturerLogIn) //returns session hash if successful, false otherwise
	router.HandleFunc("/log_manager/{id}/{password}", responseAdminLogIn)           //returns session hash if successful, false otherwise

	//Lecturer requests

	//Student requests
	router.HandleFunc("/get_courses/{sessionHash}", responseGetCourses)  //return courses if given hash is correct, false otherwise
	router.HandleFunc("/time_table/{sessionHash}", responseGetTimeTable) //returns timetable if given hash is correct, false otherwise

	//Admin requests

	return router
}

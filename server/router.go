package main

import "github.com/gorilla/mux"

func Router() *mux.Router {

	router := mux.NewRouter()
	router.HandleFunc("/get_gen_announcements", responseGetGeneralAnnouncements) //returns all general announcements
	//below requests are same for students and lecturers (POLYMORPHISM)
	router.HandleFunc("/get_home_entry/{sessionHash}", responseGetHomeEntry) //return courses if given hash is correct, false otherwise
	router.HandleFunc("/log_out/{sessionHash}", responseLogOut)
	//Lecturer requests
	router.HandleFunc("/log_lecturer/{username}/{password}", responseLecturerLogIn) //returns session hash if successful, false otherwise
	router.HandleFunc("/change_course_status/{sessionHash}/{courseId}", responseChangeActiveOfCourse)
	router.HandleFunc("/add_grade/{sessionHash}/{courseId}/{studentId}/{grade}", responseAddGrade)
	router.HandleFunc("/add_announcement/{sessionHash}/{courseId}/{title}/{content}", responseAddAnnouncement)
	router.HandleFunc("/get_student_of_course/{sessionHash}/{courseId}", responseGetStudentsOfCourse)                             //
	router.HandleFunc("/change_non_attendance/{sessionHash}/{courseId}/{studentId}/{nonAttendance}", responseChangeNonAttendance) //
	router.HandleFunc("/add_course/{sessionHash}/{courseName}/{attendanceLimit}/{credit}", responseAddCourse)                     //

	//Student requests
	router.HandleFunc("/log_student/{username}/{password}", responseStudentLogIn)                  //returns session hash if successful, false otherwise
	router.HandleFunc("/time_table/{sessionHash}", responseGetTimeTable)                           //returns timetable if given hash is correct, false otherwise
	router.HandleFunc("/get_department_of_student/{sessionToken}", responseGetDepartmentOfStudent) // responseGetPastCoursesOfStudent
	router.HandleFunc("/get_past_courses/{sessionToken}", responseGetPastCoursesOfStudent)

	//Admin requests //responseDeleteStudent
	router.HandleFunc("/log_admin/{username}/{password}", responseAdminLogIn) //returns session hash if successful, false otherwise
	router.HandleFunc("/delete_student/{sessionHash}/{studentId}", responseDeleteStudent)
	router.HandleFunc("/delete_lecturer/{sessionHash}/{lecturerId}", responseDeleteLecturer)
	router.HandleFunc("/get_students/{sessionHash}", responseGetStudents)
	router.HandleFunc("/get_lecturers/{sessionHash}", responseGetLecturers)
	router.HandleFunc("/create_lecturer/{sessionHash}/{id}/{password}/{title}/{name}/{surname}/{departmentName}", responseCreateLecturer)
	router.HandleFunc("/create_student/{sessionHash}/{id}/{password}/{name}/{surname}/{departmentName}", responseCreateStudent)
	router.HandleFunc("/get_all_department_names/{sessionHash}", responseGetAllDepartmentNames)

	return router
}

package main

import "fmt"

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
func getAllStudents() []student {
	var students []student
	query := "SELECT Student_Id,Name,Surname,Year,Department_Id,Mail,GPA,Photo_Path  FROM STUDENT where isActive=1"
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
func getLog() []logDB {
	var logRecords []logDB
	query := "select * from log"
	rows, err := DB.Query(query)
	if err != nil {
		fmt.Println(err.Error())
		return logRecords
	}
	for rows.Next() {
		var logRecord logDB
		rows.Scan(&logRecord.RecordId, &logRecord.WhoDid, &logRecord.WhoDidId, &logRecord.Operation, &logRecord.WhichTable, &logRecord.Values)
		logRecords = append(logRecords, logRecord)
	}
	return logRecords

}
func createLecturer(id int, password string, title string, name string, surname string, departmentName string) bool {
	query := "INSERT INTO lecturer (Lecturer_Id, Password, Name, Surname, Department_Id, Title) VALUES (?,?,?,?,?,?)"
	success, depId := getDepIdByName(departmentName)
	if success != true {
		fmt.Println("error occured when finding the department in createLecturer function")
		return false
	}
	_, err := DB.Exec(query, id, password, name, surname, depId, title)
	if err != nil {
		fmt.Println(err.Error())
		return false
	}
	return true
}
func createStudent(id int, password string, name string, surname string, departmentName string) bool {
	query := "INSERT INTO student (Student_Id, Password, Name, Surname, Year,GPA, Department_Id) VALUES (?,?,?,?,?,?,?)"
	success, depId := getDepIdByName(departmentName)
	if success != true {
		fmt.Println("error occured when finding the department in createLecturer function")
		return false
	}
	_, err := DB.Exec(query, id, password, name, surname, 1, 0, depId)
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
func makeFailAllCoursesOfStudent(id int) bool {
	query := "update course_has_student set Situtation='Failed' where Student_Id=?"
	_, err := DB.Exec(query, id)
	if err != nil {
		fmt.Println(err.Error())
		return false
	}
	return true
}
func deleteStudent(idStudent int) bool {
	if makeFailAllCoursesOfStudent(idStudent) == false {
		return false
	}

	query := "UPDATE student SET isActive=0 WHERE Student_Id=?"
	_, err := DB.Exec(query, idStudent)
	if err != nil {
		fmt.Println(err.Error())
		return false
	}
	return true
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
func getAllLecturers() []lecturer {
	var lecturers []lecturer
	query := "SELECT Lecturer_Id,Name,Surname,Mail,Department_Id,Title,Photo_Path FROM lecturer where isActive=1"
	rows, err := DB.Query(query)
	if err != nil {
		fmt.Println(err.Error())
		return lecturers
	}
	i := 0
	for rows.Next() {
		if i == 100 {
			break
		}
		var lecturer lecturer
		rows.Scan(&lecturer.Id, &lecturer.Name, &lecturer.Surname, &lecturer.EMail, &lecturer.DepId, &lecturer.Title, &lecturer.PhotoPath)
		lecturers = append(lecturers, lecturer)
		i = i + 1
	}
	return lecturers
}
func deleteLecturer(idLecturer int) bool {

	//First, delete its courses
	if makeFailAllStudentsOfLecturer(idLecturer) == false {
		return false
	}
	query := "UPDATE lecturer set isActive=0 WHERE Lecturer_Id=?"
	_, err := DB.Exec(query, idLecturer)
	if err != nil {
		fmt.Println(err.Error())
		return false
	}
	return true
}
func getAdmin(id int) *manager {
	var manager manager
	if err := DB.QueryRow("SELECT Manager_Id,Name,Surname,Photo_Path from mdebis.manager where Manager_Id=?", id).Scan(&manager.Id, &manager.Name, &manager.Surname, &manager.Photo_Path); err != nil {
		fmt.Println(err.Error())
		return nil
	}

	return &manager
}

/*This function makes fail all students in the courses that currently giving by the given lecturer
 */
func makeFailAllStudentsOfLecturer(lecId int) bool {
	query := "UPDATE course_has_student set Situtation='Failed' where  Situtation='Current' and " +
		"Course_Id IN (select Course_Course_Id from course_has_lecturer where Lecturer_Lecturer_Id=?)"
	_, err := DB.Exec(query, lecId)
	if err != nil {
		fmt.Println(err.Error())
		return false
	}
	return true
}

package main

import (
	"bufio"
	"database/sql"
	"fmt"
	"io"
	"net/http"
	"os"
	"time"

	_ "github.com/go-sql-driver/mysql"
)

func main() {
	DB, errD := sql.Open("mysql", "root:354152@tcp(127.0.0.1:3306)/mdebis")
	if DB != nil {

	}
	if errD != nil {
		fmt.Println(errD.Error())
	}
	/*
		f, err := os.Create("student_ids.txt")

		if err != nil {
			log.Fatal(err)
		}

		queryGetStudents := "select Student_Id from mdebis.student;"
		row, err := DB.Query(queryGetStudents)
		if err != nil {
			fmt.Println(err.Error())
		}
		for row.Next() {
			var id string
			row.Scan(&id)
			_, err2 := f.WriteString(id + "\n")

			if err2 != nil {
				log.Fatal(err2)
			}
		}
	*/
	start := time.Now()
	ch := make(chan string)
	readFile, _ := os.Open("student_ids.txt")

	fileScanner := bufio.NewScanner(readFile)
	fileScanner.Split(bufio.ScanLines)
	limit := 1000
	counter := 0
	for fileScanner.Scan() {
		id := fileScanner.Text()
		go MakeRequest("http://localhost:8080/log_student/"+id+"/354152", ch)
		counter = counter + 1
		if counter > limit {
			break
		}
	}

	readFile.Close()
	for i := 0; i < limit; i++ {
		fmt.Println(<-ch)
	}
	fmt.Printf("%.2fs elapsed\n", time.Since(start).Seconds())
}
func MakeRequest(url string, ch chan<- string) {
	resp, _ := http.Get(url)
	body, _ := io.ReadAll(resp.Body)
	sessionToken := string(body)

	//now get time table of the student
	resp2, _ := http.Get("http://localhost:8080/time_table/" + sessionToken)
	body2, _ := io.ReadAll(resp2.Body)
	fmt.Println(string(body2))

}

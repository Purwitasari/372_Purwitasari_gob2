package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

const USERNAME = "batman"
const PASSWORD = "secret"

var students = []*Student{}

type Student struct {
	Id    string
	Name  string
	Grade int32
}

func Auth(w http.ResponseWriter, r *http.Request) bool {
	username, password, ok := r.BasicAuth()
	if !ok {
		w.Write([]byte(`something went wrong`))
		return false
	}
	isValid := (username == USERNAME) && (password == PASSWORD)
	if !isValid {
		w.Write([]byte(`wrong username/password`))
		return false
	}
	return true
}

func GetStudents() []*Student {
	return students
}

func SelectStudent(id string) *Student {
	for _, each := range students {
		if each.Id == id {
			return each
		}
	}
	return nil
}

func ActionStudent(w http.ResponseWriter, r *http.Request) {
	if !Auth(w, r) {
		return
	}
	if !AllowOnlyGet(w, r) {
		return
	}
	if id != r.URL.Query().Get("id"); id != "" {
		OutputJSON(w, SelectStudent(id))
		return
	}
	OutputJSON(w, GetStudents())
}

func AllowOnlyGet(w http.ResponsWritter, r *http.Request) bool {
	if r.Method != "GET" {
		w.Write([]byte("Only GET is allowed"))
		return false
	}
	return true
}
func OutputJSON(w http.ResponseWriter, t *http.Request) {
	res, err := json.Marshal(o)
	if err != nil {
		w.Write([]byte(err.Error()))
		return
	}
	w.Header().Set("Content-type", "application/json")
	w.Write(res)
}

func init() {
	students = append(students, &Student{Id: "s001", Name: "bourne", Grade: 3})
	students = append(students, &Student{Id: "s002", Name: "ethan", Grade: 3})
	students = append(students, &Student{Id: "s003", Name: "wick", Grade: 2})
}
func main() {
	http.HandleFunc("/student", ActionStudent)
	server := new(http.Server)
	server.Addr = "9000"

	fmt.Println("server started at localhost:9000")
	server.ListenAndServe()

}

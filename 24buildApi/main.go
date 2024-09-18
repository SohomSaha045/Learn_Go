package main

import (
	"encoding/json"
	"fmt"
	"math/rand"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

// model for course - file
type Course struct {
	CourseId    string  `json:"courseid"`
	CourseName  string  `json:"coursename"`
	CrousePrice int     `json:"price"`
	Author      *Author `json:"author"`
}
type Author struct {
	Fullname string `json:"fullname"`
	Website  string `json:"website"`
}

// fake DB
var courses []Course

// middleware
func IsEmpty(c *Course) bool {
	// return c.CourseId == "" && c.CourseName == ""
	return c.CourseName == ""

}

//CONTROLERS - file

// serveHome
func serveHome(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("<h1>Welcome to GO API</h1>"))
}

// Get ALL
func getAll(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Get all Courses")
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(courses)
}

// Get with ID
func getOneCourse(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Get One Course")
	w.Header().Set("Content-Type", "application/json")
	//grab the ID from request
	params := mux.Vars(r)
	fmt.Println(params)
	//loop through the course and find matching ID
	for _, course := range courses {
		if course.CourseId == params["id"] {
			json.NewEncoder(w).Encode(course)
			return
		}
	}
	json.NewEncoder(w).Encode("No Course found with given id")

}

// post to database
func createCourse(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Create a Course")
	w.Header().Set("Content-Type", "application/json")

	//body is empty?
	if r.Body == nil {
		json.NewEncoder(w).Encode("Please send some data")
		return
	}
	// empty data - {}
	var course Course
	_ = json.NewDecoder(r.Body).Decode(&course)
	if IsEmpty(&course) {
		json.NewEncoder(w).Encode("No data found")
		return
	}
	//generate  uniqueID, string
	//append course into courses

	// rand.Seed(time.Now().UnixNano())
	course.CourseId = strconv.Itoa(rand.Intn(10000))
	courses = append(courses, course)
	json.NewEncoder(w).Encode(course)
	return
}

func updateCourse(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Update a Course")
	w.Header().Set("Content-Type", "application/json")

	// Grab id from request
	params := mux.Vars(r)
	// loop , id , remove , add new value
	var temp Course
	for index, course := range courses {
		if course.CourseId == params["id"] {
			courses = append(courses[:index], courses[index+1:]...)
			json.NewDecoder(r.Body).Decode(&temp)
			temp.CourseId = params["id"]
			courses = append(courses, temp)
			json.NewEncoder(w).Encode(courses)
			return
		}
	}
	//send a response when id not found
	json.NewEncoder(w).Encode("No Course found with given id")
	return
}
func deleteCourse(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Delte a Course")
	w.Header().Set("Content-Type", "application/json")

	// Grab id from request
	params := mux.Vars(r)
	// loop , id , remove , add new value
	for index, course := range courses {
		if course.CourseId == params["id"] {
			courses = append(courses[:index], courses[index+1:]...)
			json.NewEncoder(w).Encode(courses)
			return
		}
	}
	//send a response when id not found
	json.NewEncoder(w).Encode("No Course found with given id")
	return
}

func main() {
	fmt.Println("Build an API")
}

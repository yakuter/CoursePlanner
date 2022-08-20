package endpoint

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
)

type Courses struct {
	Course []Course `json:"course"`
}

type Course struct {
	Name string `json:"name"`
	Days Days   `json:"days"`
}

type Days struct {
	Monday    []int `json:"monday"`
	Tuesday   []int `json:"tuesday"`
	Wednesday []int `json:"wednesday"`
	Thursday  []int `json:"thursday"`
	Friday    []int `json:"friday"`
}

func PostCourse(w http.ResponseWriter, r *http.Request) {
	//Post an array of courses
	//WARNING: Overwrites files
	course := Courses{ //default values for course array
		Course: []Course{
			{
				Name: "default",
				Days: Days{
					Monday:    []int{0, 0, 0, 0, 0, 0, 0, 0, 0, 0},
					Tuesday:   []int{0, 0, 0, 0, 0, 0, 0, 0, 0, 0},
					Wednesday: []int{0, 0, 0, 0, 0, 0, 0, 0, 0, 0},
					Thursday:  []int{0, 0, 0, 0, 0, 0, 0, 0, 0, 0},
					Friday:    []int{0, 0, 0, 0, 0, 0, 0, 0, 0, 0},
				},
			},
		},
	}

	bodyBytes, _ := ioutil.ReadAll(r.Body)
	err := json.Unmarshal(bodyBytes, &course)
	if err != nil {
		log.Fatal(err)
	}

	defer r.Body.Close()

	fmt.Fprintf(w, "Added to JSON file\n")

	//Display added array and its contents
	for i := 0; i < len(course.Course); i++ {
		fmt.Fprintf(w, "Name = %s\n", course.Course[i].Name)
		fmt.Fprintf(w, "Monday = %d\n", course.Course[i].Days.Monday)
		fmt.Fprintf(w, "Tuesday = %d\n", course.Course[i].Days.Tuesday)
		fmt.Fprintf(w, "Wednesday = %d\n", course.Course[i].Days.Wednesday)
		fmt.Fprintf(w, "Thursday = %d\n", course.Course[i].Days.Thursday)
		fmt.Fprintf(w, "Friday = %d\n", course.Course[i].Days.Friday)
	}

	file, _ := json.MarshalIndent(course, "", " 	")
	_ = ioutil.WriteFile("courses.json", file, 0644)
}

func ViewCourses(w http.ResponseWriter, r *http.Request) {
	fileContent, err := os.ReadFile("courses.json")

	if err != nil {
		log.Fatal(err)
		return
	}

	fmt.Fprintln(w, "File successfully opened")

	var course Courses
	json.Unmarshal(fileContent, &course)

	//Display added array and its contents
	for i := 0; i < len(course.Course); i++ {
		fmt.Fprintf(w, "Name = %s\n", course.Course[i].Name)
		fmt.Fprintf(w, "Monday = %d\n", course.Course[i].Days.Monday)
		fmt.Fprintf(w, "Tuesday = %d\n", course.Course[i].Days.Tuesday)
		fmt.Fprintf(w, "Wednesday = %d\n", course.Course[i].Days.Wednesday)
		fmt.Fprintf(w, "Thursday = %d\n", course.Course[i].Days.Thursday)
		fmt.Fprintf(w, "Friday = %d\n", course.Course[i].Days.Friday)
	}
}

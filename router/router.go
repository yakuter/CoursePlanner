package router

import (
	"program/endpoint"

	"github.com/gorilla/mux"
)

func RegisterRoutes(router *mux.Router) { // Registers routes and CRUD operations to URL's
	router.HandleFunc("/addcourse", endpoint.PostCourse).Methods("POST")
	router.HandleFunc("/courses", endpoint.ViewCourses).Methods("GET")
}

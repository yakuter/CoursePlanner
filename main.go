package main

import (
	"program/db"
	"program/router"

	"github.com/gorilla/mux"
)

func main() {
	rt := mux.NewRouter()
	router.RegisterRoutes(rt)
	db.ConnectDatabase(rt)
}

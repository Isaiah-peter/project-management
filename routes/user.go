package routes

import (
	"project-management/controllers"

	"github.com/gorilla/mux"
)

var UserRoute = func(router *mux.Router) {
	router.HandleFunc("/rgister", controllers.NewUser).Methods("POST")
	router.HandleFunc("/login", controllers.Login).Methods("POST")
	router.HandleFunc("/getalluser", controllers.GetUser).Methods("GET")
	router.HandleFunc("/getuser/{id}", controllers.GetSingleUser).Methods("GET")
}

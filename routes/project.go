package routes

import (
	"project-management/controllers"

	"github.com/gorilla/mux"
)

var Project = func(router *mux.Router) {
	router.HandleFunc("/create_project", controllers.CreateProject).Methods("POST")
	router.HandleFunc("/get_all_project", controllers.GetProject).Methods("GET")
	router.HandleFunc("/project/singleproject/{id}", controllers.GetProjectById)
	router.HandleFunc("/project/user/singleproject/{id}", controllers.GetProjectByUserId).Methods("GET")
	router.HandleFunc("/project/update/{id}", controllers.UpdateProject).Methods("PUT")
	router.HandleFunc("/project/delete/{id}", controllers.DeleteProject).Methods("DELETE")
}

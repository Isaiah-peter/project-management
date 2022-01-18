package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"project-management/routes"

	"github.com/gorilla/mux"
)

func main() {
	port := os.Getenv("PORT")
	if port == "" {
		port = ":5000"
	}
	fmt.Println(port)
	router := mux.NewRouter()
	routes.UserRoute(router)

	http.Handle("/", router)
	log.Fatal(http.ListenAndServe(port, router))
}

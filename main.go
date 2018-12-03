package main

import (
	"log"
	"net/http"
	"os"

	"github.com/LordRahl90/contacts_manager/app"
	"github.com/LordRahl90/contacts_manager/controllers"
	"github.com/gorilla/mux"
)

func main() {
	router := mux.NewRouter()
	// router.Use(middlewares.Logger)
	router.Use(app.JwtAuthentication)

	router.HandleFunc("/api/user/register", controllers.CreateAccount).Methods("POST")
	router.HandleFunc("/api/user/login", controllers.Authenticate).Methods("POST")
	router.HandleFunc("/api/contacts/new", controllers.CreateContact).Methods("POST")
	router.HandleFunc("/api/me/contacts", controllers.GetUserContact).Methods("GET")

	port := os.Getenv("PORT")
	if port == "" {
		port = "8000"
	}

	log.Fatal(http.ListenAndServe(":"+port, router))
}

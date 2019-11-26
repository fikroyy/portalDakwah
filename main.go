package main

import (
	"github.com/gorilla/mux"
	app "portal_dakwah/app"
	con "portal_dakwah/controllers"
	"os"
	"fmt"
	"net/http"
)

func main() {

	router := mux.NewRouter()
	router.Use(app.JwtAuthentication) //attach JWT auth middleware

	router.HandleFunc("/api/user/register", con.CreateAccount).Methods("POST")
	router.HandleFunc("/api/user/login", con.Authenticate).Methods("POST")

	port := os.Getenv("PORT") //Get port from .env file, we did not specify any port so this should return an empty string when tested locally
	if port == "" {
		port = "8000" //localhost
	}

	fmt.Println(port)

	err := http.ListenAndServe(":" + port, router) //Launch the app, visit localhost:8000/api
	if err != nil {
		fmt.Print(err)
	}
}
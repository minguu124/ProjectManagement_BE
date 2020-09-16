package main

import (
	"Projectmanagement_BE/apps"
	"fmt"
	"net/http"
	"os"

	"github.com/gorilla/mux"
)

func main() {
	router := mux.NewRouter()

	// attach JWT auth middleware
	router.Use(apps.JwtAuthenticaion)

	//router.Use()
	// default api
	router.Handle("/", nil)

	// users api
	router.Handle("/user", notImplement)

	// projects api
	router.Handle("/project", notImplement)

	// logs api
	router.Handle("/log", notImplement)

	port := os.Getenv("PORT")
	if port == "" {
		port = "8000"
	}

	err := http.ListenAndServe(":"+port, router)
	if err == nil {
		fmt.Println(err)
	}
}

// in case api is not implemented yet
var notImplement = http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Not implemented"))
})

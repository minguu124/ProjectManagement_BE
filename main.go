package main

import (
	"fmt"
	"net/http"
)

func main() {
	http.HandleFunc("/", handler)

	http.ListenAndServe("5432", nil)
}

func handler(w http.ResponseWriter, r *http.Request) {
	fmt.Printf("Hello!\n")
}

package main

import (
	"fmt"
	"net/http"
)

func HandlerFunc(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "<h1>Welcome to my awesome site!</h1>")
}

func main() {
	http.HandleFunc("/", HandlerFunc)
	fmt.Println("Starting the server on :8080...")
	http.ListenAndServe(":8080", nil)
}

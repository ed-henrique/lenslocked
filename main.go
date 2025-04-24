package main

import (
	"fmt"
	"log/slog"
	"net/http"
)

func HandlerFunc(w http.ResponseWriter, r *http.Request) {
	_, err := fmt.Fprintf(w, "<h1>Welcome to my awesome site!</h1>")
	if err != nil {
		slog.Error("the message could not be sent", slog.String("err", err.Error()))
	}
}

func main() {
	http.HandleFunc("/", HandlerFunc)
	fmt.Println("Starting the server on :8080...")
	if err := http.ListenAndServe(":8080", nil); err != nil {
		slog.Error("could not start the server", slog.String("err", err.Error()))
	}
}

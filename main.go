package main

import (
	"fmt"
	"log/slog"
	"net/http"
	"os"
)

var log *slog.Logger

func init() {
	log = slog.New(slog.NewTextHandler(os.Stdout, nil))
}

func handler(w http.ResponseWriter, r *http.Request) {
	// Set the response header
	w.Header().Set("Content-Type", "text/plain")

	log.Info("request", "method", r.Method, "path", r.RequestURI)
	// Check the request method
	switch r.Method {
	case http.MethodGet, http.MethodPost:
		// Respond with "Hodor" for both GET and POST requests
		fmt.Fprint(w, "hodor... hodor... hodor")
	default:
		// Respond with a 405 Method Not Allowed for other request methods
		http.Error(w, "Method Not Allowed", http.StatusMethodNotAllowed)
	}
}

func main() {
	// Register the handler function for all routes
	http.HandleFunc("/", handler)

	// Start the server on port 8080
	log.Info("Server is running on http://localhost:8080")
	err := http.ListenAndServe("localhost:8080", nil)
	if err != nil {
		fmt.Println("Error:", err)
	}
}

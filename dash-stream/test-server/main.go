package main

import (
	"fmt"
	"net/http"
)

func main() {
	// Define a handler function to serve HTTP requests
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		// Serve files from the local directory
		w.Header().Set("Access-Control-Allow-Origin", "*")
		http.ServeFile(w, r, r.URL.Path[1:])
	})

	// Start the server on port 8000
	fmt.Println("Server listening on port 8000...")
	http.ListenAndServe(":8000", nil)
}

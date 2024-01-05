package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

func main() {
	// Define the endpoint "/hello" handler function
	http.HandleFunc("/hello", func(w http.ResponseWriter, r *http.Request) {
		// Parse the JSON request
		var requestBody map[string]int
		err := json.NewDecoder(r.Body).Decode(&requestBody)
		if err != nil {
			http.Error(w, "Error decoding JSON request", http.StatusBadRequest)
			return
		}

		// Get the "id" from the request
		id, ok := requestBody["id"]
		if !ok {
			http.Error(w, "Missing 'id' in JSON request", http.StatusBadRequest)
			return
		}
		fmt.Println("got request with id ", id)

		// Create the JSON response
		if id == 1 || id == 5 {
			http.Error(w, "got error ", http.StatusForbidden)
			return
		}
		response := map[string]int{"id": id}

		// Encode and send the JSON response
		w.Header().Set("Content-Type", "application/json")
		err = json.NewEncoder(w).Encode(response)
		if err != nil {
			http.Error(w, "Error encoding JSON response", http.StatusInternalServerError)
			return
		}
	})

	// Start the HTTP server
	port := 8080
	fmt.Printf("Server listening on :%d\n", port)
	err := http.ListenAndServe(fmt.Sprintf(":%d", port), nil)
	if err != nil {
		fmt.Println("Error starting the server:", err)
	}
}

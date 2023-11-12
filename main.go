package main

import (
    "encoding/json"
    "log"
    "net/http"
    "github.com/gorilla/mux"
		"fmt"
)

// Message struct to send as JSON response
type Message struct {
    Text string `json:"text"`
}

// handler function to respond to the root URL
func rootHandler(w http.ResponseWriter, r *http.Request) {
    msg := Message{Text: "Hello, world, UPDATED!!!!!"}
    json.NewEncoder(w).Encode(msg)
}

func main() {
    // Create a new router
    router := mux.NewRouter()

    // Register the handler function for the root URL
    router.HandleFunc("/", rootHandler).Methods("GET")

    // Start the server on localhost port 8080 and log any errors
		fmt.Println("Hello, world!")
    log.Println("API is running on http://localhost:3000")
    log.Fatal(http.ListenAndServe(":3000", router))
}
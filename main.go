package main

import (
    "encoding/json"
    "log"
    "net/http"
    "github.com/gorilla/mux"
)

// Message struct to send as JSON response
type Message struct {
    Text string `json:"text"`
}

// handler function to respond to the root URL
func rootHandler(w http.ResponseWriter, r *http.Request) {
    msg := Message{Text: "Hello, world!"}
    json.NewEncoder(w).Encode(msg)
}

func main() {
    // Create a new router
    router := mux.NewRouter()

    // Register the handler function for the root URL
    router.HandleFunc("/", rootHandler).Methods("GET")

    // Start the server on localhost port 8080 and log any errors
    log.Println("API is running on http://localhost:8080")
    log.Fatal(http.ListenAndServe(":8080", router))
}
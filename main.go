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

type Message2 struct {
  Text string `json:"text"`
}

// handler function to respond to the root URL
func rootHandler(w http.ResponseWriter, r *http.Request) {
    msg := Message{Text: "Hello, world, UPDATED!!!!!"}
    json.NewEncoder(w).Encode(msg)
}


func testHandler(w http.ResponseWriter, r *http.Request) {
  msg := Message{Text: "Hello World, testing method"}
  json.NewEncoder(w).Encode(msg)
}

func test2Handler(w http.ResponseWriter, r *http.Request) {
  msg := Message{Text: "MEOWWWW"}
  json.NewEncoder(w).Encode(msg)
}


func main() {
    // Create a new router
    router := mux.NewRouter()

    // Register the handler function for the root URL
    router.HandleFunc("/", rootHandler).Methods("GET")
    router.HandleFunc("/test", testHandler).Methods("GET")
    router.HandleFunc("/MEOW", test2Handler).Methods("GET")

		port := 8001
    machineName := "high-fructose-corn-syrup"

    // Start the server on localhost port 8080 and log any errors
		fmt.Println("Hello, world!")
		// format string
    log.Println(fmt.Sprintf("Server started on localhost:%v", port))
		log.Fatal(http.ListenAndServe(fmt.Sprintf("%s:%v", machineName, port), router))
}

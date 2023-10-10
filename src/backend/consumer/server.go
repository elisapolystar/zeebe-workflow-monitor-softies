package main

import (
	"fmt"
	"net/http"
)

func rootHandler(response http.ResponseWriter, request *http.Request) {
	fmt.Fprintf(response, "Hello there!")
}

func main() {

	fmt.Println("Backend started!")
	go Consume()
	fmt.Println("goroutine in action")
	//Call function that handles requests arriving at root
	http.HandleFunc("/", rootHandler)
	//Start server and listen port 8000
	http.ListenAndServe(":8001", nil)
}

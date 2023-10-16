package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/websocket"
)

var upgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
}

func rootHandler(response http.ResponseWriter, request *http.Request) {
	fmt.Fprintf(response, "Hello there!")
}

func reader(conn *websocket.Conn) {
	for {
		messageType, p, err := conn.ReadMessage()
		if err != nil {
			log.Println(err)
			return
		}

		log.Println(string(p))

		if err := conn.WriteMessage(messageType, p); err != nil {
			log.Println(err)
			return
		}
	}
}

func wsEndpoint(w http.ResponseWriter, r *http.Request) {
	upgrader.CheckOrigin = func(r *http.Request) bool { return true }

	ws, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Println(err)
	}

	log.Println("Client Succesfully Connected...")
	reader(ws)
}

func setupRoutes() {
	//Call function that handles requests arriving at root
	http.HandleFunc("/", rootHandler)
	http.HandleFunc("/ws", wsEndpoint)

}

func main() {

	fmt.Println("Backend started!")
	go Consume()
	setupRoutes()
	//Start server and listen port 8000
	http.ListenAndServe(":8001", nil)
}

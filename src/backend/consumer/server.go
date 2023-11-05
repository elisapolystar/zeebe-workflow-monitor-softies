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
	CheckOrigin:     func(r *http.Request) bool { return true },
}

var messageChannel chan string

func rootHandler(response http.ResponseWriter, request *http.Request) {
	fmt.Fprintf(response, "Hello there!")
}

func writer(conn *websocket.Conn) {
	for {

		message, ok := <-messageChannel

		if !ok {
			fmt.Println("Channel is closed.")
			break
		}

		err := conn.WriteMessage(websocket.TextMessage, []byte(message))
		if err != nil {
			return
		}

	}
}

func wsEndpoint(w http.ResponseWriter, r *http.Request) {

	ws, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Println(err)
	}

	log.Println("Client Succesfully Connected...")
	writer(ws)

}

//saves only process as test
func storeConsumerData(process Process) {
	SaveData(process)
}

func main() {

	fmt.Println("Backend started!")

	messageChannel = make(chan string)
	go Consume(messageChannel)
	go TestDatabase()
	http.HandleFunc("/", rootHandler)
	http.HandleFunc("/ws", wsEndpoint)
	http.HandleFunc("/test", testDatabase)

	//Start server and listen port 8000
	http.ListenAndServe(":8001", nil)
}

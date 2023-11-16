package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/websocket"
)

type topicMessagePair struct {
	Topic   string
	Message []byte
}

var upgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
	CheckOrigin:     func(r *http.Request) bool { return true },
}

var messageChannel chan topicMessagePair

func rootHandler(response http.ResponseWriter, request *http.Request) {
	fmt.Fprintf(response, "Hello there!")
}

func reader(conn *websocket.Conn) {

	for {
		messageType, p, err := conn.ReadMessage()
		if err != nil {
			log.Println("Error reading message from websocket: ", err)
			return
		}
		log.Println(string(p))

		//maybe use conn.ReadJSON()
		frontMessage, err3 := parseCommunicationItem(p)
		if err3 != nil {
			log.Println("Error parsing the message from frontend(!): ", err3)
		}

		fmt.Println()
		fmt.Println("Front message process value: ", frontMessage.Process)
		fmt.Println()
		fmt.Println("The message from front: ", string(p))
		fmt.Println()

		if frontMessage.Process == "" {

			fmt.Println()
			fmt.Println("FRONTEND IS WANTING ALL PROCESSES FRONTEND IS WANTING ALL PROCESSES FRONTEND IS WANTING ALL PROCESSES FRONTEND IS WANTING ALL PROCESSES ")
			fmt.Println()
			allProcesses := RetrieveProcesses()
			if len(string(allProcesses)) == 0 {
				fmt.Println("No processes to return from database")
				continue
			}
			fmt.Println(string(allProcesses))

			processesData := WebsocketMessage{
				Type: "process",
				Data: string(allProcesses),
			}

			processesDataJson, err := json.Marshal(processesData)
			if err != nil {
				fmt.Println("Error JSON Unmarshalling in the websocket comm section")
				fmt.Println(err.Error())
			}

			err2 := conn.WriteMessage(messageType, processesDataJson)
			if err2 != nil {
				fmt.Println("Error sending message to frontend: ", err2)
				return
			}

		} else {
			//processItem = hae_prosessi(frontMessage.Process)
			//conn.WriteMessage(processItem)
			fmt.Println("Kurwa")
		}

		/*fmt.Println()
		fmt.Println("Echo trying:D")
		fmt.Println()
		err2 := conn.WriteMessage(messageType, p)
		if err2 != nil {
			log.Println("Error sending message to frontend: ", err2)
			return
		}*/
	}
}

func wsEndpoint(w http.ResponseWriter, r *http.Request) {
	ws, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Println(err)
	}
	log.Println("Client Succesfully Connected...")

	reader(ws)
}

// Listen for the messages from the consumer and parse the messages into structs
func listenTmChannel() {

	for {

		tmPair, ok := <-messageChannel
		if !ok {
			fmt.Println("Channel is closed")
			break
		}

		fmt.Println()
		fmt.Println()
		fmt.Printf("Received message from topic %s: %s\n", tmPair.Topic, string(tmPair.Message))
		fmt.Println()
		fmt.Println("------------------------------------------------------------")
		fmt.Println()

		// Make a struct of a process JSON
		if tmPair.Topic == "zeebe-process" {

			process, err := parseProcessJson(tmPair.Message)
			if err != nil {
				fmt.Println("Error parsing the json: ", err)
			}

			SaveData(*process)
			processes := RetrieveProcesses()
			fmt.Println(string(processes))

			// Talenna_tietokantaan(process)

			fmt.Println()
			fmt.Println("++++++++++++++++++++++++++++++++++++++++++++++++++")
			fmt.Println()
			fmt.Println("Process key of the process item: ", process.Key)
			fmt.Println("bpmnProcessId: ", process.Value.BpmnProcessId)
			fmt.Println("Version: ", process.Value.Version)
			fmt.Println()
			fmt.Println("++++++++++++++++++++++++++++++++++++++++++++++++++")
			fmt.Println()

			// Structi muutetaan fronttiin lähetettäväksi JSONiksi
			jsonString, err2 := structToJson(&process)
			if err2 != nil {
				fmt.Println("Error turning struct to json: ", err2)
			}

			fmt.Println()
			fmt.Println("JSON STRING - JSON STRING - JSON STRING - JSON STRING - JSON STRING -")
			fmt.Println()
			fmt.Print(jsonString)
			fmt.Println()
			fmt.Println("JSON STRING - JSON STRING - JSON STRING - JSON STRING - JSON STRING -")
			fmt.Println()

		}

		// Make a struct of a process instance JSON
		if tmPair.Topic == "zeebe-process-instance" {

			processInstanceItem, err := parseProcessInstanceJson(tmPair.Message)
			if err != nil {
				fmt.Println("Error parsing the process instance json: ", err)
			}

			fmt.Println()
			fmt.Println("//////////////////////////////////////////////////")
			fmt.Println()
			fmt.Println("Process key of the process instance item: ", processInstanceItem.Key)
			fmt.Println("PartitionId: ", processInstanceItem.PartitionId)
			fmt.Println("Process definition key: ", processInstanceItem.Value.ProcessDefinitionKey)
			fmt.Println("Process instance process id: ", processInstanceItem.Value.BpmnProcessId)
			fmt.Println("Version: ", processInstanceItem.Value.Version)
			fmt.Println("Parent process instance key: ", processInstanceItem.Value.ParentProcessInstanceKey)
			fmt.Println("Parent element instance key: ", processInstanceItem.Value.ParentElementInstanceKey)
			fmt.Println()
			fmt.Println("//////////////////////////////////////////////////")
			fmt.Println()

			// Structi muutetaan fronttiin lähetettäväksi JSONiksi
			jsonString, err2 := structToJson(&processInstanceItem)
			if err2 != nil {
				fmt.Println("Error turning struct to json: ", err2)
			}

			fmt.Println()
			fmt.Println("JSON STRING - JSON STRING - JSON STRING - JSON STRING - JSON STRING -")
			fmt.Println()
			fmt.Print(jsonString)
			fmt.Println()
			fmt.Println("JSON STRING - JSON STRING - JSON STRING - JSON STRING - JSON STRING -")
			fmt.Println()

		}

		// Make a struct of a variable JSON
		if tmPair.Topic == "zeebe-variable" {

			variableItem, err := parseVariableJson(tmPair.Message)
			if err != nil {
				fmt.Println("Error parsing the variable JSON: ", err)
			}

			fmt.Println()
			fmt.Println("******************************************************")
			fmt.Println()
			fmt.Println("Partition id of variable : ", variableItem.PartitionId)
			fmt.Println("Position: ", variableItem.Position)
			fmt.Println("Name: ", variableItem.Value.Name)
			fmt.Println("Value: ", variableItem.Value.Value)
			fmt.Println("Processinstancekey: ", variableItem.Value.ProcessInstanceKey)
			fmt.Println("Scope key: ", variableItem.Value.ScopeKey)
			fmt.Println()
			fmt.Println("******************************************************")
			fmt.Println()
		}

		// Make a struct of a job JSON
		if tmPair.Topic == "zeebe-job" {

			jobItem, err := parseJobJson(tmPair.Message)
			if err != nil {
				fmt.Println("Error parsing the job JSON: ", err)
			}

			fmt.Println()
			fmt.Println("-.-.-.-.-.-.-.-.-.-.-.-.-.-.-.-.-.-.-.-.-.-.-.-.-.-.-.-")
			fmt.Println()
			fmt.Println("Key of job item: ", jobItem.Key)
			fmt.Println("Timestamp: ", jobItem.Timestamp)
			fmt.Println("Type of job: ", jobItem.Value.JobType)
			fmt.Println("Worker: ", jobItem.Value.Worker)
			fmt.Println("Process instance key: ", jobItem.Value.ProcessInstanceKey)
			fmt.Println("Element instance key: ", jobItem.Value.ElementInstanceKey)
			fmt.Println()
			fmt.Println("-.-.-.-.-.-.-.-.-.-.-.-.-.-.-.-.-.-.-.-.-.-.-.-.-.-.-.-")
			fmt.Println()

		}

		// Make a struct of a incident JSON
		if tmPair.Topic == "zeebe-incident" {

			incidentItem, err := parseIncidentJson(tmPair.Message)
			if err != nil {
				fmt.Println("Error parsing the incident JSON: ", err)
			}

			fmt.Println()
			fmt.Println("INCIDENT - INCIDENT - INCIDENT - INCIDENT - INCIDENT - ")
			fmt.Println()
			fmt.Println("Key of the incident item: ", incidentItem.Key)
			fmt.Println("Process Id: ", incidentItem.Value.BpmnProcessId)
			fmt.Println()
			fmt.Println("Error type: ", incidentItem.Value.ErrorType)
			fmt.Println("Error message: ", incidentItem.Value.ErrorMessage)
			fmt.Println()
			fmt.Println("INCIDENT - INCIDENT - INCIDENT - INCIDENT - INCIDENT - ")
			fmt.Println()

		}

		// Make a struct of a message JSON
		if tmPair.Topic == "zeebe-message" {

			messageItem, err := parseMessageJson(tmPair.Message)
			if err != nil {
				fmt.Println("Error parsing the message JSON: ", err)
			}

			fmt.Println()
			fmt.Println("MESSAGE - MESSAGE - MESSAGE - MESSAGE - MESSAGE - MESSAGE")
			fmt.Println()
			fmt.Println("Key of message: ", messageItem.Key)
			fmt.Println("Correlation key: ", messageItem.Value.CorrelationKey)
			fmt.Println("Name: ", messageItem.Value.Name)
			fmt.Println()
			fmt.Println()
			fmt.Println("MESSAGE - MESSAGE - MESSAGE - MESSAGE - MESSAGE - MESSAGE")
			fmt.Println()
		}

		// Make a struct of a timer JSON
		if tmPair.Topic == "zeebe-timer" {

			timerItem, err := parseTimerJson(tmPair.Message)
			if err != nil {
				fmt.Println("Error parsing the timer JSON: ", err)
			}

			fmt.Println("TIMER TIMER TIMER TIMER TIMER TIMER TIMER TIMER ")
			fmt.Println("Timer key: ", timerItem.Key)
			fmt.Println("Duedate: ", timerItem.Value.Duedate)
			fmt.Println("Repetitions: ", timerItem.Value.Repetitions)
			fmt.Println("Element instance key: ", timerItem.Value.ElementInstanceKey)
			fmt.Println("Process definition key: ", timerItem.Value.ProcessDefinitionKey)

			fmt.Println("TIMER TIMER TIMER TIMER TIMER TIMER TIMER TIMER ")

		}

	}
}

func setupRoutes() {
	http.HandleFunc("/", rootHandler)
	http.HandleFunc("/ws", wsEndpoint)
}

func main() {

	fmt.Println("Backend started!")
	messageChannel = make(chan topicMessagePair)
	go Consume(messageChannel)
	go listenTmChannel()
	fmt.Println("We are here")
	setupRoutes()
	//Start server and listen port 8000
	http.ListenAndServe(":8001", nil)
}

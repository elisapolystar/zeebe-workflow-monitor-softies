package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strconv"
	"strings"

	"github.com/gorilla/websocket"
)

// Contains the message from kafka and the topic the message came from
type topicMessagePair struct {
	Topic   string
	Message []byte
}

var upgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
	CheckOrigin:     func(r *http.Request) bool { return true },
}

// Channel that stores topic and message pairs received from kafka
var messageChannel chan topicMessagePair

func rootHandler(response http.ResponseWriter, request *http.Request) {
	fmt.Fprintf(response, "Hello there!")
}

// Interpret and answer messages from frontend
func reader(conn *websocket.Conn) {
	db, err := connectToDatabase()
	if err != nil {
		fmt.Println("Database connection failed")
	}
	for {
		// Read messages sent by frontend
		messageType, p, err := conn.ReadMessage()
		if err != nil {
			log.Println("Error reading message from websocket: ", err)
			return
		}
		if messageType == websocket.CloseMessage {
			fmt.Println("Websocket connection closed.")
			break
		}

		// Print the message received
		fmt.Println("The message from front: ", string(p))

		// Turn the json from frontend into a map so we can check wich fields it contains
		var messageData map[string]interface{}
		if err := json.Unmarshal([]byte(p), &messageData); err != nil {
			log.Println("Error unmarshalling frontend message JSON: ", err)
			return
		}

		// If the json from fron contains field process we do stuff for process request
		if processValue, ok := messageData[process_string]; ok {
			fmt.Println("Process: ", processValue)

			// Parse the json into struct so we can access the value of the "process" -field
			processMessage, err3 := parseProcessRequest(p)
			if err3 != nil {
				log.Println("Error parsing the message from frontend(!): ", err3)
			}

			if processMessage.Process == empty {

				// If the value of the "process"-field is empty then we want to return all the processes to the frontend
				fmt.Println("Fetching all processes")

				// Retrieve the processes from the database
				allProcesses := RetrieveProcesses(db)
				fmt.Println("(J)The processes retrieved: ", string(allProcesses))
				// Transfrom the retrieved processes to the correct json format that can be sent to the front
				processesData := WebsocketMessage{
					Type: all_processes,
					Data: string(allProcesses),
				}
				processesDataJson, err := json.Marshal(processesData)
				if err != nil {
					fmt.Println("Error JSON marshalling in the websocket comm section")
					fmt.Println(err.Error())
				}

				// Send all processes to the frontend in the correct format
				err2 := conn.WriteMessage(messageType, processesDataJson)
				if err2 != nil {
					fmt.Println("Error sending message to frontend: ", err2)
					return
				}

			} else if processMessage.Process != empty {

				// If the value of the "Process"-field is not empty then we want to retrieve only one process from the database
				fmt.Println("Fetching one specific process")
				key, err := strconv.ParseInt(processMessage.Process, 10, 64)
				if err != nil {
					log.Println("Error turning string to int: ", err)
				}

				// Fetch the process and turn it into correct format to be sent to front
				process := RetrieveProcessByID(db, key)
				processData := WebsocketMessage{
					Type: process_string,
					Data: string(process),
				}
				processDataJson, err2 := json.Marshal(processData)
				if err2 != nil {
					log.Println("Error marshalling websocketmessage struct: ", err2)
				}

				err3 := conn.WriteMessage(messageType, processDataJson)
				if err3 != nil {
					log.Println("Error sending single process message to frontend: ", err2)
					return
				}

			} else {
				fmt.Println("For some reason process message value is not empty and does not contain anything?")
			}

		} else if instanceValue, ok := messageData[instance_string]; ok {

			// If frontend asks for instances
			fmt.Println("Instance?: ", instanceValue)
			//Parse the message from frontend into a struct
			instanceMessage, err := parseInstanceRequest(p)
			if err != nil {
				log.Println("Error parsing instance request to struct: ", err)
			}

			if instanceMessage.Instance == empty {

				// Fetch all instances
				fmt.Println("Fetching all instances")
				allInstances := RetrieveInstances(db)
				instancesData := WebsocketMessage{
					Type: all_instances,
					Data: string(allInstances),
				}
				instancesDataJson, err := json.Marshal(instancesData)
				if err != nil {
					fmt.Println("Error marshalling the instancesData item to json")
					fmt.Println(err.Error())
				}

				err2 := conn.WriteMessage(messageType, instancesDataJson)
				if err2 != nil {
					fmt.Println("Error sending instances to frontend", err2)
				}

			} else if instanceMessage.Instance != empty {

				// Fetch a specific instance
				fmt.Println("Fetching a specific instance")
				key, err := strconv.ParseInt(instanceMessage.Instance, 10, 64)
				if err != nil {
					log.Println("Error turning string to int: ", err)
				}

				// Retrieve the instance item so we can get the corresponding process with the definitionkey of the instance
				instance, err2 := RetrieveInstanceByID(db, "ProcessInstanceKey", key)
				if err2 != nil {
					fmt.Println("Error getting instance by id: ", err2)
				}
				var instanceItem ProcessInst
				err3 := json.Unmarshal([]byte(instance), &instanceItem)
				if err3 != nil {
					fmt.Println("error unmarshalin instance json: ", err3)
				}

				// Fetch everything thats needed to build the instance response json
				processJson := RetrieveProcessByID(db, instanceItem.ProcessDefinitionKey)
				elements := RetrieveElementByID(db, key)
				variables := RetrieveVariableByID(db, key)
				timers := RetrieveTimerByID(db, key)
				incidents := RetrieveIncidentByID(db, key)

				// Combine the jsons
				combinedJSON, err4 := concatenateInstanceJSON(
					[]byte(processJson),
					[]byte(elements),
					[]byte(variables),
					[]byte(timers),
					[]byte(incidents))
				if err4 != nil {
					fmt.Println("Error combining the jsons: ", err4)
					return
				}

				// Add an outer layer to the json so frontend knows what we are sending them
				instanceData := WebsocketMessage{
					Type: instance_string,
					Data: string(*combinedJSON),
				}
				instanceDataJson, err5 := json.Marshal(instanceData)
				if err5 != nil {
					fmt.Println("(#asd123J) Error JSON marshalling the instanceData block")
					fmt.Println(err5.Error())
				}

				// Send the instance json to frontend
				err6 := conn.WriteMessage(messageType, instanceDataJson)
				if err6 != nil {
					fmt.Println("Error sending instances to frontend", err6)
				}
			}

		} else if incidentValue, ok := messageData[incident_string]; ok {

			// Frontend requests incidents
			fmt.Println("Incident value: ", incidentValue)
			// retrieve all incidents from the database and create a WS message with the data
			allIncidents := RetrieveIncidents(db)
			incidentsData := WebsocketMessage{
				Type: all_incidents,
				Data: string(allIncidents),
			}
			// parse the message to a json format
			incidentsDataJson, err := json.Marshal(incidentsData)
			if err != nil {
				fmt.Println("Error marshalling the incidentsData item to json")
				fmt.Println(err.Error())
			}
			// send the json data to the frontend
			fmt.Println("Incidents json we are sending to front: ", string(incidentsDataJson))
			err2 := conn.WriteMessage(messageType, incidentsDataJson)
			if err2 != nil {
				fmt.Println("Error sending instances to frontend", err2)
			}
		} else {
			fmt.Println("Unrecognized message")
		}
	}
}

// Create the websocket connection
func wsEndpoint(w http.ResponseWriter, r *http.Request) {
	ws, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Println(err)
	}
	log.Println("Client Succesfully Connected...")

	defer ws.Close()
	reader(ws)
}

// Listen for the messages from the consumer, parse the messages into structs and save the structs to the database
func listenTmChannel() {
	db, err := connectToDatabase()
	if err != nil {
		fmt.Println("Database connection failed")
	}
	for {

		// Receive messages from the consumer
		tmPair, ok := <-messageChannel
		if !ok {
			fmt.Println("Channel is closed")
			break
		}

		// Check from which topic the message is, turn the message into a struct and save it to the database
		if tmPair.Topic == zeebe {

			zeebeItem, err := parseZeebeJson(tmPair.Message)
			if err != nil {
				fmt.Println("Error parsing the process instance json: ", err)
			}

			SaveData(db, *zeebeItem)
		}

		// Make a struct of a process JSON
		if tmPair.Topic == zeebe_process {

			process, err := parseProcessJson(tmPair.Message)
			if err != nil {
				fmt.Println("Error parsing the json: ", err)
			}

			SaveData(db, *process)
			processes := RetrieveProcessByID(db, process.Key)
			fmt.Println("The process we just saved: ", string(processes))
		}

		// Make a struct of a process instance JSON
		if tmPair.Topic == zeebe_process_instance {

			if strings.Contains(string(tmPair.Message), intent_str) {

				elementItem, err := parseElementJson(tmPair.Message)
				if err != nil {
					fmt.Println("Error parsing the process instance json: ", err)
				}

				SaveData(db, *elementItem)
			} else {

				// Probably useless
				processInstanceItem, err := parseProcessInstanceJson(tmPair.Message)
				if err != nil {
					fmt.Println("Error parsing the process instance json: ", err)
				}
				fmt.Println("process instance item: ", processInstanceItem)
			}
		}

		// Make a struct of a variable JSON
		if tmPair.Topic == zeebe_variable {

			variableItem, err := parseVariableJson(tmPair.Message)
			if err != nil {
				fmt.Println("Error parsing the variable JSON: ", err)
			}
			SaveData(db, *variableItem)
		}

		// Make a struct of a job JSON
		if tmPair.Topic == zeebe_job {

			jobItem, err := parseJobJson(tmPair.Message)
			if err != nil {
				fmt.Println("Error parsing the job JSON: ", err)
			}
			SaveData(db, *jobItem)
		}

		// Make a struct of a incident JSON
		if tmPair.Topic == zeebe_incident {

			incidentItem, err := parseIncidentJson(tmPair.Message)
			if err != nil {
				fmt.Println("Error parsing the incident JSON: ", err)
			}
			SaveData(db, *incidentItem)
		}

		// Make a struct of a message JSON
		if tmPair.Topic == zeebe_message {

			messageItem, err := parseMessageJson(tmPair.Message)
			if err != nil {
				fmt.Println("Error parsing the message JSON: ", err)
			}
			SaveData(db, *messageItem)
		}

		// Make a struct of a timer JSON
		if tmPair.Topic == zeebe_timer {

			timerItem, err := parseTimerJson(tmPair.Message)
			if err != nil {
				fmt.Println("Error parsing the timer JSON: ", err)
			}
			SaveData(db, *timerItem)
		}
	}
}

func setupRoutes() {
	http.HandleFunc("/", rootHandler)
	http.HandleFunc("/ws", wsEndpoint)
}

func connectToDatabase() (*sql.DB, error) {
	//pass variables to the connection string
	DBConnection := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable", host, port, user, password, DBname)
	// Open a database connection, and check that it works
	db, err := sql.Open("postgres", DBConnection)
	if err != nil {
		return nil, err
	}
	fmt.Println("Connected to the database!")
	return db, nil
}

func main() {

	fmt.Println("Backend started!")
	// Create a channel where messages from the consumer will be put and read from
	messageChannel = make(chan topicMessagePair)
	go Consume(messageChannel)
	go listenTmChannel()
	fmt.Println("We are here")
	setupRoutes()
	//Start server and listen port 8000
	http.ListenAndServe(":8001", nil)

}

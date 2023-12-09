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

		// Print the message received for debuging reasons
		fmt.Println()
		fmt.Println("The message from front: ", string(p))
		fmt.Println()

		// Turn the json from frontend into a map so we can check wich fields it contains
		var messageData map[string]interface{}
		if err := json.Unmarshal([]byte(p), &messageData); err != nil {
			log.Println("Error unmarshalling frontend message JSON: ", err)
			return
		}

		// If the json from fron contains field process we do stuff for process request
		if processValue, ok := messageData["process"]; ok {
			fmt.Println("Process: ", processValue)

			// Parse the json into struct so we can access the value of the "process" -field
			processMessage, err3 := parseProcessRequest(p)
			if err3 != nil {
				log.Println("Error parsing the message from frontend(!): ", err3)
			}
			fmt.Println("The value of the process field in the json front sent: ", processMessage.Process)

			// If the value of the "process"-field is empty then we want to return all the processes to the frontend
			if processMessage.Process == "" {

				fmt.Println()
				fmt.Println("FRONTEND IS WANTING ALL PROCESSES ")
				fmt.Println()

				// Retrieve the processes from the database
				allProcesses := RetrieveProcesses(db)
				fmt.Println("(J)The processes retrieved: ", string(allProcesses))

				// Transfrom the retrieved processes to the correct json format that can be sent to the front
				processesData := WebsocketMessage{
					Type: "all-processes",
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

				// Then if the value of the "process"-field is not empty then we want to retrieve only one process from the database
			} else if processMessage.Process != "" {

				fmt.Println()
				fmt.Println("FRONTEND IS WANTING ONLY ONE PROCESS ")
				fmt.Println()

				key, err := strconv.ParseInt(processMessage.Process, 10, 64)
				if err != nil {
					log.Println("Error turning string to int: ", err)
				}
				fmt.Println(key)

				process := RetrieveProcessByID(db, key)
				fmt.Println("(J) The retrieved process: ", string(process))

				processData := WebsocketMessage{
					Type: "process",
					Data: string(process),
				}
				processDataJson, err2 := json.Marshal(processData)
				if err2 != nil {
					log.Println("Error marshalling websocketmessage struct: ", err2)
				}

				fmt.Println("The JSON we are sending to front: ", processDataJson)
				err3 := conn.WriteMessage(messageType, processDataJson)
				if err3 != nil {
					log.Println("Error sending single process message to frontend: ", err2)
					return
				}

			} else {
				fmt.Println("For some reason process message value is not empty and does not contain anything?")
			}

			// If frontend asks for instances
		} else if instanceValue, ok := messageData["instance"]; ok {
			fmt.Println("Instance?: ", instanceValue)

			//Parse the message into a struct
			instanceMessage, err := parseInstanceRequest(p)
			if err != nil {
				log.Println("Error parsing instance request to struct: ", err)
			}

			if instanceMessage.Instance == "" {
				// get all instances

				allInstances := RetrieveInstances(db)

				fmt.Println()
				fmt.Println("-.-.-.-.-.-.-.-.-.-.-.-.-.-.-.-.-.-.-.-.-.-.-.-.-.-.-.-.")
				fmt.Println()
				fmt.Println("(J) All the retrieved instances: ", string(allInstances))
				fmt.Println()
				fmt.Println("-.-.-.-.-.-.-.-.-.-.-.-.-.-.-.-.-.-.-.-.-.-.-.-.-.-.-.-.")
				fmt.Println()

				instancesData := WebsocketMessage{
					Type: "all-instances",
					Data: string(allInstances),
				}
				instancesDataJson, err := json.Marshal(instancesData)
				if err != nil {
					fmt.Println("Error marshalling the instancesData item to json")
					fmt.Println(err.Error())
				}

				fmt.Println("Instances json we are sending to front: ", string(instancesDataJson))
				err2 := conn.WriteMessage(messageType, instancesDataJson)
				if err2 != nil {
					fmt.Println("Error sending instances to frontend", err2)
				}

				// Get a specific instance
			} else if instanceMessage.Instance != "" {

				// Muunnetaan parametrina saatu avain intiksi
				key, err := strconv.ParseInt(instanceMessage.Instance, 10, 64)
				if err != nil {
					log.Println("Error turning string to int: ", err)
				}

				// Haetaan yksi instanssi jotta voidaan hakea prosessi instanssin processDefinition avaimella
				instance, err2 := RetrieveInstanceByID(db, "ProcessInstanceKey", key)
				if err2 != nil {
					fmt.Println("Error getting instance by id: ", err2)
				}

				fmt.Println()
				fmt.Println("hhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhh")
				fmt.Println()
				fmt.Println("The retrieved INSTANCE: ", string(instance))
				fmt.Println()
				fmt.Println("hhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhh")
				fmt.Println()

				// Luodaan haetusta instanssi-jsonista strukti
				var instanceItem ProcessInst
				err3 := json.Unmarshal([]byte(instance), &instanceItem)
				if err3 != nil {
					fmt.Println("error unmarshalin instance json: ", err3)
				}

				fmt.Println()
				fmt.Println("Processinstance key: ", instanceItem.ProcessInstanceKey)
				fmt.Println()

				// Haetaan prosessi
				processJson := RetrieveProcessByID(db, instanceItem.ProcessDefinitionKey)
				//fmt.Println("The retrieved process for the instance item: ", string(processJson))

				elements := RetrieveElementByID(db, key)

				fmt.Println()
				fmt.Println("E      E     E      E      E      E   E      E    E      E ")
				fmt.Println()
				fmt.Println("The retrieved ELEMENTS: ", string(elements))
				fmt.Println()
				fmt.Println("E      E     E      E      E      E   E      E    E      E ")
				fmt.Println()

				variables := RetrieveVariableByID(db, key)

				fmt.Println()
				fmt.Println("vvvvvvvvvvvvvvvvvvvvvvvvvvvvvvvvvvvvvvvvvvvvvvvvvvvvvvvvv")
				fmt.Println()
				fmt.Println("The retrieved VARIABLES: ", string(variables))
				fmt.Println()
				fmt.Println("vvvvvvvvvvvvvvvvvvvvvvvvvvvvvvvvvvvvvvvvvvvvvvvvvvvvvvvvv")
				fmt.Println()

				timers := RetrieveTimerByID(db, key)

				fmt.Println()
				fmt.Println("tttttttttttttttttttttttttttttttttttttttttttttttttttttttttt")
				fmt.Println()
				fmt.Println("The retrieved TIMERS: ", string(timers))
				fmt.Println()
				fmt.Println("tttttttttttttttttttttttttttttttttttttttttttttttttttttttttt")
				fmt.Println()

				incidents := RetrieveIncidentByID(db, key)

				fmt.Println()
				fmt.Println("iiiiiiiiiiiiiiiiiiiiiiiiiiiiiiiiiiiiiiiiiiiiiiiiiiiiiiiiiii")
				fmt.Println()
				fmt.Println("The retrieved INCIDENTS: ", string(incidents))
				fmt.Println()
				fmt.Println("iiiiiiiiiiiiiiiiiiiiiiiiiiiiiiiiiiiiiiiiiiiiiiiiiiiiiiiiiii")
				fmt.Println()

				// concatenateJSON funktion ois tarkotus yhistää noi haetut jsonit
				// concatenateJSON funktio löytyy jsonHandler.gon pohjalta
				combinedJSON, err4 := concatenateJSON([]byte(processJson),
					[]byte(elements),
					[]byte(variables),
					[]byte(timers),
					[]byte(incidents))

				if err4 != nil {
					fmt.Println("Error combining the jsons: ", err4)
					return
				}
				// Tällä printillä näkee onko halutun vastauksen sisemmällä jsonilla haluttu rakenne
				// Googlesta jsonformatterilla on helppo tarkastella onko oikeen näkönen json
				fmt.Println()
				fmt.Println("!     !     !     !     !     !     !     !     !")
				fmt.Println()
				fmt.Println("The combined json: ", string(*combinedJSON))
				fmt.Println()
				fmt.Println("!     !     !     !     !     !     !     !     !")
				fmt.Println()

				/* Siihen lopulliseen jsoniin saa lisättyä semmosen halutun
				ulkokuoren tekemällä näistä jutuista WebsocketMessage structin
				ja muuttamalla sen jsoniks
				*/
				instanceData := WebsocketMessage{
					Type: "instance",
					Data: string(*combinedJSON),
				}
				instanceDataJson, err5 := json.Marshal(instanceData)
				if err5 != nil {
					fmt.Println("(#asd123J) Error JSON marshalling the instanceData block")
					fmt.Println(err5.Error())
				}

				// tätä json printtiä on vaikee tarkastella ku tähän tulee semmosia kenoviivoja niin toi ylempi on parempi
				fmt.Println()
				fmt.Println("!     !     !     !     !     !     !     !     !")
				fmt.Println()
				fmt.Println("The final json to be sent: ", string(instanceDataJson))
				fmt.Println()
				fmt.Println("!     !     !     !     !     !     !     !     !")
				fmt.Println()

				// Send the instance json to frontend
				err6 := conn.WriteMessage(messageType, instanceDataJson)
				if err6 != nil {
					fmt.Println("Error sending instances to frontend", err6)
				}
			}
		// frontend requests incidents	
		} else if incidentValue, ok := messageData["incident"]; ok {
			// retrieve all incidents from the database and create a WS message with the data
			allIncidents := RetrieveIncidents(db)
			incidentsData := WebsocketMessage{
				Type: "all-incidents",
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

func wsEndpoint(w http.ResponseWriter, r *http.Request) {
	ws, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Println(err)
	}
	log.Println("Client Succesfully Connected...")

	defer ws.Close()
	reader(ws)
}

// Listen for the messages from the consumer and parse the messages into structs
func listenTmChannel() {
	db, err := connectToDatabase()
	if err != nil {
		fmt.Println("Database connection failed")
	}
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

		if tmPair.Topic == "zeebe" {

			zeebeItem, err := parseZeebeJson(tmPair.Message)
			if err != nil {
				fmt.Println("Error parsing the process instance json: ", err)
			}

			SaveData(db, *zeebeItem)

			// Structi muutetaan fronttiin lähetettäväksi JSONiksi
			jsonString, err2 := structToJson(&zeebeItem)
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

		// Make a struct of a process JSON
		if tmPair.Topic == "zeebe-process" {

			process, err := parseProcessJson(tmPair.Message)
			if err != nil {
				fmt.Println("Error parsing the json: ", err)
			}

			SaveData(db, *process)
			processes := RetrieveProcessByID(db, process.Key)
			fmt.Println("The process we just saved: ", string(processes))

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

			if strings.Contains(string(tmPair.Message), intent_str) {

				elementItem, err := parseElementJson(tmPair.Message)
				if err != nil {
					fmt.Println("Error parsing the process instance json: ", err)
				}

				SaveData(db, *elementItem)

				// Structi muutetaan fronttiin lähetettäväksi JSONiksi
				jsonString, err2 := structToJson(&elementItem)
				if err2 != nil {
					fmt.Println("Error turning struct to json: ", err2)
				}

				fmt.Print(jsonString)
			} else {

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
			fmt.Println("Timestamp: ", incidentItem.Timestamp)
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
	messageChannel = make(chan topicMessagePair)
	go Consume(messageChannel)
	go listenTmChannel()
	fmt.Println("We are here")
	setupRoutes()
	//Start server and listen port 8000
	http.ListenAndServe(":8001", nil)

}

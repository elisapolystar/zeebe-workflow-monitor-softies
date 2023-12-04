package main


import (
	"fmt"
	"strconv"
	"errors"
	"database/sql"
	"encoding/json"
	_ "github.com/lib/pq"
)

const (
	host = "postgres"
	port = 5432
	user = "postgres"
	password = "password"
	DBname = "workflow"
	ProcessesQuery = "SELECT p.Key, p.BpmnProcessId, p.Version, p.Timestamp, COUNT(i.ProcessDefinitionKey) FROM process p LEFT JOIN process_instance i ON p.Key = i.ProcessDefinitionKey GROUP BY p.Key ORDER BY Timestamp DESC"
	ProcessByIDQuery = "SELECT * FROM process WHERE key = %s"
	InstanceByIDQuery = "SELECT * FROM process_instance WHERE %s = %s ORDER BY Timestamp DESC"
	InstancesQuery = "SELECT * FROM process_instance ORDER BY Timestamp DESC"
	VariableByIDQuery = "SELECT * FROM variable WHERE ProcessInstanceKey = %s"
	IncidentByIDQuery = "SELECT * FROM incident WHERE ProcessInstanceKey = %s"
	IncidentsQuery = "SELECT * FROM incident ORDER BY Timestamp DESC"
	MessageByIDQuery = "SELECT * FROM message WHERE key = %s"
	TimerByIDQuery = "SELECT * FROM timer WHERE ProcessInstanceKey = %s"
	ElementByIDQuery = "SELECT * FROM element WHERE ProcessInstanceKey = %s"
)

type SimpleProcess struct {
	Key				int64	`json:"key"`
	BpmnProcessId 	string 	`json:"bpmnProcessId"`
	Version       	int64  	`json:"version"`
	Timestamp 		int64	`json:"timestamp"`
	Instances		int64	`json:"instances"`
}
type FullProcess struct {
	Key				int64	`json:"key"`
	BpmnProcessId 	string 	`json:"bpmnProcessId"`
	Version       	int64  	`json:"version"`
	Resource 		string 	`json:"resource"`
	Timestamp 		int64	`json:"timestamp"`
}
type ProcessInst struct {
	ProcessInstanceKey	int64	`json:"ProcessInstanceKey"`
	PartitionID		int64	`json:"PartitionID"`
	ProcessDefinitionKey	int64	`json:"ProcessDefinitionKey"`
	BpmnProcessId	string	`json:"BpmnProcessId"`
	Version		int64	`json:"Version"`
	Timestamp	int64	`json:"Timestamp"`
	Active		bool	`json:"Active"`
}
type variable struct {
	PartitionId		int64	`json:"PartitionID"`
	Position		int64	`json:"Position"`
	Name			string	`json:"Name"`
	Value			string	`json:"Value"`
	ProcessInstanceKey	int64	`json:"ProcessInstanceKey"`
	ScopeKey	int64	`json:"ScopeKey"`
}
type incident struct {
	Key				int64	`json:"Key"`
	BpmnProcessId	string	`json:"BpmnProcessId"`
	ProcessInstanceKey	int64	`json:"ProcessInstanceKey"`
	ElementInstanceKey	int64	`json:"ElementInstanceKey"`
	JobKey			int64	`json:"JobKey"`
	ErrorType		string	`json:"ErrorType"`
	ErrorMessage	string	`json:"ErrorMessage"`
	Timestamp		int64	`json:"Timestamp"`
}
type message struct {
	Key				int64	`json:"Key"`
	Name			string	`json:"Name"`
	CorrelationKey				string	`json:"Key"`
	MessageId		string	`json:"MessageId"`
	Timestamp		int64	`json:"Timestamp"`
}
type timer struct {
	Key				int64	`json:"Key"`
	Timestamp		int64	`json:"Timestamp"`
	ProcessDefinitionKey	int64	`json:"ProcessDefinitonKey"`
	ProcessInstanceKey	int64	`json:"ProcessInstanceKey"`
	ElementInstanceKey	int64	`json:"ElementInstanceKey"`
	TargetElementId		string	`json:"TargetElementId"`
	Duedate				int64	`json:"Duedate"`
	Repetitions			int64	`json:"Repetitions"`
}
type element struct {
	Key				int64	`json:"Key"`
	ProcessInstanceKey	int64	`json:"ProcessInstanceKey"`
	ProcessDefinitionKey	int64	`json:"ProcessDefinitionKey"`
	BpmnProcessId	string	`json:"BpmnProcessId"`
	ElementId		string	`json:"ElementId"`
	BpmnElementType	string	`json:"BpmnElementType"`
	Intent			string	`json:"Intent"`
}

func SaveData(db *sql.DB, entity interface{}) {
	//check what type of entity we are saving
	switch d := entity.(type) {
	//save a process entity
	case Process:
		process := d;
		fmt.Println("saving process")
		insertProcess := `INSERT INTO process (Key, BpmnProcessId, Version, Resource, Timestamp) VALUES ($1, $2, $3, $4, $5)`
		//execute the insertion command with entity as parameters
		_, err := db.Exec(insertProcess, process.Key, process.Value.BpmnProcessId, process.Value.Version, process.Value.Resource, process.Timestamp)
		if err != nil {
			fmt.Println("Data insertion into database failed")
			fmt.Println(err)
		} else {
			fmt.Println("saved process to database!")
		}
	//save a zeebe entity
	case Zeebe:
		fmt.Println("saving zeebe")
		zeebe := d;
		insertZeebe := `INSERT INTO process_instance (ProcessInstanceKey, PartitionID, ProcessDefinitionKey, BpmnProcessId, Version, Timestamp, Active) VALUES ($1, $2, $3, $4, $5, $6, $7)`
		_, err := db.Exec(insertZeebe, zeebe.Value.ProcessInstanceKey, zeebe.PartitionId, zeebe.Value.ProcessDefinitionKey, zeebe.Value.BpmnProcessId, zeebe.Value.Version, zeebe.Timestamp, zeebe.Active)
		if err != nil {
			fmt.Println("Failed to save instance to the database")
			fmt.Println(err)
		} else {
			fmt.Println("saved instance to database!")
		}	
	// save a timer entity
	case Timer:
		fmt.Println("saving timer")
		timer := d;
		insertTimer := `INSERT INTO timer (Key, Timestamp, ProcessDefinitionKey, ProcessInstanceKey, ElementInstanceKey, TargetElementId, Duedate, Repetitions) VALUES ($1, $2, $3, $4, $5, $6, $7, $8)`
		_, err := db.Exec(insertTimer, timer.Key, timer.Timestamp, timer.Value.ProcessDefinitionKey, timer.Value.ProcessInstanceKey, timer.Value.ElementInstanceKey, timer.Value.TargetElementId, timer.Value.Duedate, timer.Value.Repetitions)
		if err != nil {
			fmt.Println("Failed to save timer to the database")
			fmt.Println(err)
		} else {
			fmt.Println("saved timer to the database!")
		}
	// save a job entity
	case Job:
		fmt.Println("saving job")
		job := d;
		insertJob := `INSERT INTO job (Key, Timestamp, ProcessInstanceKey, ElementInstanceKey, JobType, Worker, Retries) VALUES ($1, $2, $3, $4, $5, $6, $7)`
		_, err := db.Exec(insertJob, job.Key, job.Timestamp, job.Value.ProcessInstanceKey, job.Value.ElementInstanceKey, job.Value.JobType, job.Value.Worker, job.Value.Retries)
		if err != nil {
			fmt.Println("Failed to save job to the database")
			fmt.Println(err)
		} else {
			fmt.Println("saved job to the database!")
		}
	// save a message entity
	case Message:
		fmt.Println("saving message")
		msg := d;
		insertMsg := `INSERT INTO message (Key, Name, CorrelationKey, MessageId, Timestamp) VALUES ($1, $2, $3, $4, $5)`
		_, err := db.Exec(insertMsg, msg.Key, msg.Value.Name, msg.Value.CorrelationKey, msg.Value.MessageId, msg.Timestamp)
		if err != nil {
			fmt.Println("Failed to save message to the database")
			fmt.Println(err)
		} else {
			fmt.Println("saved message to the database!")
		}
	// save an incident entity
	case Incident:
		fmt.Println("saving incident")
		incident := d;
		insertIncident := `INSERT INTO incident (Key, BpmnProcessId, ProcessInstanceKey, ElementInstanceKey, JobKey, ErrorType, ErrorMessage, Timestamp) VALUES ($1, $2, $3, $4, $5, $6, $7, $8)`
		_, err := db.Exec(insertIncident, incident.Key, incident.Value.BpmnProcessId, incident.Value.ProcessInstanceKey, incident.Value.ElementInstanceKey, incident.Value.JobKey, incident.Value.ErrorType, incident.Value.ErrorMessage, incident.Timestamp)
		if err != nil {
			fmt.Println("Failed to save incident to the database")
			fmt.Println(err)
		} else {
			fmt.Println("saved incident to the database!")
		}
	// save a variable entity to the database
	case Variable:
		fmt.Println("saving variable")
		variable := d;
		insertVariable := `INSERT INTO variable (PartitionID, Position, Name, Value, ProcessInstanceKey, ScopeKey) VALUES ($1, $2, $3, $4, $5, $6)`
		_, err := db.Exec(insertVariable, variable.PartitionId, variable.Position, variable.Value.Name, variable.Value.Value, variable.Value.ProcessInstanceKey, variable.Value.ScopeKey)	
		if err != nil {
			fmt.Println("Failed to save variable to the database")
			fmt.Println(err)
		} else {
			fmt.Println("saved variable to the database!")
		}
	
	// save an element entity to the database
	case Element:
		fmt.Println("saving element")
		element := d;
		// check if the element already exists. If yes, only update then intent value
		CheckIfExists := `SELECT * FROM element WHERE Key = $1`
		var storeIntent string = element.Intent
		err := db.QueryRow(CheckIfExists, element.Key).Scan(&element.Key, &element.Value.ProcessInstanceKey, &element.Value.ProcessDefinitionKey, &element.Value.BpmnProcessId, &element.Value.ElementId, &element.Value.BpmnElementType, &element.Intent)
		if err == sql.ErrNoRows{
			insertElement := `INSERT INTO ELEMENT (Key, ProcessInstanceKey, ProcessDefinitionKey, BpmnProcessId, ElementId, BpmnElementType, Intent) VALUES ($1, $2, $3, $4, $5, $6, $7)`	
			_, err = db.Exec(insertElement, element.Key, element.Value.ProcessInstanceKey, element.Value.ProcessDefinitionKey, element.Value.BpmnProcessId, element.Value.ElementId, element.Value.BpmnElementType, element.Intent)
			if err != nil {
				fmt.Println("Failed to save element to the database")
				fmt.Println(err)
			} else {
				fmt.Println("saved element to the database!")
			}
		} else if err != nil {
			fmt.Println(err)
		} else {
			UpdateIntent := `UPDATE element SET Intent = $1 WHERE Key = $2`
			_, err = db.Exec(UpdateIntent, storeIntent, element.Key)
			if (err != nil){
				fmt.Println("Failed to update element: ", err)
			}
			fmt.Println("Successfully updated element")
		}

	default:
        fmt.Println("Unsupported entity")
	}

}

// retrieves every process, along with the number of instances for each process.
func RetrieveProcesses(db *sql.DB) string {
	fmt.Println("retrieving processes from the database")
	fmt.Println("processes retrieved succesfully")
	rows, err := db.Query(ProcessesQuery)
	if err != nil {
		fmt.Println("Query failed!")
	}
	defer rows.Close()

	//array for the processes
	var processes []SimpleProcess

	for rows.Next(){
		var p SimpleProcess
		err := rows.Scan(&p.Key, &p.BpmnProcessId, &p.Version, &p.Timestamp, &p.Instances)
		if err != nil {
			fmt.Println("Failed to scan rows")
		}
		processes = append(processes, p)
	}
	jsonData, err := json.Marshal(processes)
	if err != nil {
		fmt.Println("Failed to transform to json")
	}
	return string(jsonData)
}
// retrieves a process with the given Key (Process definition key)
func RetrieveProcessByID(db *sql.DB, key int64) string {
	fmt.Println("Retrieving the Process...")
	// Perform the query
	var strkey string = strconv.FormatInt(key, 10)
	db_query := fmt.Sprintf(ProcessByIDQuery, strkey)
	rows, err := db.Query(db_query)
	if err != nil {
		fmt.Println("Query failed")
	}
	defer rows.Close()

	counter := 0
	fmt.Println("Process retrieved successfully!")
	fmt.Println("Converting data to JSON...")
	// Try to convert data to a JSON format
	var p FullProcess
	for rows.Next(){
		counter++
		err := rows.Scan(&p.Key, &p.BpmnProcessId, &p.Version, &p.Resource, &p.Timestamp)
		if err != nil {
			fmt.Println("Failed to scan row")
		}
	}
	//check if rows exist and if not return an error JSON.
	fmt.Println("Checking if query return rows")
	if counter == 0 {
		fmt.Println("results empty. Returning error.")
		message := GenerateErrorMessage("Process not found")	
		return message
	//convert to string otherwise
	} else {
		json, err := json.Marshal(p)
		if err != nil {
			fmt.Println("Failed to convert data to JSON")
		}
		return string(json)	
	}
  
}
// Retrieves an instance from the database. 
// column = The name of a specific column, key = the value we want to find from the column
// accepts "ProcessDefinitionKey" or "ProcessInstanceKey" as the column parameter
func RetrieveInstanceByID(db *sql.DB, column string, key int64) (string, error) {
	if (column == "ProcessDefinitionKey") || (column == "ProcessInstanceKey"){
		fmt.Println("retrieving an instance")
		db, err := connectToDatabase()
		if err != nil {
			fmt.Println("Error opening database connection")
		}
		// perform the query
		var strkey string = strconv.FormatInt(key, 10)
		db_query := fmt.Sprintf(InstanceByIDQuery, column, strkey)
		rows, err := db.Query(db_query)
		if err != nil {
			fmt.Println("Query failed!")
		}
		defer rows.Close()
		// create the JSON and return it
		var p ProcessInst
		for rows.Next(){
			err := rows.Scan(&p.ProcessInstanceKey, &p.PartitionID, &p.ProcessDefinitionKey, &p.BpmnProcessId, &p.Version, &p.Timestamp, &p.Active)
			if err != nil {
				fmt.Println("Failed to scan row:", err)
			}
		}
		json, err := json.Marshal(p)
		if err != nil {
			fmt.Println("Failed to conver data to JSON")
		}
		return string(json), nil
		} else {
		return "", errors.New("invalid column")
	}
}
// retrieves all process instances from the database, and returns them ordered from newest to oldest.
func RetrieveInstances(db *sql.DB) string {
	fmt.Println("retrieving all instances from the database")
	fmt.Println("processes retrieved succesfully")
	rows, err := db.Query(InstancesQuery)
	if err != nil {
		fmt.Println("Query failed!")
	}
	defer rows.Close()

	//array for the process instances
	var instances []ProcessInst

	for rows.Next(){
		var i ProcessInst
		err := rows.Scan(&i.ProcessInstanceKey, &i.PartitionID, &i.ProcessDefinitionKey, &i.BpmnProcessId, &i.Version, &i.Timestamp, &i.Active)
		if err != nil {
			fmt.Println("Failed to scan rows")
		}
		instances = append(instances, i)
	}
	jsonData, err := json.Marshal(instances)
	if err != nil {
		fmt.Println("Failed to transform to json")
	}
	return string(jsonData)

}

// retrieves a variable with the ProcessInstanceKey specified in the parameter
func RetrieveVariableByID(db *sql.DB, key int64) (string){
	fmt.Println("Retrieving the variable...")
	// Perform the query
	var strkey string = strconv.FormatInt(key, 10)
	db_query := fmt.Sprintf(VariableByIDQuery, strkey)
	rows, err := db.Query(db_query)
	if err != nil {
		fmt.Println("Query failed!")
	}
	defer rows.Close()
	fmt.Println("Variable retrieved successfully!")
	fmt.Println("Converting data to JSON...")
	// Convert data to a JSON format
	var v variable
	for rows.Next(){
		err := rows.Scan(&v.PartitionId, &v.Position, &v.Name, &v.Value, &v.ProcessInstanceKey, &v.ScopeKey)
		if err != nil {
			fmt.Println("Failed to scan row")
		}		
	}
	json, err := json.Marshal(v)
	if err != nil {
		fmt.Println("Failed to convert data to JSON")
	}
	return string(json)	
}
// Retrieves an incident with the given ProcessInstanceKey
func RetrieveIncidentByID(db *sql.DB, key int64) string{
	fmt.Println("Retrieving the incident...")
	// Perform the query
	var strkey string = strconv.FormatInt(key, 10)
	db_query := fmt.Sprintf(IncidentByIDQuery, strkey)
	rows, err := db.Query(db_query)
	if err != nil {
        fmt.Println("Query failed")
    }
	defer rows.Close()
	fmt.Println("incident retrieved successfully!")
	fmt.Println("Converting data to JSON...")
	// Convert data to a JSON format
	var i incident
	for rows.Next(){
		err := rows.Scan(&i.Key, &i.BpmnProcessId, &i.ProcessInstanceKey, &i.ElementInstanceKey, &i.JobKey, &i.ErrorType, &i.ErrorMessage, &i.Timestamp)
		if err != nil {
			fmt.Println("Failed to scan row")
		}		
	}
	json, err := json.Marshal(i)
	if err != nil {
		fmt.Println("Failed to convert data to JSON")
	}
	return string(json)		
}
// Returns all incidents in the database, sorted from newest to oldest.
func RetrieveIncidents(db *sql.DB) string {
	fmt.Println("retrieving all incidents from the database")
	fmt.Println("incidents retrieved succesfully")
	rows, err := db.Query(IncidentsQuery)
	if err != nil {
        fmt.Println("Query failed")
    }
	defer rows.Close()

	//array for the incidents
	var incidents []incident

	for rows.Next(){
		var i incident
		err := rows.Scan(&i.Key, &i.BpmnProcessId, &i.ProcessInstanceKey, &i.ElementInstanceKey, &i.JobKey, &i.ErrorType, &i.ErrorMessage, &i.Timestamp)
		if err != nil {
			fmt.Println("Failed to scan rows")
		}
		incidents = append(incidents, i)
	}
	jsonData, err := json.Marshal(incidents)
	if err != nil {
		fmt.Println("Failed to transform to json")
	}
	return string(jsonData)	
}
// retrieves a message with a given key
func RetrieveMessageByID(db *sql.DB, key int64) string {
	fmt.Println("Retrieving the Message...")
	// Perform the query
	var strkey string = strconv.FormatInt(key, 10)
	db_query := fmt.Sprintf(MessageByIDQuery, strkey)
	rows, err := db.Query(db_query)
	if err != nil {
        fmt.Println("Query failed")
    }
	defer rows.Close()
	fmt.Println("Message retrieved successfully!")
	fmt.Println("Converting data to JSON...")
	// Convert data to a JSON format
	var m message
	for rows.Next(){
		err := rows.Scan(&m.Key, &m.Name, &m.CorrelationKey, &m.MessageId, &m.Timestamp)
		if err != nil {
			fmt.Println("Failed to scan row")
		}		
	}
	json, err := json.Marshal(m)
	if err != nil {
		fmt.Println("Failed to convert data to JSON")
	}
	return string(json)		
}
// retrieves a timer with the specified ProcessInstanceKey from the db
func RetrieveTimerByID(db *sql.DB, key int64) string {
	fmt.Println("Retrieving the timer...")
	// Perform the query
	var strkey string = strconv.FormatInt(key, 10)
	db_query := fmt.Sprintf(TimerByIDQuery, strkey)
	rows, err := db.Query(db_query)
	if err != nil {
        fmt.Println("Query failed")
    }
	defer rows.Close()
	fmt.Println("Message retrieved successfully!")
	fmt.Println("Converting data to JSON...")
	// Convert data to a JSON format
	var t timer
	for rows.Next(){
		err := rows.Scan(&t.Key, &t.Timestamp, &t.ProcessDefinitionKey, &t.ProcessInstanceKey, &t.ElementInstanceKey, &t.TargetElementId, &t.Duedate, &t.Repetitions)
		if err != nil {
			fmt.Println("Failed to scan row")
		}		
	}
	json, err := json.Marshal(t)
	if err != nil {
		fmt.Println("Failed to convert data to JSON")
	}
	return string(json)		
}
// retrieves an element with the given ProcessInstanceKey
func RetrieveElementByID(db *sql.DB, key int64) string {
	fmt.Println("Retrieving the element...")
	// Perform the query
	var strkey string = strconv.FormatInt(key, 10)
	db_query := fmt.Sprintf(ElementByIDQuery, strkey)
	rows, err := db.Query(db_query)
	if err != nil {
        fmt.Println("Query failed")
    }
	defer rows.Close()
	fmt.Println("Element retrieved successfully!")
	fmt.Println("Converting data to JSON...")
	// Convert data to a JSON format
	var e element
	for rows.Next(){
		err := rows.Scan(&e.Key, &e.ProcessInstanceKey, &e.ProcessDefinitionKey, &e.BpmnProcessId, &e.ElementId, &e.BpmnElementType, &e.Intent)
		if err != nil {
			fmt.Println("Failed to scan row")
		}		
	}
	json, err := json.Marshal(e)
	if err != nil {
		fmt.Println("Failed to convert data to JSON")
	}
	return string(json)		
}

//generates message for error
func GenerateErrorMessage(message string) string {
	fmt.Println("results empty. Returning error.")
	errorMessageValue := ErrorMessageValue {
		Error: message,
	}
	errorJSON, err := json.MarshalIndent(errorMessageValue, "", "  ")
	if err != nil {
		fmt.Println("generated error message could not be parsed.")
	}
	return string(errorJSON)
}

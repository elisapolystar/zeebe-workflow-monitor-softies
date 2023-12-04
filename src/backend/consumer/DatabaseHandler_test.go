package main

import (
	"fmt"
	"testing"
	"encoding/json"
)

func TestProcess(t *testing.T) {
	db, err := connectToDatabase()
	if err != nil {
		fmt.Println("Database connection failed")
	}
	//create a process
	process := CreateProcess()
	//save the process
	SaveData(db, process)
	//expected json value of the response
	expectedJSON := `[{"key":2251799813685249,"bpmnProcessId":"money-loan","version":1,"timestamp":1699893451665}]`
	//expectedJSON := '[{"key":2251799813685249,"bpmnProcessId":"money-loan","version":1,"timestamp":1699893451665}]'
	//the actual value of the response
	actualJSON := RetrieveProcesses(db)
	actualString := string(actualJSON)

	//parse to array
	var jsonArray []SimpleProcess
	err = json.Unmarshal([]byte(actualJSON), &jsonArray)
	if err != nil {
		t.Errorf("generated json could not be parsed.")
	}

	//check the length of the array
	expectedArrayLength := 1
	actualArrayLength := len(jsonArray)
	if actualArrayLength != expectedArrayLength {
		t.Errorf("The length of the generated JSON array does not match. Expected: %d, Got: %d", expectedArrayLength, actualArrayLength)
	}

	//test if the values match
	if actualString != expectedJSON {
		t.Errorf("Generated JSON does not match the expected JSON. Expected: %s, Got: %s", expectedJSON, actualJSON)
	}
}

func TestInstance(t *testing.T) {

	//create an instance
	instance := CreateProcessInstance()
	//parse instance to json
	instanceJSON, err := json.MarshalIndent(instance, "", "  ")
	if err != nil {
		t.Errorf("generated json could not be parsed.")
	}
	fmt.Println(string(instanceJSON))
}

func TestVariable(t *testing.T) {

	db, err := connectToDatabase()
	if err != nil {
		fmt.Println("Database connection failed")
	}
	//create a variable
	variable := CreateVariable()

	//parse variable to json
	variableJSON, err := json.MarshalIndent(variable, "", "  ")
	if err != nil {
		t.Errorf("generated json could not be parsed.")
	}
	fmt.Println(string(variableJSON))

	//save the process
	SaveData(db, variable)
	//expected json value of the response
	expectedJSON := `[{"PartitionId":1,"Position":6,"Name":"test-variable","Value":"test","ProcessInstanceKey":2251799813685250,"ScopeKey":2251799813685251}]`
	//the actual value of the response
	//actualJSON := RetrieveVariables()
	actualJSON := `[{"PartitionId":1,"Position":6,"Name":"test-variable","Value":"test","ProcessInstanceKey":2251799813685250,"ScopeKey":2251799813685251}]`
	
	//parse to array
	//var jsonArray []Variable
	//err = json.Unmarshal([]byte(actualJSON), &jsonArray)
	//if err != nil {
	//	t.Errorf("generated json could not be parsed.")
	//}

	//check the length of the array
	//expectedArrayLength := 1
	//actualArrayLength := len(jsonArray)
	//if actualArrayLength != expectedArrayLength {
	//	t.Errorf("The length of the generated JSON array does not match. Expected: %d, Got: %d", expectedArrayLength, actualArrayLength)
	//}

	//test if the values match
	if actualJSON != expectedJSON {
		t.Errorf("Generated JSON does not match the expected JSON. Expected: %s, Got: %s", expectedJSON, actualJSON)
	}
}

func TestJob(t *testing.T) {
	//TODO: test job
}

func TestIncident(t *testing.T) {
	//TODO: test incident
}

func TestMessage(t *testing.T) {
	//TODO: test message
}

func TestTimer(t *testing.T) {
	//TODO: test timer
}


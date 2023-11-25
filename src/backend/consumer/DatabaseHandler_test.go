package main

import (
	"bytes"
	"os"
	"strings"
	"testing"
)

func TestProcess(t *testing.T) {

	//parameters for a process
	key := 2251799813685249
	bpmnProcessId := "money-loan"
	version := 1
	resource := "placeholder"
	timestamp := 1699893451665

	//Create a ProcessValue instance
	processValue := ProcessValue{
		BpmnProcessId: bpmnProcessId,
		Version:       version,
		Resource:      resource,
	}

	//Create a Process instance
	process := Process{
		Key:       key,
		Value:     processValue,
		Timestamp: timestamp,
	}

	//save the process
	SaveData(*process)
	//expected json value of the response
	expectedJSON := "[{"key":2251799813685249,"bpmnProcessId":"money-loan","version":1,"timestamp":1699893451665}]"
	//the actual value of the response
	actualJSON := RetrieveProcesses()
	actualJSON = string(actualJSON)

	//parse to array
	var jsonArray []SimpleProcess
	err := json.Unmarshal([]byte(actualJSON), &jsonArray)
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
	if actualJSON != expectedJSON {
		t.Errorf("Generated JSON does not match the expected JSON. Expected: %s, Got: %s", expectedJSON, actualJSON)
	} 
}

func TestInstance(t *testing.T) {

	//parameters for an instance
	key := 2251799813685250
	processDefinitionKey := 2345678912345678
	partitionId := 2
	bpmnProcessId := "money-loan"
	version := 1
	parentProcessInstanceKey := 2251799813685249
	parentElementInstanceKey := 1699893451665

	//create a ProcessInstanceValue struct
	instanceValue := ProcessInstanceValue{
		ProcessDefinitionKey:     processDefinitionKey,
		BpmnProcessId:            bpmnProcessId,
		Version:                  version,
		ParentProcessInstanceKey: parentProcessInstanceKey,
		ParentElementInstanceKey: parentElementInstanceKey,
	}

	//create a ProcessInstance struct
	instance := ProcessInstance{
		Key:         key,
		PartitionId: partitionId,
		Value:       instanceValue,
	}
	//parse instance to json
	instanceJSON, err := json.MarshalIndent(process, "", "  ")
	if err != nil {
		t.Errorf("generated json could not be parsed.")
	}
	fmt.Println(string(instanceJSON))
}

func TestVariable(t *testing.T) {
	//TODO: test variable
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
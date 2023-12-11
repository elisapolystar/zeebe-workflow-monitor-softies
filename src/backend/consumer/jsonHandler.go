package main

import (
	"encoding/json"
	"fmt"
)

// Turn JSON data into struct of type Process
func parseProcessJson(msg []byte) (*Process, error) {

	var processItem Process
	err := json.Unmarshal(msg, &processItem)
	if err != nil {
		fmt.Println("Error JSON Unmarshalling")
		fmt.Println(err.Error())
	}

	return &processItem, nil
}

// Turn JSON data into struct of type ProcessInstance
func parseProcessInstanceJson(msg []byte) (*ProcessInstance, error) {

	var processInstanceItem ProcessInstance
	err := json.Unmarshal(msg, &processInstanceItem)
	if err != nil {
		fmt.Println("Error Unmarshalling processInstance JSON")
		fmt.Println(err.Error())
	}

	return &processInstanceItem, nil
}

// Turn JSON data into struct of type Variable
func parseVariableJson(msg []byte) (*Variable, error) {

	var variableItem Variable
	err := json.Unmarshal(msg, &variableItem)
	if err != nil {
		fmt.Println("Error Unmarshalling variable JSON")
		fmt.Println(err.Error())
	}

	return &variableItem, nil
}

// Turn JSON data into a struct of type Job
func parseJobJson(msg []byte) (*Job, error) {

	var jobItem Job
	err := json.Unmarshal(msg, &jobItem)
	if err != nil {
		fmt.Println("Error Unmarshalling job JSON")
		fmt.Println(err.Error())
	}

	return &jobItem, nil
}

// Turn JSON data into a struct of type Incident
func parseIncidentJson(msg []byte) (*Incident, error) {

	var incidentItem Incident
	err := json.Unmarshal(msg, &incidentItem)
	if err != nil {
		fmt.Println("Error Unmarshalling incident JSON")
		fmt.Println(err.Error())
	}

	return &incidentItem, nil
}

// Turn JSON data into a struct of type Message
func parseMessageJson(msg []byte) (*Message, error) {

	var messageItem Message
	err := json.Unmarshal(msg, &messageItem)
	if err != nil {
		fmt.Println("Error Unmarshalling message JSON")
		fmt.Println(err.Error())
	}

	return &messageItem, nil
}

// Turn JSON data into struct of type Timer
func parseTimerJson(msg []byte) (*Timer, error) {

	var timerItem Timer
	err := json.Unmarshal(msg, &timerItem)
	if err != nil {
		fmt.Println("Error Unmarshalling timer JSON")
		fmt.Println(err.Error())
	}

	return &timerItem, nil
}

func parseZeebeJson(msg []byte) (*Zeebe, error) {

	var zeebeItem Zeebe
	err := json.Unmarshal(msg, &zeebeItem)
	if err != nil {
		fmt.Println("Error Unmarshalling Zeebe item JSON")
		fmt.Println(err.Error())
	}

	zeebeItem.Active = true

	return &zeebeItem, nil
}

func parseElementJson(msg []byte) (*Element, error) {

	var elementItem Element
	err := json.Unmarshal(msg, &elementItem)
	if err != nil {
		fmt.Println("Error Unmarshalling Element item JSON")
		fmt.Println(err.Error())
	}

	return &elementItem, nil
}

// Turn a process request message from the frontend into a struct
func parseProcessRequest(msg []byte) (*ProcessRequest, error) {

	var communicationItem ProcessRequest
	err := json.Unmarshal(msg, &communicationItem)
	if err != nil {
		fmt.Println("Le` error in the front json parsings :-(")
		fmt.Println(err.Error())
	}

	return &communicationItem, nil
}

// Turn a instance request message into a struct
func parseInstanceRequest(msg []byte) (*InstanceRequest, error) {

	var instanceMessage InstanceRequest
	err := json.Unmarshal(msg, &instanceMessage)
	if err != nil {
		fmt.Println("Error turning json to struct: ")
		fmt.Print(err.Error())
	}

	return &instanceMessage, nil
}

// Add JSONs together to form one json that represents an instance
func concatenateInstanceJSON(json1, json2, json3, json4, json5 []byte) (*[]byte, error) {
	var process ProcessForFrontend
	var elements []interface{}
	var variables []interface{}
	var timers []interface{}
	var incidents []interface{}

	err1 := json.Unmarshal(json1, &process)
	if err1 != nil {
		fmt.Println("Error turning json to struct: ", err1)
		return nil, err1
	}

	processContainer := ProcessContainer{
		Process: process,
	}

	err2 := json.Unmarshal(json2, &elements)
	if err2 != nil {
		fmt.Println("Error turning instance to struct: ", err2)
		return nil, err2
	}

	err3 := json.Unmarshal(json3, &variables)
	if err3 != nil {
		fmt.Println("Error turning variables to struct: ", err3)
		return nil, err3
	}

	err4 := json.Unmarshal(json4, &timers)
	if err4 != nil {
		fmt.Println("Error turning timers to struct: ", err4)
		return nil, err4
	}

	err5 := json.Unmarshal(json5, &incidents)
	if err5 != nil {
		fmt.Println("Error turning incidents to struct: ", err5)
		return nil, err5
	}

	combinedItem := struct {
		ProcessContainer
		Elements  []interface{}
		Variables []interface{}
		Timers    []interface{}
		Incidents []interface{}
	}{
		ProcessContainer: processContainer,
		Elements:         elements,
		Variables:        variables,
		Timers:           timers,
		Incidents:        incidents,
	}

	combinedJSON, err := json.Marshal(combinedItem)
	if err != nil {
		fmt.Println("Failed to marshal json: ", err)
		return nil, err
	}
	return &combinedJSON, nil
}

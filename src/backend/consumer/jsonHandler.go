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

// Turn a process struct into JSON data
func structToJson(data interface{}) (string, error) {

	jsonData, err := json.Marshal(data)
	if err != nil {
		fmt.Println("Error: ", err)
		fmt.Println(err.Error())
	}

	jsonString := string(jsonData)
	return jsonString, err
}

// Turn a message from the frontend into a struct
func parseCommunicationItem(msg []byte) (*FrontCommunication, error) {

	var communicationItem FrontCommunication
	err := json.Unmarshal(msg, &communicationItem)
	if err != nil {
		fmt.Println("Le` error in the front json parsings :-(")
		fmt.Println(err.Error())
	}

	return &communicationItem, nil
}
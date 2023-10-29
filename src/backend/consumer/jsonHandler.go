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

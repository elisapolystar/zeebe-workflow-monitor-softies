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

// Add two JSONs together
func concatenateJSON(json1, json2, json3, json4, json5 []byte) (*[]byte, error) {
	var process ProcessForFrontend
	var elements ElementsContainer
	var variables VariablesContainer
	var timers TimersContainer
	var incidents IncidentsContainer

	err1 := json.Unmarshal(json1, &process)
	if err1 != nil {
		fmt.Println("Error turning json to struct: ", err1)
		return nil, err1
	}

	processContainer := ProcessContainer{
		Process: process,
	}

	/*fmt.Println()
	fmt.Println("Processcontainer item: ", processContainer.Process.BpmnProcessId)
	fmt.Println("Processcontainer item: ", processContainer.Process.Key)
	fmt.Println("jarkko")*/

	err2 := json.Unmarshal(json2, &elements)
	if err2 != nil {
		fmt.Println("Error turning elements json to struct: ", err2)
		return nil, err2
	}

	err3 := json.Unmarshal(json3, &variables)
	if err3 != nil {
		fmt.Println("Error turning json to struct: ", err3)
		return nil, err3
	}

	err4 := json.Unmarshal(json4, &timers)
	if err4 != nil {
		fmt.Println("Error turning json to struct: ", err4)
		return nil, err4
	}

	err5 := json.Unmarshal(json5, &incidents)
	if err5 != nil {
		fmt.Println("Error turning json to struct: ", err5)
		return nil, err5
	}

	combinedItem := struct {
		ProcessContainer
		//ElementsContainer
		VariablesContainer
		TimersContainer
		IncidentsContainer
	}{
		ProcessContainer: processContainer,
		//ElementsContainer:  elements,
		VariablesContainer: variables,
		TimersContainer:    timers,
		IncidentsContainer: incidents,
	}

	fmt.Println()
	fmt.Println("combined item process field: ", combinedItem.ProcessContainer.Process.Key)
	//fmt.Println("combined item fields: ", combinedItem.Elements)

	combinedJSON, err := json.Marshal(combinedItem)
	if err != nil {
		fmt.Println("Failed to marshal json: ", err)
		return nil, err
	}

	fmt.Println()
	fmt.Println("1 combinded json ", string(combinedJSON))
	fmt.Println()

	return &combinedJSON, nil
}

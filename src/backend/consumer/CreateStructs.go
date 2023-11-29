package main

import (
	"fmt"
	"encoding/json"
)

func CreateProcess() Process {
	//parameters for a process
	var key int64 = 2251799813685249
	bpmnProcessId := "money-loan"
	var version int64 = 1
	//version := 1
	resource := "placeholder"
	var timestamp int64 = 1699893451665

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
	return process
}
func CreateProcessInstance() Zeebe {
	//parameters for an instance
	var key int64 = 2251799813685250
	var processDefinitionKey int64 = 2345678912345678
	var partitionId int64 = 2
	bpmnProcessId := "money-loan"
	var version int64 = 1
	var parentProcessInstanceKey int64 = 2251799813685249
	var parentElementInstanceKey int64 = 1699893451665
	active := true

	//create a ProcessInstanceValue struct
	instanceValue := ProcessInstanceValue{
		ProcessDefinitionKey:     processDefinitionKey,
		BpmnProcessId:            bpmnProcessId,
		Version:                  version,
		ParentProcessInstanceKey: parentProcessInstanceKey,
		ParentElementInstanceKey: parentElementInstanceKey,
	}

	//create a ProcessInstance struct
	instance := Zeebe{
		Key:         key,
		PartitionId: partitionId,
		Value:       instanceValue,
		Active:      active,
	}
	return instance
}
func CreateVariable() Variable {
	//variable parameters
	name := "test-variable"
	value := "test"
	var processInstanceKey int64 = 2251799813685250
	var scopeKey int64 = 2251799813685251
	var partitionId int64 = 1
	var position int64 = 6

	//Create a variableValue instance
	variableValue := VariableValue{
		Name:				name,
		Value:				value,
		ProcessInstanceKey: processInstanceKey,
		ScopeKey:			scopeKey,
	}

	//Create a Variable instance
	variable := Variable{
		PartitionId: partitionId,
		Position:	 position,
		Value: 		 variableValue,
	}
	return variable
}
func CreateJob() Job {
	var key int64 = 2251799813685251
	var timestamp int64 = 1699893451665
	var processInstanceKey int64 = 2251799813685250
	var elementInstanceKey int64 = 2251799813685252
	jobType := "test-job"
	worker := "test-worker"
	var retries int64 = 3

	jobValue := JobValue{
		ProcessInstanceKey: processInstanceKey,
		ElementInstanceKey: elementInstanceKey,
		JobType:			jobType,
		Worker:				worker,
		Retries:			retries,
	}

	job := Job{
		Key:	   key,
		Timestamp: timestamp,
		Value:	   jobValue,
	}
	return job
}
func CreateIncident() Incident {

	bpmnProcessId := "money-loan"
	var processDefinitionKey int64 = 2251799813685249
	var processInstanceKey int64 = 2251799813685250
	var elementInstanceKey int64 = 2251799813685252
	var jobKey int64 = 2251799813685251
	errorType := "test-error"
	errorMessage := "placeholder"
	var key int64 = 2251799813685253

	incidentValue := IncidentValue{
		BpmnProcessId:bpmnProcessId,
		ProcessDefinitionKey:processDefinitionKey,
		ProcessInstanceKey:processInstanceKey,
		ElementInstanceKey:elementInstanceKey,
		JobKey:jobKey,
		ErrorType:errorType,
		ErrorMessage:errorMessage,
	}

	incident := Incident{
		Key:key,
		Value:incidentValue,
	}
	return incident
}
func CreateMessage() Message {
	name := "test-name"
	correlationKey := "test-correlation-key"
	messageId := "id"
	var key int64 = 2251799813685253
	var timestamp int64 = 1699893451665

	messageValue := MessageValue{
		Name:name,
		CorrelationKey:correlationKey,
		MessageId:messageId,
	}

	message := Message{
		Key:key,
		Value:messageValue,
		Timestamp:timestamp,
	}
	return message
}
func CreateTimer() Timer {
	var processDefinitionKey int64 = 2251799813685249
	var processInstanceKey int64 = 2251799813685250
	var elementInstanceKey int64 = 2251799813685252
	targetElementId := "id"
	var dueDate int64 = 1699893451665
	var repetitions int64 = 3
	var key int64 = 2251799813685254
	var timestamp int64 = 1699893451665

	timerValue := TimerValue{
		ProcessDefinitionKey:processDefinitionKey,
		ProcessInstanceKey:processInstanceKey,
		ElementInstanceKey:elementInstanceKey,
		TargetElementId:targetElementId,
		Duedate:dueDate,
		Repetitions:repetitions,
	}
	timer := Timer{
		Key:key,
		Timestamp:timestamp,
		Value:timerValue,
	}
	return timer
}

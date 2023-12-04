package main

func CreateProcess() Process {
	//parameters for a process
	var key int64 = 2345678912345678
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
	const processDefinitionKey int64 = 2345678912345678
	const processInstanceKey int64 = 2593275030505839
	const partitionId int64 = 2
	const bpmnProcessId string = "money-loan"
	const version int64 = 1
	const timestamp = 1700323209172
	const active bool = true

	//create a ProcessInstanceValue struct
	instanceValue := ZeebeValue{
		ProcessDefinitionKey: processDefinitionKey,
		ProcessInstanceKey:   processInstanceKey,
		BpmnProcessId:        bpmnProcessId,
		Version:              version,

	}

	//create a ProcessInstance struct
	instance := Zeebe{
		PartitionId: partitionId,
		Value:       instanceValue,
		Timestamp: timestamp,
		Active:      active,
	}
	return instance

}
func CreateVariable() Variable {
	//variable parameters
	name := "test-variable"
	value := "test"
	var processInstanceKey int64 = 2593275030505839
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
	var processInstanceKey int64 = 2593275030505839
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
	var processInstanceKey int64 = 2593275030505839
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
	var processInstanceKey int64 = 2593275030505839
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

func CreateElement() Element {
	var key int64 = 1
	var intent string = "first"
	var processinstancekey int64 = 2593275030505839
	var processdefinitionkey int64 = 2251799813685249
	var bpmnprocessid string = "placeholder"
	var elementid string = "placeholder id"
	var bpmnelementtype string = "placeholder elementtype"

	elementvalue := ElementValue{
		ProcessInstanceKey:processinstancekey,
		ProcessDefinitionKey:processdefinitionkey,
		BpmnProcessId:bpmnprocessid,
		ElementId:elementid,
		BpmnElementType:bpmnelementtype,
	}
	element := Element{
		Key:key,
		Value:elementvalue,
		Intent:intent,
	}
	return element
}
func UpdateElement() Element {
	var key int64 = 1
	var intent string = "second"
	var processinstancekey int64 = 2593275030505839
	var processdefinitionkey int64 = 2251799813685249
	var bpmnprocessid string = "placeholder"
	var elementid string = "placeholder id"
	var bpmnelementtype string = "placeholder elementtype"

	elementvalue := ElementValue{
		ProcessInstanceKey:processinstancekey,
		ProcessDefinitionKey:processdefinitionkey,
		BpmnProcessId:bpmnprocessid,
		ElementId:elementid,
		BpmnElementType:bpmnelementtype,
	}
	element := Element{
		Key:key,
		Value:elementvalue,
		Intent:intent,
	}
	return element	
}
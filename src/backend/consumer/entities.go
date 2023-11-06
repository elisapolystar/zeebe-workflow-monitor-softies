package main

type Process struct {
	Key       int64        `json:"key"`
	Value     ProcessValue `json:"value"`
	Timestamp int64        `json:"timestamp"`
}

type ProcessValue struct {
	BpmnProcessId string `json:"bpmnProcessId"`
	Version       int64  `json:"version"`
	Resource      string `json:"resource"`
}

type ProcessInstance struct {
	// state, start, end?
	Key         int64                `json:"key"`
	PartitionId int64                `json:"partitionId"`
	Value       ProcessInstanceValue `json:"value"`
}

type ProcessInstanceValue struct {
	ProcessDefinitionKey     int64  `json:"processDefinitionKey"`
	BpmnProcessId            string `json:"bpmnProcessId"`
	Version                  int64  `json:"version"`
	ParentProcessInstanceKey int64  `json:"parentProcessInstanceKey"`
	ParentElementInstanceKey int64  `json:"parentElementInstanceKey"`
}

type Variable struct {
	//id?
	PartitionId int64         `json:"partitionId"`
	Position    int64         `json:"position"`
	Value       VariableValue `json:"value"`
}

type VariableValue struct {
	Name               string `json:"name"`
	Value              string `json:"value"`
	ProcessInstanceKey int64  `json:"processInstanceKey"`
	ScopeKey           int64  `json:"scopeKey"`
}

type Job struct {
	//state?
	Key       int64    `json:"key"`
	Timestamp int64    `json:"timestamp"`
	Value     JobValue `json:"value"`
}

type JobValue struct {
	ProcessInstanceKey int64  `json:"processInstanceKey"`
	ElementInstanceKey int64  `json:"elementInstanceKey"`
	JobType            string `json:"type"`
	Worker             string `json:"worker"`
	Retries            int64  `json:"retries"`
}

type Incident struct {
	Key   int64         `json:"key"`
	Value IncidentValue `json:"value"`
}

type IncidentValue struct {
	BpmnProcessId        string `json:"bpmnProcessId"`
	ProcessDefinitionKey int64  `json:"processDefinitionKey"`
	ProcessInstanceKey   int64  `json:"processInstanceKey"`
	ElementInstanceKey   int64  `json:"elementInstanceKey"`
	JobKey               int64  `json:"jobKey"`
	ErrorType            string `json:"errorType"`
	ErrorMessage         string `json:"errorMessage"`
}

type Message struct {
	Key       int64        `json:"key"`
	Value     MessageValue `json:"value"`
	Timestamp int64        `json:"timestamp"`
}

type MessageValue struct {
	Name           string `json:"name"`
	CorrelationKey string `json:"correlationKey"`
	MessageId      string `json:"messageId"`
}

package main

type Process struct {
	Key       int64 `json:"key"`
	Value     Value `json:"value"`
	Timestamp int64 `json:"timestamp"`
}

type Value struct {
	BpmnProcessId string `json:"bpmnProcessId"`
	Version       int64  `json:"version"`
	Resource      string `json:"resource"`
}

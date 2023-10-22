package main

import (
	"encoding/json"
	"fmt"
)

func parseJson(msg []byte) (*Process, error) {

	var processItem Process
	err := json.Unmarshal(msg, &processItem)
	if err != nil {
		fmt.Println("Error JSON Unmarshalling")
		fmt.Println(err.Error())
	}

	return &processItem, nil
}

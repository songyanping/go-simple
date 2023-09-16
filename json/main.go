package main

import (
	"encoding/json"
	"fmt"
)

func main() {
	var body string
	body = `
{
	"source": "k8s-err-with-logs-events",
	"data": {
		"APIVersion": "v1",
		"Action": "",
		"Actions": null,
		"Cluster": "ft-k8s",
		"Code": "",
		"Count": 1,
		"Error": "",
		"Kind": "Pod",
		"Level": "error",
		"Messages": ["Error: ImagePullBackOff"],
		"Name": "nginx-test-767c779d45-z8ht6",
		"Namespace": "default",
		"Reason": "Failed",
		"Recommendations": null,
		"Resource": "v1/pods",
		"TimeStamp": "2023-05-10T07:40:29Z",
		"Title": "v1/pods error",
		"Type": "error",
		"Warnings": null
	},
	"timeStamp": "0001-01-01T00:00:00Z"
}
`

	//fmt.Println(body)
	var event Event
	json.Unmarshal([]byte(body), &event)

	msg := event.Data["Messages"]
	fmt.Println(msg)
	//var m []map[string]string

	msg1 := msg.([]interface{})[0]
	fmt.Printf("%T", msg1)
	fmt.Println("=========")
	fmt.Println(msg1)

	//json.Unmarshal([]byte(msg.(string)), &m)
	//
	//fmt.Println(m)
}

type Event struct {
	Source    string                 `json:"source"`
	Data      map[string]interface{} `json:"data"`
	TimeStamp string                 `json:"timeStamp"`
}

package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"text/template"
	"time"
)

type Alert struct {
	Title    string            `json:"title"`
	RuleName string            `json:"ruleName"`
	Level    string            `json:"level"`
	Type     string            `json:"type"`
	DateTime time.Time         `json:"datetime"`
	Location string            `json:"location"`
	Business map[string]string `json:"business"`
	Details  string            `json:"details"`
	Notes    string            `json:"notes"`
}

type Alert2 struct {
	Title    string `json:"title"`
	RuleName string `json:"ruleName"`
	Level    string `json:"level"`
	Type     string `json:"type"`
	DateTime string `json:"datetime"`
	Location string `json:"location"`
	Business string `json:"business"`
	Details  string `json:"details"`
	Notes    string `json:"notes"`
}

type Writer interface {
	Write(p []byte) (n int, err error)
}

type Bus struct {
}

func (b *Bus) Write(p []byte) (n int, err error) {
	return n, err
}

func main() {
	business := map[string]string{
		"service":     "bff",
		"environment": "ft1",
	}
	midd := Alert{
		RuleName: "alert test",
		Level:    "p1",
		Type:     "mysql",
		DateTime: time.Now(),
		Location: "local",
		Business: business,
		Details:  "details",
		Notes:    "notes",
	}

	tmpl, err := template.ParseFiles("template/demo.tmpl")
	if err != nil {
		fmt.Println(err.Error())
	}
	var buf bytes.Buffer
	if err := tmpl.Execute(&buf, midd); err != nil {
		fmt.Println(err.Error())
	}
	fmt.Println(string(buf.Bytes()))

	var alert2 Alert2
	json.Unmarshal(buf.Bytes(), &alert2)
	fmt.Println(alert2)

	//a, _ := json.Marshal(buf.Bytes())
	//fmt.Printf(string(a))
	/*
			{"ruleName": "alert test"
			"level": "p1"
			...
		    }
	*/
	//以上是输出示例，我想要json格式字符串或byte

}

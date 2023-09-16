package main

import (
	"bytes"
	"fmt"
	"text/template"
)

type KV map[string]string

type Alert struct {
	Level    string
	Type     string
	Business map[string]string
}

func main() {
	alert := Alert{
		Level:    "P1",
		Type:     "mysql",
		Business: map[string]string{"env": "pd", "service": "bff=h5"},
	}
	tmpl, err := template.ParseFiles("./hello.tmpl")
	if err != nil {
		fmt.Println(err)
	}
	var buf bytes.Buffer
	if err := tmpl.Execute(&buf, alert); err != nil {
		fmt.Println(err)
	}
	fmt.Println(buf.String())

}

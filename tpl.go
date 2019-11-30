package main

import (
	"bytes"
	"fmt"
	"os"
	"text/template"
)

type Person struct {
	Name string
	Age  int
}

func main() {
	var cmd *bytes.Buffer
	//p := Person{"longshuai",23}
	p := &map[string]string{
		"Name": "11",
		"Age":  "22",
	}
	tmpl, _ := template.New("test").Parse("Name: {{.Name}}, Age: {{.Age}}")
	_ = tmpl.Execute(os.Stdout, p)
	fmt.Printf("%#+v", cmd)
}

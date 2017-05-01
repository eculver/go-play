package main

import (
	"encoding/json"
	"fmt"
	"log"

	"gopkg.in/yaml.v2"
)

type Foo struct {
	FieldName      string
	OtherFieldName string `yaml:"ofn"`
}

var rawJSON = `{
	"fieldName": "foo",
	"OtherFieldName": "bar"
}`

var rawYAML = `
fieldname: foo
ofn: bar`

func main() {
	f1 := Foo{}
	if err := json.Unmarshal([]byte(rawJSON), &f1); err != nil {
		log.Fatalf("JSON error: %v", err)
	}
	f2 := Foo{}
	if err := yaml.Unmarshal([]byte(rawYAML), &f2); err != nil {
		log.Fatalf("YAML error: %v", err)
	}

	fmt.Printf("JSON - FieldName: %v, OtherFieldName: %v\n", f1.FieldName, f1.OtherFieldName)
	fmt.Printf("YAML - FieldName: %v, OtherFieldName: %v\n", f2.FieldName, f2.OtherFieldName)

}

package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"os"
	"text/template"
)

type Response struct {
	ID   int64  `json:"Id"`
	Msg  string `json:"msg"`
	Data string `json:"data"`
}

var rawElement = []byte(` {"Id":200,"msg":"success","data":"Macbook"}`)

func main() {
	var raw interface{}
	//var aa Response
	//buffer := new(bytes.Buffer)
	rdr := bytes.NewReader(rawElement)
	dec := json.NewDecoder(rdr)
	dec.UseNumber()

	if rawErr := dec.Decode(&raw); rawErr != nil {
		fmt.Errorf("unable to read inspect data: %v", rawErr)
	}
	fmt.Println(raw)

	tmpl := template.New("t1")

	tmplMissingKey, _ := tmpl.Option("missingkey=error").Parse("{{ .Id }}")
	if rawErr := tmplMissingKey.Execute(os.Stdout, raw); rawErr != nil {
		fmt.Errorf("Template parsing error: %v", rawErr)
	}

	//fmt.Println(buffer)
}

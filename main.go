package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"log"
	"strings"
	"text/template"
)

type typedElem struct {
	ID   string
	Size float64
}

func main() {
	fmt.Println("hello")
	/*typedElem := struct {
		ID string `json:"Id"`
	}{"ad3"}*/
	testcases := []struct {
		raw []byte
		exp string
	}{
		{raw: []byte(`{"Id": "ad3", "Size": 53317}`), exp: "53317 ad3\n"},
		{raw: []byte(`{"Id": "ad3", "Size": 53317.102}`), exp: "53317.102 ad3\n"},
		{raw: []byte(`{"Id": "ad3", "Size": 53317.0}`), exp: "53317.0 ad3\n"},
	}
	f := "{{.ID}} ++++ {{.Size}}\n"
	if strings.Contains(f, ".ID") {
		f = strings.Replace(f, ".ID", ".Id", -1)
	}
	buf := new(bytes.Buffer)
	tmpl, err := template.New("test").Parse(f)
	if err != nil {
		log.Fatalf(err.Error())
	}
	for _, test := range testcases {
		rdr := bytes.NewReader(test.raw)
		dec := json.NewDecoder(rdr)
		dec.UseNumber()
		var raw interface{}
		if rawErr := dec.Decode(&raw); rawErr != nil {
			log.Fatalf("unable to read inspect data: %v", rawErr)
		}
		rawT := raw
		fmt.Println("raw = : ", raw)
		fmt.Println("rawT = : ", rawT)
		tmplMissingKey := tmpl.Option("missingkey=error")
		if rawErr := tmplMissingKey.Execute(buf, rawT); rawErr != nil {
			log.Fatalf("Template parsing error: %v", rawErr)
		}

		fmt.Println(buf.String())

	}

}

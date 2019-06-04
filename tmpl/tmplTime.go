package main

import (
	"fmt"
	"github.com/sirupsen/logrus"
	"os"
	"text/template"
	"time"
)

var muban string = `
{{ .TT.UTC }}
{{.TT.UTC.Format "2006-01-02T15:04:05.781Z" }}
`

type test struct {
	TT time.Time
}

func main() {
	//model()
	TT := time.Now()
	fmt.Println(TT.UTC())
	fmt.Println(TT.String())

	tmpl, err := template.New("tmpl").Parse(muban)
	if err != nil {
		logrus.Error(err)
	}
	err = tmpl.Execute(os.Stdout, test{TT})
	if err != nil {
		fmt.Println(err)
	}
}

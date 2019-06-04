package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"github.com/sirupsen/logrus"
	"io/ioutil"
)

var CfgTmpl = `
 - name: {{ .Name }}
  rules:
  - alert: {{ .Alert }}
    expr: {{ .Expr }} {{ .Compare }} {{ .Threshold }}
    for: {{ .ForTime }}
    labels:
      severity: warning
      job: alertmanager
      layer: site
    annotations:
	  summary: {{ .Summary }} {{ .Compare }} {{ .Threshold }}
      description: {{ .Description }}
`

type CfgStu struct {
	Name        string
	Alert       string
	Expr        string
	Compare     string
	Threshold   float64
	ForTime     string
	Summary     string
	Description string
}

type Config struct {
	Sites []string
	Rules []CfgStu
}

var (
	// ConfigFile defines something.
	ConfigFile = flag.String("config", "tmpl/config.json", "Config File")
)

func main() {
	flag.Parse()
	config, err := LoadConfig(*ConfigFile)
	if err != nil {
		logrus.Fatalf("unable to load config: %v", err)
	}
	fmt.Println(config)
}

// LoadConfig loads configs from local file.
func LoadConfig(cfgPath string) (*Config, error) {
	contents, err := ioutil.ReadFile(cfgPath)
	if err != nil {
		return nil, fmt.Errorf("error reading config file %q: %v", cfgPath, err)
	}

	var config Config
	err = json.Unmarshal(contents, &config)
	if err != nil {
		return nil, fmt.Errorf("error unmarshalling config: %v", err)
	}

	return &config, nil
}

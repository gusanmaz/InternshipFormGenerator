package globals

import (
	"flag"
	"fmt"
	"gopkg.in/yaml.v2"
	"io/ioutil"
)

type Config struct {
	OutputDirPath        string `yaml:"outputDirPath"`
	FormTemplateFilePath string `yaml:"formTemplateFilePath"`
	DispatchTemplateFilePath string `yaml:"dispatchTemplateFilePath"`
	FormTemplateFileName string `yaml:"formTemplateFileName"`
	DispatchTemplateFileName string `yaml:"dispatchTemplateFileName"`
	CSVFilePath          string `yaml:"CSVFilePath"`
	IntegrityCheck       bool `yaml:"integrityCheck"`
	JSONTestBaseURL      string `yaml:"JSONTestBaseURL"`
	InstitutionName      string `yaml:"institutionName"`
	DepartmentName       string `yaml:"departmentName"`
	Head                 string `yaml:"head"`
	Member1              string `yaml:"member1"`
	Member2              string `yaml:"member2"`
	EducationYear        string `yaml:"educationYear"`
	BeginDate            string `yaml:"beginDate"`
	EndDate              string `yaml:"endDate"`
	OnlineText           string `yaml:"onlineText"`
	OnsiteText        string `yaml:"onsiteText"`
	OnlineOnsitetext  string `yaml:"onlineOnsiteText"`
}

var configFilePath = "config.yaml"
var Globals Config

func init(){
	configFilePathP := flag.String("config", "config.yaml", "Configuration file")
	flag.Parse()
	configFilePath = *configFilePathP

	yamlFile, err := ioutil.ReadFile(configFilePath)
	if err != nil {
		fmt.Printf("Error reading YAML file: %s\n", err)
		return
	}

	err = yaml.Unmarshal(yamlFile, &Globals)
	if err != nil {
		fmt.Printf("Error parsing YAML file: %s\n", err)
	}
}

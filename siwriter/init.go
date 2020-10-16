package siwriter

import (
	"errors"
	"github.com/gusanmaz/InternshipFormGenerator2/globals"
	"github.com/gusanmaz/InternshipFormGenerator2/templates"
	"io/ioutil"
	"os"
)

var formTemplateText string
var tableRowTemplateText string
var dispatchTemplateText string
var g globals.Config

func init(){
	g = globals.Globals
	var err error
	formTemplateText, err = getTemplateText(g.FormTemplateFilePath)
	if err != nil{
		formTemplateText = templates.FormFileContent

	}
	//tableRowTemplateText = getTemplateText(g.FormTemplateFilePath)
	dispatchTemplateText, err = getTemplateText(g.FormTemplateFilePath)
	if err != nil{
		dispatchTemplateText = templates.DispatchFileContent
	}
}

func getTemplateText(templateFilePath string)(string, error){
	templateFile, err := os.Open(templateFilePath)
	if err == nil{
		bytes, _ := ioutil.ReadAll(templateFile)
		templateFile.Close()
		return string(bytes), nil
	}
	return "", errors.New("File not found")
}

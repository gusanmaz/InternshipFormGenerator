package main

import (
	"io"
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"
)

// Reads all .txt files in the current folder
// and encodes them as strings literals in textfiles.go

var templatesDir string = "./templates"
var goVariablesFilePath string = "templatesText.go"
var variableSuffix string = "FileContent"
func main() {
	fs, _ := ioutil.ReadDir(templatesDir)
	out, _ := os.Create(filepath.Join(templatesDir, goVariablesFilePath))
	out.Write([]byte("package templates \n\nconst (\n"))
	for _, f := range fs {
		if strings.HasSuffix(f.Name(), ".fodt") {
			variableName := strings.TrimSuffix(f.Name(), ".fodt")
			variableName = strings.Title(variableName)
			variableName = variableName + variableSuffix
			//strings.TrimSuffix(f.Name(), ".fodt") + variableSuffix
			out.Write([]byte(variableName + " = `"))
			f, _ := os.Open(filepath.Join(templatesDir,f.Name()))
			io.Copy(out, f)
			out.Write([]byte("`\n"))
		}
	}
	out.Write([]byte(")\n"))
}

package siwriter

import (
	"github.com/gusanmaz/InternshipFormGenerator/si"
	"github.com/gusanmaz/InternshipFormGenerator/templates"
	"os"
	"os/exec"
	"path/filepath"
	"strings"
	"text/template"
)

type FormData struct{
	Department string
	Name string
	ID string
	Begin string
	End string
	DaysRequested string
	DaysAccepted string
	Member1 string
	Head string
	Type string
}

func SubjectFileName(s string) string{
	r := strings.Replace(s, "(", "", -1)
	r = strings.Replace(r, ")", "", -1)
	r = strings.Replace(r, "-", "_", -1)
	r = strings.Replace(r, "/", "", -1)
	r = strings.Replace(r, " ", "_", -1)
	return r
}


func CreateAllFormFiles(siArr []si.StudentInternship){
	outputDir := g.OutputDirPath
	for _, si := range siArr{
		CreateFormFiles(outputDir, formTemplateText, si)
	}
}

func GenerateFormData(siData si.StudentInternship) FormData{
	ret := FormData{}
	ret.ID = siData.ID
	ret.Name = strings.Title(strings.ToLower(siData.Name)) + " " + strings.ToUpper(siData.Surname)
	ret.Begin = siData.BeginDateS
	ret.End   = siData.EndDateS
	ret.DaysAccepted = siData.DaysAcceptedS
	ret.DaysRequested = siData.DaysRequestedS
	ret.Member1 = g.Member1
	ret.Head = g.Head
	ret.Department = g.DepartmentName
	ret.Type = ""
	return ret
}

func CreateFormFiles(outputDir, draftContent string, siData si.StudentInternship) bool{
	if siData.InternshipType == si.Onsite{
		return false
	}

	fd := GenerateFormData(siData)

	fileID := siData.ID + "_" + SubjectFileName(siData.Subject)
	odtFileName := fileID + ".fodt"
	pdfFileName := fileID + ".pdf"
	DirExists(outputDir, true)

	formFile, err := os.Create(filepath.Join(g.OutputDirPath, odtFileName))
	defer formFile.Close()

	formTemp := template.New(g.FormTemplateFileName)
	_, err = os.Stat(g.FormTemplateFilePath)
	if os.IsExist(err){
		_, err = formTemp.ParseFiles(g.DispatchTemplateFilePath)
	}else{
		formTemp.Parse(templates.FormFileContent)
	}
	formTemp.Execute(formFile, fd)

	cmd := exec.Command("libreoffice", "--headless", "--convert-to", "pdf", odtFileName, "output", pdfFileName)
	cmd.Dir = filepath.Join(".", outputDir)
	err = cmd.Run()
	return true
}

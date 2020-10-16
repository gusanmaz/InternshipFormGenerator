package siwriter

import (
	"fmt"
	"github.com/gusanmaz/InternshipFormGenerator/si"
	"github.com/gusanmaz/InternshipFormGenerator/templates"
	"os"
	"os/exec"
	"path/filepath"
	"strings"
	"text/template"
	"time"
	"unicode"
)

type RowData struct{
	Order         string
	ID            string
	Name          string
	Surname       string
	Institution   string
	BeginDate     string
	EndDate       string
	DaysRequested string
	DaysAccepted  string
}

type DispatchData struct{
	Date             string
	EducationYear    string
	Department       string
	DepartmentAllCap string
	Head             string
	Member1          string
	Member2          string
	Text             string
	Rows             []RowData
}

func GenerateDispatchData(siArr []si.StudentInternship) DispatchData{
	onlineRows  := false
	offlineRows := false
	rows := make([]RowData, len(siArr))

	for i, siElem := range siArr{
		rows[i].Order = fmt.Sprintf("%v", i + 1)
		rows[i].ID = siElem.ID
		rows[i].Name = strings.Title(strings.ToLower(siElem.Name))
		rows[i].Surname = strings.ToUpper(siElem.Surname)
		rows[i].BeginDate = siElem.BeginDateS
		rows[i].EndDate = siElem.EndDateS
		rows[i].DaysAccepted = siElem.DaysAcceptedS
		rows[i].DaysRequested = siElem.DaysRequestedS
		rows[i].Institution = siElem.Location

		if siElem.InternshipType == si.Online || siElem.InternshipType == si.OnlineVerifiable{
			onlineRows = true
		}
		if siElem.InternshipType == si.Onsite{
			offlineRows = true
		}
	}

	text := ""
	if onlineRows && offlineRows{
		text = g.OnlineOnsitetext
	}

	if onlineRows && !offlineRows{
		text = g.OnlineText
	}

	if !onlineRows && offlineRows{
		text = g.OnsiteText
	}

	d := time.Now().Day()
	m := int(time.Now().Month())
	y := time.Now().Year()
	date := fmt.Sprintf("%v.%v.%v", d, m ,y)
	ret := DispatchData{
		Date:             date,
		EducationYear:    g.EducationYear,
		Department:       g.DepartmentName,
		DepartmentAllCap: strings.ToUpperSpecial(unicode.TurkishCase, g.DepartmentName),
		Head:             g.Head,
		Member1:          g.Member1,
		Member2:          g.Member2,
		Text:             text,
		Rows:             rows,
	}
	return ret
}

func DispatchFODTFileName()string{
	t := time.Now()
	return fmt.Sprintf("%v_%v_%v_staj_yazisi.fodt", t.Day(), int(t.Month()), t.Year())
}

func DispatchPDFFileName()string{
	t := time.Now()
	return fmt.Sprintf("%v_%v_%v_staj_yazisi.pdf", t.Day(), int(t.Month()), t.Year())
}



func CreateDispatchFiles(siArr []si.StudentInternship){
	d := GenerateDispatchData(siArr)
	dispatchTemp := template.New(g.DispatchTemplateFileName)
	_, err := os.Stat(g.DispatchTemplateFilePath)
	if os.IsExist(err){
		_, err = dispatchTemp.ParseFiles(g.DispatchTemplateFilePath)
	}else{
		dispatchTemp.Parse(templates.DispatchFileContent)
	}

	DirExists(g.OutputDirPath, true)
	dispatchFODTFilePath := filepath.Join(g.OutputDirPath, DispatchFODTFileName())

	dispatchFile, err := os.Create(dispatchFODTFilePath)
	err = dispatchTemp.Execute(dispatchFile, d)
	dispatchFile.Close()

	cmd := exec.Command("libreoffice", "--headless", "--convert-to", "pdf", DispatchFODTFileName() , "output", DispatchPDFFileName())
	cmd.Dir = filepath.Join(filepath.Join(".", g.OutputDirPath))
	cmd.Run()
}


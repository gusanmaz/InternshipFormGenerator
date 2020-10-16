package main

import (
	"github.com/gusanmaz/InternshipFormGenerator/globals"
	"github.com/gusanmaz/InternshipFormGenerator/si"
	"github.com/gusanmaz/InternshipFormGenerator/sireader"
	"github.com/gusanmaz/InternshipFormGenerator/siwriter"
	"log"
	"os"
	"sort"
)

var g globals.Config

func init(){
	g = globals.Globals
}

//go:generate go run scripts/serializeTemplateFiles.go
func main(){
	siArr := sireader.ParseCSVFile(g.CSVFilePath, g.InstitutionName)
	if g.IntegrityCheck{
		integrity := sireader.ValidateIntegrity(siArr)
		if !integrity{
			log.Fatalln("CSV and JSON files do not match!")
			os.Exit(-1)
		}
	}
	sort.Sort(si.ByID(siArr))
	siwriter.CreateDispatchFiles(siArr)
	siwriter.CreateAllFormFiles(siArr)
}


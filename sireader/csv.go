package sireader

import (
	"encoding/csv"
	"github.com/gusanmaz/InternshipFormGenerator/date"
	"github.com/gusanmaz/InternshipFormGenerator/si"
	"os"
	"strconv"
	"strings"
)


func ParseCSVFile(path, institutionName string)[]si.StudentInternship {
	sFile, _ := os.Open(path)
	reader := csv.NewReader(sFile)
	studentRecords, _ := reader.ReadAll()


	siArr := make([]si.StudentInternship, 0)
	for _, record := range studentRecords {
		id := record[0]
		name := record[1]
		surname := record[2]
		subjectLoc := record[3]
		beginS := record[4]
		endS := record[5]
		daysRequestedS := record[6]
		daysAcceptedS := record[7]

		// Conversions
		beginD := date.ParseDate(beginS)
		endD   := date.ParseDate(endS)
		daysAcceptedI, _ := strconv.Atoi(daysAcceptedS)
		daysRequestedI, _ := strconv.Atoi(daysRequestedS)

		// TODO() Needs some refactoring

		oneAt := strings.HasPrefix(subjectLoc, "@")
		twoAt := strings.HasPrefix(subjectLoc, "@@")
		iType := 0
		subject := ""
		location := ""

		if twoAt{
			iType = Onsite
			subjectLoc := strings.TrimPrefix(subjectLoc, "@@")
			location = subjectLoc
			subject = ""
		}else if oneAt{
			iType = Onsite // was Online
			subjectLoc := strings.TrimPrefix(subjectLoc, "@")
			subject = subjectLoc
			location = "" // was institutionName
		}else{
			iType = OnlineVerifiable
			subject = subjectLoc
			location = institutionName
		}

		si := si.StudentInternship{
			Name:           name,
			Surname:        surname,
			ID:             id,

			Location:       location,
			Subject:        subject,
			InternshipType: iType,

			BeginDateS:     beginS,
			EndDateS:       endS,
			BeginDateD:     beginD,
			EndDateD:       endD,

			DaysRequestedI: daysRequestedI,
			DaysAcceptedI:  daysAcceptedI,
			DaysRequestedS: daysRequestedS,
			DaysAcceptedS:  daysAcceptedS,
		}

		siArr = append(siArr, si)
	}
	return siArr
}



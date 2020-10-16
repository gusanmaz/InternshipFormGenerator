package sireader

import (
	"encoding/json"
	"fmt"
	"github.com/gusanmaz/InternshipFormGenerator2/si"
	"io/ioutil"
	"net/http"
	"strconv"
	"strings"
)

type StudentJSON struct {
	StudentInfo struct {
		ID      int    `json:"id"`
		Name    string `json:"name"`
		Surname string `json:"surname"`
	} `json:"studentInfo"`
	ProjectInfo []struct {
		ID      int    `json:"id"`
		Subject string `json:"subject"`
		Days    int    `json:"days"`
	} `json:"projectInfo"`
}

func CleanSubject(subject string) string{
	s := strings.TrimPrefix(subject, " ")
	s = strings.TrimSuffix(s, " ")
	return s
}

func IsSubjectsSame(sub1, sub2 string) bool{
	return strings.Contains(sub1, sub2) || strings.Contains(sub2, sub1)
}

func GetStudentJSON(baseUrl string, studentID int) StudentJSON {
	url := fmt.Sprintf("%v/%v.json", baseUrl, studentID)
	resp, _ := http.Get(url)
	var s StudentJSON
	b, _ := ioutil.ReadAll(resp.Body)
	resp.Body.Close()
	json.Unmarshal(b, &s)
	return s
}

func ValidateIntegrity(siArr []si.StudentInternship) bool{
	corruptData := false
	for _, si := range siArr{
		if si.InternshipType != OnlineVerifiable {
			continue
		}

		id := si.ID
		idI, _ := strconv.Atoi(id)
		jsonData := GetStudentJSON(g.JSONTestBaseURL, idI)
		findMatch := false
		for _, hw := range jsonData.ProjectInfo{
			jsonSub := hw.Subject
			if IsSubjectsSame(jsonSub, si.Subject){
				findMatch = true
				jsonDays := hw.Days
				if jsonDays != si.DaysRequestedI{
					fmt.Println("Different day values for", hw, jsonData)
					corruptData = true
				}
				break
			}
		}
		
		if !findMatch{
			fmt.Println("Mismatch", jsonData, si)
			corruptData = true
		}
	}
	return !corruptData
}

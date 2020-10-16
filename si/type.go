package si

import "github.com/gusanmaz/InternshipFormGenerator/date"

const(
	Online = iota
	OnlineVerifiable
	Onsite
)

type StudentInternship struct {
	Name           string
	Surname        string
	ID             string
	InternshipType int
	Location       string
	Subject        string
	BeginDateS     string
	EndDateS       string
	BeginDateD     date.Date
	EndDateD       date.Date
	DaysRequestedI int
	DaysAcceptedI  int
	DaysRequestedS string
	DaysAcceptedS  string
}

type ByID []StudentInternship

func (si ByID) Len() int           { return len(si) }
func (si ByID) Less(i, j int) bool { return si[i].ID < si[j].ID }
func (si ByID) Swap(i, j int)      { si[i], si[j] = si[j], si[i] }

type ByName []StudentInternship

func (si ByName) Len() int           { return len(si) }
func (si ByName) Less(i, j int) bool { return si[i].Name < si[j].Name }
func (si ByName) Swap(i, j int)      { si[i], si[j] = si[j], si[i] }

type BySurname []StudentInternship

func (si BySurname) Len() int           { return len(si) }
func (si BySurname) Less(i, j int) bool { return si[i].Name < si[j].Name }
func (si BySurname) Swap(i, j int)      { si[i], si[j] = si[j], si[i] }




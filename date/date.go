package date

import (
	"log"
	"strconv"
	"strings"
)

type Date struct{
	Day int
	Month int
	Year int
}

func ParseDate(s string)Date{
	errorDate := Date{-1,-1,-1,}
	elements := []string{}
	seperators := []string{".", "-", "/", " "}
	for _, sep := range seperators{
		elements = strings.Split(s, sep)
		if len(elements) == 3{
			break
		}
	}

	if len(elements) != 3{
		log.Fatalln("Wrong File Format")
		return errorDate
	}

	for i, v := range elements{
		elements[i] = strings.ReplaceAll(v, " ", "")
	}

	//elements = strings.Split(s, ".")
	di, errD := strconv.Atoi(elements[0])
	mi, errM := strconv.Atoi(elements[1])
	yi, errY := strconv.Atoi(elements[2])

	if errD == nil && errM == nil && errY == nil{
		return Date{Day:di, Month:mi, Year:yi}
	}

	log.Fatalln("Parse resulted in non-integer date values.")
	return errorDate
}

func (d1 Date) IsEqual(d2 Date)bool{
	return (d1.Day == d2.Day) && (d1.Month == d2.Month) && (d1.Year == d2.Year)
}

func (d1 Date) IsBefore(d2 Date) bool{
	if d1.IsEqual(d2){
		return false
	}

	if d1.Year > d2.Year{
		return false
	}
	if d1.Year < d2.Year{
		return true
	}

	if d1.Month > d2.Month{
		return false
	}
	if d1.Month < d2.Month{
		return true
	}

	if d1.Day > d2.Day{
		return false
	}
	if d1.Day < d2.Day{
		return true
	}

	return false
}

func (d1 Date) IsAfter(d2 Date) bool{
	if d1.IsEqual(d2){
		return false
	}
	return d2.IsBefore(d1)
}

func (d Date) Between(begin, end Date)bool{
	afterBegin := d.IsAfter(begin) || d.IsEqual(begin)
	beforeEnd  := d.IsBefore(end) || d.IsEqual(end)

	return afterBegin && beforeEnd
}

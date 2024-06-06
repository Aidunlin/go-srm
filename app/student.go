package app

import (
	"fmt"
	"net/url"
	"strconv"
)

type StudentRecord struct {
	IdRaw           string
	GpaRaw          string
	FinancialAidRaw string

	Id             int64
	FirstName      string
	LastName       string
	Email          string
	Phone          string
	DegreeProgram  string
	Gpa            float64
	FinancialAid   int64
	GraduationDate string
}

func NewStudentFromCreateForm(params url.Values) (StudentRecord, []string) {
	data := StudentRecord{}
	errors := []string{}

	firstName := params.Get("first_name")
	if len(firstName) == 0 {
		errors = append(errors, "A <strong>first name</strong> is required.")
	} else {
		data.FirstName = firstName
	}

	lastName := params.Get("last_name")
	if len(lastName) == 0 {
		errors = append(errors, "A <strong>last name</strong> is required.")
	} else {
		data.LastName = lastName
	}

	data.GpaRaw = params.Get("gpa")
	gpa, err := strconv.ParseFloat(data.GpaRaw, 64)
	if err == nil {
		data.Gpa = gpa
	}

	degreeProgram := params.Get("degree_program")
	if isDegreeProgram(degreeProgram) {
		data.DegreeProgram = degreeProgram
	} else {
		errors = append(errors, "Invalid <strong>degree program</strong>.")
	}

	graudationDate := params.Get("graduation_date")
	if isGraduationDate(graudationDate) {
		data.GraduationDate = graudationDate
	} else {
		errors = append(errors, "Invalid <strong>graduation date</strong>.")
	}

	data.FinancialAidRaw = params.Get("financial_aid")
	financialAid, err := strconv.Atoi(data.FinancialAidRaw)
	if err != nil || !isFinancialAid(financialAid) {
		errors = append(errors, "An option for <strong>financial aid</strong> is required.")
	} else {
		data.FinancialAid = int64(financialAid)
	}

	email := params.Get("email")
	if len(email) > 0 {
		data.Email = email
	} else {
		errors = append(errors, "An <strong>email address</strong> is required.")
	}

	phone := params.Get("phone")
	if len(phone) > 0 {
		data.Phone = phone
	} else {
		errors = append(errors, "A <strong>phone number</strong> is required.")
	}

	return data, errors
}

func NewStudentFromUpdateForm(params url.Values) (StudentRecord, []string) {
	return NewStudentFromCreateForm(params)
}

func NewStudentFromAdvancedSearchForm(params url.Values) StudentRecord {
	data := StudentRecord{}

	firstName := params.Get("first_name")
	if len(firstName) > 0 {
		data.FirstName = firstName
	}

	lastName := params.Get("last_name")
	if len(lastName) > 0 {
		data.LastName = lastName
	}

	data.GpaRaw = params.Get("gpa")
	gpa, err := strconv.ParseFloat(data.GpaRaw, 64)
	if err == nil {
		data.Gpa = gpa
	}

	degreeProgram := params.Get("degree_program")
	if isDegreeProgram(degreeProgram) {
		data.DegreeProgram = degreeProgram
	}

	graudationDate := params.Get("graduation_date")
	if isGraduationDate(graudationDate) {
		data.GraduationDate = graudationDate
	}

	data.FinancialAidRaw = params.Get("financial_aid")
	financialAid, err := strconv.Atoi(data.FinancialAidRaw)
	if err == nil || isFinancialAid(financialAid) {
		data.FinancialAid = int64(financialAid)
	}

	email := params.Get("email")
	if len(email) > 0 {
		data.Email = email
	}

	phone := params.Get("phone")
	if len(phone) > 0 {
		data.Phone = phone
	}

	return data
}

func (p StudentRecord) GetRawMap() map[string]string {
	return map[string]string{
		"id":              p.IdRaw,
		"first_name":      p.FirstName,
		"last_name":       p.LastName,
		"email":           p.Email,
		"phone":           p.Phone,
		"degree_program":  p.DegreeProgram,
		"gpa":             p.GpaRaw,
		"financial_aid":   p.FinancialAidRaw,
		"graduation_date": p.GraduationDate,
	}
}

func (p StudentRecord) GetMap() map[string]string {
	return map[string]string{
		"id":              fmt.Sprint(p.Id),
		"first_name":      p.FirstName,
		"last_name":       p.LastName,
		"email":           p.Email,
		"phone":           p.Phone,
		"degree_program":  p.DegreeProgram,
		"gpa":             fmt.Sprint(p.Gpa),
		"financial_aid":   fmt.Sprint(p.FinancialAid),
		"graduation_date": p.GraduationDate,
	}
}

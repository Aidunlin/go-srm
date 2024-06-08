package model

import (
	"fmt"
	"net/url"
	"strconv"

	"github.com/Aidunlin/go-srm/value"
	"github.com/go-mysql-org/go-mysql/mysql"
)

type StudentRecord struct {
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

// Excludes the record id.
func (s StudentRecord) ToMap() map[string]string {
	return map[string]string{
		"first_name":      s.FirstName,
		"last_name":       s.LastName,
		"email":           s.Email,
		"phone":           s.Phone,
		"degree_program":  s.DegreeProgram,
		"gpa":             fmt.Sprint(s.Gpa),
		"financial_aid":   fmt.Sprint(s.FinancialAid),
		"graduation_date": s.GraduationDate,
	}
}

func NewStudentRecord(params url.Values) (StudentRecord, []string) {
	form := StudentRecord{}
	errors := []string{}

	firstName := params.Get("first_name")
	if len(firstName) == 0 {
		errors = append(errors, "A <strong>first name</strong> is required.")
	} else {
		form.FirstName = firstName
	}

	lastName := params.Get("last_name")
	if len(lastName) == 0 {
		errors = append(errors, "A <strong>last name</strong> is required.")
	} else {
		form.LastName = lastName
	}

	gpa, err := strconv.ParseFloat(params.Get("gpa"), 64)
	if err == nil {
		form.Gpa = gpa
	}

	degreeProgram := params.Get("degree_program")
	if value.IsDegreeProgram(degreeProgram) {
		form.DegreeProgram = degreeProgram
	} else {
		errors = append(errors, "Invalid <strong>degree program</strong>.")
	}

	graudationDate := params.Get("graduation_date")
	if value.IsGraduationDate(graudationDate) {
		form.GraduationDate = graudationDate
	} else {
		errors = append(errors, "Invalid <strong>graduation date</strong>.")
	}

	financialAid, err := strconv.Atoi(params.Get("financial_aid"))
	if err != nil || !value.IsFinancialAid(financialAid) {
		errors = append(errors, "An option for <strong>financial aid</strong> is required.")
	} else {
		form.FinancialAid = int64(financialAid)
	}

	email := params.Get("email")
	if len(email) > 0 {
		form.Email = email
	} else {
		errors = append(errors, "An <strong>email address</strong> is required.")
	}

	phone := params.Get("phone")
	if len(phone) > 0 {
		form.Phone = phone
	} else {
		errors = append(errors, "A <strong>phone number</strong> is required.")
	}

	return form, errors
}

func GetStudentRecordFromResult(result *mysql.Result, row int) StudentRecord {
	id, _ := result.GetIntByName(row, "id")
	firstName, _ := result.GetStringByName(row, "first_name")
	lastName, _ := result.GetStringByName(row, "last_name")
	email, _ := result.GetStringByName(row, "email")
	phone, _ := result.GetStringByName(row, "phone")
	degreeProgram, _ := result.GetStringByName(row, "degree_program")
	gpa, _ := result.GetFloatByName(row, "gpa")
	financialAid, _ := result.GetIntByName(row, "financial_aid")
	graduationDate, _ := result.GetStringByName(row, "graduation_date")

	return StudentRecord{
		Id:             id,
		FirstName:      firstName,
		LastName:       lastName,
		Email:          email,
		Phone:          phone,
		DegreeProgram:  degreeProgram,
		Gpa:            gpa,
		FinancialAid:   financialAid,
		GraduationDate: graduationDate,
	}
}

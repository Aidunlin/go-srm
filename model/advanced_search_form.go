package model

import (
	"context"
	"fmt"
	"net/url"
	"strconv"

	"github.com/Aidunlin/go-srm/value"
)

type AdvancedSearchForm struct {
	Searched       bool
	FirstName      string
	LastName       string
	Email          string
	Phone          string
	DegreeProgram  string
	Gpa            float64
	FinancialAid   int64
	GraduationDate string
}

func (s AdvancedSearchForm) ToMap() map[string]string {
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

func NewAdvancedSearchForm(params url.Values) AdvancedSearchForm {
	form := AdvancedSearchForm{}

	firstName := params.Get("first_name")
	if len(firstName) > 0 {
		form.FirstName = firstName
		form.Searched = true
	}

	lastName := params.Get("last_name")
	if len(lastName) > 0 {
		form.LastName = lastName
		form.Searched = true
	}

	gpa := params.Get("gpa")
	if len(gpa) > 0 {
		gpa, err := strconv.ParseFloat(params.Get("gpa"), 64)
		if err == nil {
			form.Gpa = gpa
			form.Searched = true
		}
	}

	degreeProgram := params.Get("degree_program")
	if value.IsDegreeProgram(degreeProgram) {
		form.DegreeProgram = degreeProgram
		form.Searched = true
	}

	graudationDate := params.Get("graduation_date")
	if value.IsGraduationDate(graudationDate) {
		form.GraduationDate = graudationDate
		form.Searched = true
	}

	aidString := params.Get("financial_aid")
	if len(aidString) > 0 {
		financialAid, err := strconv.Atoi(aidString)
		if err == nil && value.IsFinancialAid(financialAid) {
			form.FinancialAid = int64(financialAid)
			form.Searched = true
		}
	}

	email := params.Get("email")
	if len(email) > 0 {
		form.Email = email
		form.Searched = true
	}

	phone := params.Get("phone")
	if len(phone) > 0 {
		form.Phone = phone
		form.Searched = true
	}

	return form
}

func GetAdvancedSearchForm(ctx context.Context) AdvancedSearchForm {
	if params, ok := ctx.Value("form").(AdvancedSearchForm); ok {
		return params
	}
	return NewAdvancedSearchForm(nil)
}

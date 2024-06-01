package app

import (
	"context"
	"fmt"
	"net/url"
	"slices"
	"strconv"
	"time"
)

func DisplayDate(value string) string {
	date, err := time.Parse(time.DateOnly, value)
	if err != nil {
		return ""
	}
	return fmt.Sprintf("%02d/%02d/%04d", date.Month(), date.Day(), date.Year())
}

type StudentRecordColumn struct {
	Name                string
	Label               string
	title               string
	BasicSearch         bool
	BasicSearchExact    bool
	AdvancedSearchExact bool
}

func (c StudentRecordColumn) Title() string {
	if len(c.title) > 0 {
		return c.title
	}
	return c.Label
}

func GetColumns() []StudentRecordColumn {
	return []StudentRecordColumn{
		{
			Name:                "student_id",
			Label:               "SID",
			title:               "Student ID",
			BasicSearch:         true,
			BasicSearchExact:    true,
			AdvancedSearchExact: true,
		}, {
			Name:        "first_name",
			Label:       "First Name",
			BasicSearch: true,
		}, {
			Name:        "last_name",
			Label:       "Last Name",
			BasicSearch: true,
		}, {
			Name:                "gpa",
			Label:               "GPA",
			AdvancedSearchExact: true,
		}, {
			Name:        "degree_program",
			Label:       "Degree Program",
			BasicSearch: true,
		}, {
			Name:                "graduation_date",
			Label:               "Graduation",
			title:               "Graduation Date",
			AdvancedSearchExact: true,
		}, {
			Name:                "financial_aid",
			Label:               "Aid",
			title:               "Financial Aid",
			AdvancedSearchExact: true,
		}, {
			Name:                "email",
			Label:               "Email",
			BasicSearch:         true,
			BasicSearchExact:    true,
			AdvancedSearchExact: true,
		}, {
			Name:                "phone",
			Label:               "Phone",
			BasicSearch:         true,
			BasicSearchExact:    true,
			AdvancedSearchExact: true,
		},
	}
}

func GetColumn(name string) StudentRecordColumn {
	for _, column := range GetColumns() {
		if column.Name == name {
			return column
		}
	}
	return StudentRecordColumn{}
}

func GetDegrees() []string {
	return []string{
		"Advanced Manufacturing",
		"Business Administration",
		"Cuisine Management",
		"Cybersecurity",
		"Digital Media Arts",
		"Fine Arts",
		"Management",
		"Marketing",
		"Music",
		"Network Technology",
		"Professional Baking and Pastry Arts",
		"Web Development",
	}
}

func getAlphabet() []string {
	alpha := []string{}
	for char := 'a'; char <= 'z'; char++ {
		alpha = append(alpha, string(char))
	}
	return alpha
}

func isFilter(value string) bool {
	return slices.Contains(getAlphabet(), value)
}

func isSortBy(value string) bool {
	for _, column := range GetColumns() {
		if value == column.Name {
			return true
		}
	}
	return false
}

func isOrdering(value string) bool {
	return value == "asc" || value == "desc"
}

func isDegreeProgram(value string) bool {
	if value == "" {
		return true
	}

	for _, degree := range GetDegrees() {
		if value == degree {
			return true
		}
	}

	return false
}

func isGraduationDate(value string) bool {
	if value == "" {
		return true
	}

	_, err := time.Parse(time.DateOnly, value)
	return err == nil
}

func isFinancialAid(value int) bool {
	return value == 0 || value == 1
}

type RecordTableParams struct {
	Filter string
	Sortby string
	Order  string
	Page   int
	Search string
}

func NewRecordTableParams(params url.Values) RecordTableParams {
	filter := params.Get("filter")
	if !isFilter(filter) {
		filter = ""
	}

	sortby := params.Get("sortby")
	if !isSortBy(sortby) {
		sortby = "last_name"
	}

	order := params.Get("order")
	if !isOrdering(order) {
		order = "asc"
	}

	page, err := strconv.Atoi(params.Get("page"))
	if err != nil {
		page = 1
	}
	page = max(page, 1)

	search := params.Get("q")

	return RecordTableParams{
		Filter: filter,
		Sortby: sortby,
		Order:  order,
		Page:   page,
		Search: search,
	}
}

func AdvancedSearchParams(input url.Values) map[string]string {
	output := map[string]string{}

	for _, column := range GetColumns() {
		if input.Has(column.Name) && len(input.Get(column.Name)) > 0 {
			output[column.Name] = input.Get(column.Name)
		}
	}

	return output
}

type StudentRecord struct {
	IdRaw           string
	StudentIdRaw    string
	GpaRaw          string
	FinancialAidRaw string

	Id             int64
	StudentId      int64
	FirstName      string
	LastName       string
	Email          string
	Phone          string
	DegreeProgram  string
	Gpa            float64
	FinancialAid   int64
	GraduationDate string
}

func NewStudentRecord(params url.Values) (StudentRecord, []string) {
	data := StudentRecord{}
	errors := []string{}

	data.StudentIdRaw = params.Get("student_id")
	studentId, err := strconv.Atoi(data.StudentIdRaw)
	if err != nil || studentId == 0 {
		errors = append(errors, "A <strong>student ID</strong> is required.")
	} else {
		data.StudentId = int64(studentId)
	}

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

func NewAdvancedSearchForm(params url.Values) StudentRecord {
	data := StudentRecord{}

	data.StudentIdRaw = params.Get("student_id")
	studentId, err := strconv.Atoi(data.StudentIdRaw)
	if err == nil && studentId != 0 {
		data.StudentId = int64(studentId)
	}

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
		"student_id":      p.StudentIdRaw,
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
		"student_id":      fmt.Sprint(p.StudentId),
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

type MessageParams struct {
	Success string
	Error   string
}

func NewMessageParams(params url.Values) MessageParams {
	return MessageParams{
		Success: params.Get("success"),
		Error:   params.Get("error"),
	}
}

func GetTableParams(ctx context.Context) RecordTableParams {
	if params, ok := ctx.Value("table").(RecordTableParams); ok {
		return params
	}
	return NewRecordTableParams(nil)
}

func GetMessageParams(ctx context.Context) MessageParams {
	if params, ok := ctx.Value("message").(MessageParams); ok {
		return params
	}
	return NewMessageParams(nil)
}

func GetFormParams(ctx context.Context) map[string]string {
	if params, ok := ctx.Value("form").(map[string]string); ok {
		return params
	}
	return nil
}

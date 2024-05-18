package app

import (
	"fmt"
	"math"
	"net/url"
	"slices"
	"strconv"
	"time"

	"github.com/a-h/templ"
)

const AppCopyright = "&copy; 2024 Aidan Linerud"
const AppVersion = 3.0
const AppName = "CTEC 127 Record Manager"
const AppStatus = "Development"
const PaginateLimit = 10

type ParamMap = map[string]string

func DisplayDate(value string) string {
	date, err := time.Parse(time.DateOnly, value)
	if err != nil {
		return ""
	}
	return fmt.Sprintf("%02d/%02d/%04d", date.Month(), date.Day(), date.Year())
}

func GetTotalPages(totalResults int64) int {
	return int(math.Ceil(float64(totalResults) / float64(PaginateLimit)))
}

type StudentRecordColumn struct {
	Name  string
	Label string
	Title string
}

func GetColumns() []StudentRecordColumn {
	return []StudentRecordColumn{
		{Name: "student_id", Label: "SID", Title: "Student Id"},
		{Name: "first_name", Label: "First Name", Title: "First Name"},
		{Name: "last_name", Label: "Last Name", Title: "Last Name"},
		{Name: "gpa", Label: "GPA", Title: "GPA"},
		{Name: "degree_program", Label: "Degree Program", Title: "Degree Program"},
		{Name: "graduation_date", Label: "Graduation", Title: "Graduation Date"},
		{Name: "financial_aid", Label: "Aid", Title: "Financial Aid"},
		{Name: "email", Label: "Email", Title: "Email"},
		{Name: "phone", Label: "Phone", Title: "Phone"},
	}
}

func GetColumnLabel(name string) string {
	for _, column := range GetColumns() {
		if column.Name == name {
			return column.Label
		}
	}
	return ""
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

	search := params.Get("search")

	return RecordTableParams{
		Filter: filter,
		Sortby: sortby,
		Order:  order,
		Page:   page,
		Search: search,
	}
}

func (p RecordTableParams) QueryString(with ParamMap, without ...string) templ.SafeURL {
	v := url.Values{}
	if len(p.Filter) > 0 {
		v.Set("filter", p.Filter)
	}
	v.Set("sortby", p.Sortby)
	v.Set("order", p.Order)
	v.Set("page", fmt.Sprint(p.Page))
	if len(p.Search) > 0 {
		v.Set("search", p.Search)
	}
	for key, value := range with {
		v.Set(key, value)
	}
	for _, value := range without {
		v.Del(value)
	}
	return templ.URL(fmt.Sprintf("?%v", v.Encode()))
}

type StudentRecord struct {
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

type RecordFormParams struct {
	StudentId      int
	FirstName      string
	LastName       string
	Gpa            float64
	DegreeProgram  string
	GraduationDate string
	FinancialAid   int64
	Email          string
	Phone          string
	Id             int
}

func NewRecordFormParams(params url.Values, requireId bool) (RecordFormParams, []string) {
	data := RecordFormParams{}
	errors := []string{}

	if requireId {
		id, err := strconv.Atoi(params.Get("id"))
		if err != nil {
			errors = append(errors, "Missing or invalid record id.")
		} else {
			data.Id = id
		}
	}

	studentId, err := strconv.Atoi(params.Get("student_id"))
	if err != nil || studentId == 0 {
		errors = append(errors, "A <strong>student ID</strong> is required.")
	} else {
		data.StudentId = studentId
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

	gpa, err := strconv.ParseFloat(params.Get("gpa"), 64)
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

	financialAid, err := strconv.Atoi(params.Get("financial_aid"))
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

func (p RecordFormParams) GetMap(withId bool) ParamMap {
	paramMap := ParamMap{
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
	if withId {
		paramMap["id"] = fmt.Sprint(p.Id)
	}
	return paramMap
}

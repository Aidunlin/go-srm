package main

import (
	"fmt"
	"net/http"
	"net/url"
	"slices"
	"strconv"
	"time"

	"github.com/a-h/templ"
	"github.com/go-mysql-org/go-mysql/client"
	"github.com/labstack/echo"
)

type ParamMap = map[string]string

func getDateString(value string) string {
	date, err := time.Parse(time.DateOnly, value)
	if err != nil {
		return ""
	}
	return fmt.Sprintf("%02d/%02d/%04d", date.Month(), date.Day(), date.Year())
}

const AppCopyright = "&copy; 2024 Aidan Linerud"
const AppVersion = 3.0
const AppName = "CTEC 127 Record Manager"
const DBTable = "student_v2"
const AppStatus = "Development"
const PaginateLimit = 10

type StudentRecordColumn struct {
	Name  string
	Label string
	Title string
}

func getColumns() []StudentRecordColumn {
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

func getColumnLabel(name string) string {
	for _, column := range getColumns() {
		if column.Name == name {
			return column.Label
		}
	}
	return ""
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

func isSortby(value string) bool {
	for _, column := range getColumns() {
		if value == column.Name {
			return true
		}
	}
	return false
}

func isOrdering(value string) bool {
	return value == "asc" || value == "desc"
}

type RecordTableParams struct {
	Filter string
	Sortby string
	Order  string
	Page   int
	Search string
}

func newRecordTableParams(params url.Values) RecordTableParams {
	filter := params.Get("filter")
	if !isFilter(filter) {
		filter = ""
	}

	sortby := params.Get("sortby")
	if !isSortby(sortby) {
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

func selectRecords(params RecordTableParams) (int64, []StudentRecord, error) {
	db, err := client.Connect("localhost:3306", "root", "", "ctec")
	if err != nil {
		return 0, nil, err
	}

	whereSql := ""

	if len(params.Filter) > 0 {
		whereSql = fmt.Sprintf("WHERE last_name LIKE '%v%%'", params.Filter)
	}

	totalSql := fmt.Sprintf("SELECT COUNT(*) AS total from %v %v", DBTable, whereSql)
	totalResult, err := db.Execute(totalSql)
	if err != nil {
		return 0, nil, err
	}
	total, err := totalResult.GetIntByName(0, "total")
	if err != nil {
		return 0, nil, err
	}

	orderSql := fmt.Sprintf("ORDER BY %v %v", params.Sortby, params.Order)

	offset := (params.Page - 1) * PaginateLimit
	pageSql := fmt.Sprintf("LIMIT %v OFFSET %v", PaginateLimit, offset)

	recordsSql := fmt.Sprintf("SELECT * FROM %v %v %v %v", DBTable, whereSql, orderSql, pageSql)
	recordsResult, err := db.Execute(recordsSql)
	if err != nil {
		return 0, nil, err
	}

	records := []StudentRecord{}

	for i := range recordsResult.Values {
		id, _ := recordsResult.GetIntByName(i, "id")
		studentId, _ := recordsResult.GetIntByName(i, "student_id")
		firstName, _ := recordsResult.GetStringByName(i, "first_name")
		lastName, _ := recordsResult.GetStringByName(i, "last_name")
		email, _ := recordsResult.GetStringByName(i, "email")
		phone, _ := recordsResult.GetStringByName(i, "phone")
		degreeProgram, _ := recordsResult.GetStringByName(i, "degree_program")
		gpa, _ := recordsResult.GetFloatByName(i, "gpa")
		financialAid, _ := recordsResult.GetIntByName(i, "financial_aid")
		graduationDate, _ := recordsResult.GetStringByName(i, "graduation_date")

		records = append(records, StudentRecord{
			Id:             id,
			StudentId:      studentId,
			FirstName:      firstName,
			LastName:       lastName,
			Email:          email,
			Phone:          phone,
			DegreeProgram:  degreeProgram,
			Gpa:            gpa,
			FinancialAid:   financialAid,
			GraduationDate: graduationDate,
		})
	}
	return total, records, nil
}

// Renders a templ component using echo.
func render(c echo.Context, code int, t templ.Component) error {
	buf := templ.GetBuffer()
	defer templ.ReleaseBuffer(buf)

	if err := t.Render(c.Request().Context(), buf); err != nil {
		return err
	}

	return c.HTML(code, buf.String())
}

// Entry point for the web server.
func main() {
	e := echo.New()
	e.Static("/css", "css")
	e.Static("/js", "js")

	e.GET("/", func(c echo.Context) error {
		params := newRecordTableParams(c.QueryParams())
		total, records, err := selectRecords(params)
		if err != nil {
			return c.String(http.StatusInternalServerError, fmt.Sprint(err))
		}
		return render(c, http.StatusOK, page(indexPage(total, records, params), c.Path()))
	})

	e.Logger.Fatal((e.Start(":3000")))
}
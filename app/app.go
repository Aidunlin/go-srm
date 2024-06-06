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
	for _, column := range GetStudentColumns() {
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

	for _, column := range GetStudentColumns() {
		if input.Has(column.Name) && len(input.Get(column.Name)) > 0 {
			output[column.Name] = input.Get(column.Name)
		}
	}

	return output
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

package model

import (
	"context"
	"net/url"
	"slices"
	"strconv"
)

type StudentTableParams struct {
	Filter string
	Sortby string
	Order  string
	Page   int
	Search string
}

func NewStudentTableParams(params url.Values) StudentTableParams {
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

	return StudentTableParams{
		Filter: filter,
		Sortby: sortby,
		Order:  order,
		Page:   page,
		Search: search,
	}
}

func GetStudentTableParams(ctx context.Context) StudentTableParams {
	if params, ok := ctx.Value("table").(StudentTableParams); ok {
		return params
	}
	return NewStudentTableParams(nil)
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

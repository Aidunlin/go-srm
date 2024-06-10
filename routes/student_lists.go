package routes

import (
	"net/http"

	"github.com/Aidunlin/go-srm/db"
	"github.com/Aidunlin/go-srm/model"
	"github.com/Aidunlin/go-srm/templates"
	"github.com/labstack/echo/v4"
)

func AddMainRoutes(e *echo.Echo) {
	e.GET("/", getIndex)
	e.GET("/search", getSearch)
	e.GET("/advanced-search", getAdvancedSearch)
}

func getIndex(c echo.Context) error {
	queryParams := c.QueryParams()
	tableParams := model.NewStudentTableParams(queryParams)
	total, students := db.SelectStudents(tableParams)
	params := map[any]any{
		"table":   tableParams,
		"message": model.NewMessageParams(queryParams),
	}
	return render(c, http.StatusOK, templates.IndexPage(total, students, isAuthenticated(c)), params)
}

func getSearch(c echo.Context) error {
	queryParams := c.QueryParams()
	tableParams := model.NewStudentTableParams(queryParams)
	total, students := db.SelectStudents(tableParams)
	params := map[any]any{
		"table": tableParams,
	}
	return render(c, http.StatusOK, templates.SearchPage(total, students, isAuthenticated(c)), params)
}

func getAdvancedSearch(c echo.Context) error {
	formParams, _ := c.FormParams()
	form := model.NewAdvancedSearchForm(formParams)
	if !form.Searched {
		return render(c, http.StatusOK, templates.AdvancedSearchPage(model.AdvancedSearchForm{}, 0, nil, false, isAuthenticated(c)), nil)
	}
	queryParams := c.QueryParams()
	tableParams := model.NewStudentTableParams(queryParams)
	total, students := db.SelectStudentsWithForm(tableParams, form)
	params := map[any]any{
		"table": tableParams,
		"form":  form,
	}
	return render(c, http.StatusOK, templates.AdvancedSearchPage(form, total, students, true, isAuthenticated(c)), params)
}

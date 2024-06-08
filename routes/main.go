package routes

import (
	"context"
	"net/http"

	"github.com/Aidunlin/go-srm/db"
	"github.com/Aidunlin/go-srm/model"
	"github.com/Aidunlin/go-srm/templates"
	"github.com/a-h/templ"
	"github.com/labstack/echo/v4"
)

func AddMainRoutes(e *echo.Echo) {
	e.GET("/", getIndex)
	e.GET("/search", getSearch)
	e.GET("/advanced-search", getAdvancedSearch)
}

// Renders a templ component using echo.
func render(c echo.Context, code int, t templ.Component, params map[any]any) error {
	buf := templ.GetBuffer()
	defer templ.ReleaseBuffer(buf)

	ctx := c.Request().Context()

	for key, value := range params {
		ctx = context.WithValue(ctx, key, value)
	}

	if err := t.Render(ctx, buf); err != nil {
		return err
	}

	return c.HTML(code, buf.String())
}

func getIndex(c echo.Context) error {
	queryParams := c.QueryParams()
	tableParams := model.NewStudentTableParams(queryParams)
	total, students := db.SelectStudents(tableParams)
	params := map[any]any{
		"table":   tableParams,
		"message": model.NewMessageParams(queryParams),
	}
	return render(c, http.StatusOK, templates.IndexPage(total, students), params)
}

func getSearch(c echo.Context) error {
	queryParams := c.QueryParams()
	tableParams := model.NewStudentTableParams(queryParams)
	total, students := db.SelectStudents(tableParams)
	params := map[any]any{
		"table": tableParams,
	}
	return render(c, http.StatusOK, templates.SearchPage(total, students), params)
}

func getAdvancedSearch(c echo.Context) error {
	formParams, _ := c.FormParams()
	form := model.NewAdvancedSearchForm(formParams)
	if !form.Searched {
		return render(c, http.StatusOK, templates.AdvancedSearchPage(model.AdvancedSearchForm{}, 0, nil, false), nil)
	}
	queryParams := c.QueryParams()
	tableParams := model.NewStudentTableParams(queryParams)
	total, students := db.SelectStudentsWithForm(tableParams, form)
	params := map[any]any{
		"table": tableParams,
		"form":  form,
	}
	return render(c, http.StatusOK, templates.AdvancedSearchPage(form, total, students, true), params)
}

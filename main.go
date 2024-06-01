package main

import (
	"context"
	"fmt"
	"net/http"
	"strconv"

	"github.com/Aidunlin/go-srm/app"
	"github.com/Aidunlin/go-srm/db"
	"github.com/Aidunlin/go-srm/templates"
	"github.com/a-h/templ"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

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

// Entry point for the web server.
func main() {
	e := echo.New()
	e.Use(middleware.Logger())
	e.Static("/css", "css")

	e.GET("/", func(c echo.Context) error {
		queryParams := c.QueryParams()
		tableParams := app.NewRecordTableParams(queryParams)
		total, records := db.SelectRecords(tableParams)
		params := map[any]any{
			"table":   tableParams,
			"message": app.NewMessageParams(queryParams),
		}
		return render(c, http.StatusOK, templates.IndexPage(total, records), params)
	})

	e.GET("/search", func(c echo.Context) error {
		queryParams := c.QueryParams()
		tableParams := app.NewRecordTableParams(queryParams)
		total, records := db.SelectRecords(tableParams)
		params := map[any]any{
			"table": tableParams,
		}
		return render(c, http.StatusOK, templates.SearchPage(total, records), params)
	})

	e.GET("/advanced-search", func(c echo.Context) error {
		formParams, _ := c.FormParams()
		form := app.AdvancedSearchParams(formParams)
		if len(form) == 0 {
			return render(c, http.StatusOK, templates.AdvancedSearchPage(app.StudentRecord{}, 0, nil, false), nil)
		}
		record := app.NewAdvancedSearchForm(formParams)
		queryParams := c.QueryParams()
		tableParams := app.NewRecordTableParams(queryParams)
		total, records := db.SelectRecordsWithForm(tableParams, record)
		params := map[any]any{
			"table": tableParams,
			"form":  form,
		}
		return render(c, http.StatusOK, templates.AdvancedSearchPage(record, total, records, true), params)
	})

	e.GET("/create", func(c echo.Context) error {
		return render(c, http.StatusOK, templates.CreatePage(app.StudentRecord{}, []string{}), nil)
	})

	e.POST("/create", func(c echo.Context) error {
		formParams, _ := c.FormParams()
		record, errors := app.NewStudentRecord(formParams)
		if len(errors) > 0 {
			return render(c, http.StatusOK, templates.CreatePage(record, errors), nil)
		}
		success := db.InsertRecord(record)
		if !success {
			return render(c, http.StatusOK, templates.CreatePage(record, []string{"Could not save that record!"}), nil)
		}
		return c.Redirect(http.StatusSeeOther, fmt.Sprintf("/?success=%v created!", record.FirstName))
	})

	e.GET("/update/:id", func(c echo.Context) error {
		id, err := strconv.Atoi(c.Param("id"))
		if err != nil {
			return c.Redirect(http.StatusSeeOther, "/?error=Invalid id.")
		}
		success, record := db.SelectRecord(id)
		if !success {
			return c.Redirect(http.StatusSeeOther, "/?error=Could not get record!")
		}
		return render(c, http.StatusOK, templates.UpdatePage(record, []string{}), nil)
	})

	e.POST("/update/:id", func(c echo.Context) error {
		id, err := strconv.Atoi(c.Param("id"))
		if err != nil {
			return c.Redirect(http.StatusSeeOther, "/?error=Invalid id.")
		}
		formParams, _ := c.FormParams()
		record, errors := app.NewStudentRecord(formParams)
		if len(errors) > 0 {
			return render(c, http.StatusOK, templates.UpdatePage(record, errors), nil)
		}
		success := db.UpdateRecord(id, record)
		if !success {
			return render(c, http.StatusOK, templates.UpdatePage(record, []string{"Could not update that record!"}), nil)
		}
		return c.Redirect(http.StatusSeeOther, fmt.Sprintf("/?success=%v updated!", record.FirstName))

	})

	e.GET("/delete/:id", func(c echo.Context) error {
		id, err := strconv.Atoi(c.Param("id"))
		if err != nil {
			return c.Redirect(http.StatusSeeOther, "/?error=Invalid id.")
		}
		success := db.DeleteRecord(id)
		if !success {
			return c.Redirect(http.StatusSeeOther, "/?error=Could not delete record!")
		}
		return c.Redirect(http.StatusSeeOther, "/?success=Deleted record.")
	})

	e.Logger.Fatal((e.Start(":3000")))
}

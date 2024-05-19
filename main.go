package main

import (
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
	e.Use(middleware.Logger())
	e.Static("/css", "css")
	e.Static("/js", "js")

	e.GET("/", func(c echo.Context) error {
		queryParams := c.QueryParams()
		tableParams := app.NewRecordTableParams(queryParams)
		total, records := db.SelectRecords(tableParams)
		messageParams := app.NewMessageParams(queryParams)
		return render(c, http.StatusOK, templates.IndexPage(total, records, tableParams, messageParams))
	})

	e.GET("/create", func(c echo.Context) error {
		return render(c, http.StatusOK, templates.CreatePage(app.StudentRecord{}, []string{}))
	})

	e.POST("/create", func(c echo.Context) error {
		formParams, _ := c.FormParams()
		record, errors := app.NewStudentRecord(formParams)
		if len(errors) > 0 {
			return render(c, http.StatusOK, templates.CreatePage(record, errors))
		}
		success := db.InsertRecord(record)
		if success {
			return c.Redirect(http.StatusSeeOther, fmt.Sprintf("/?success=%v created!", record.FirstName))
		} else {
			return render(c, http.StatusOK, templates.CreatePage(record, []string{"Could not save that record!"}))
		}
	})

	e.GET("/update/:id", func(c echo.Context) error {
		id, err := strconv.Atoi(c.Param("id"))
		if err != nil {
			return c.Redirect(http.StatusSeeOther, "/?error=Invalid id.")
		}
		success, record := db.SelectRecord(id)
		if success {
			return render(c, http.StatusOK, templates.UpdatePage(record, []string{}))
		} else {
			return c.Redirect(http.StatusSeeOther, "/?error=Could not get record!")
		}
	})

	e.POST("/update/:id", func(c echo.Context) error {
		id, err := strconv.Atoi(c.Param("id"))
		if err != nil {
			return c.Redirect(http.StatusSeeOther, "/?error=Invalid id.")
		}
		formParams, _ := c.FormParams()
		record, errors := app.NewStudentRecord(formParams)
		if len(errors) > 0 {
			return render(c, http.StatusOK, templates.UpdatePage(record, errors))
		}
		success := db.UpdateRecord(id, record)
		if success {
			return c.Redirect(http.StatusSeeOther, fmt.Sprintf("/?success=%v updated!", record.FirstName))
		} else {
			return render(c, http.StatusOK, templates.UpdatePage(record, []string{"Could not update that record!"}))
		}
	})

	e.GET("/delete/:id", func(c echo.Context) error {
		id, err := strconv.Atoi(c.Param("id"))
		if err != nil {
			return c.Redirect(http.StatusSeeOther, "/?error=Invalid id.")
		}
		success := db.DeleteRecord(id)
		if success {
			return c.Redirect(http.StatusSeeOther, "/?success=Deleted record.")
		} else {
			return c.Redirect(http.StatusSeeOther, "/?error=Could not delete record!")
		}
	})

	e.Logger.Fatal((e.Start(":3000")))
}

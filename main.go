package main

import (
	"encoding/gob"
	"fmt"
	"net/http"
	"strconv"

	"github.com/Aidunlin/go-srm/app"
	"github.com/Aidunlin/go-srm/db"
	"github.com/Aidunlin/go-srm/templates"
	"github.com/a-h/templ"
	"github.com/gorilla/securecookie"
	"github.com/gorilla/sessions"
	"github.com/labstack/echo-contrib/session"
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
	gob.Register(app.StudentRecord{})

	e := echo.New()
	e.Use(middleware.Logger(), session.Middleware(sessions.NewCookieStore(securecookie.GenerateRandomKey(32))))
	e.Static("/css", "css")
	e.Static("/js", "js")

	e.GET("/", func(c echo.Context) error {
		queryParams := c.QueryParams()
		tableParams := app.NewRecordTableParams(queryParams)
		total, records := db.SelectRecords(tableParams)
		messageParams := app.NewMessageParams(queryParams)
		return render(c, http.StatusOK, templates.IndexPage(total, records, tableParams, messageParams))
	})

	e.GET("/search", func(c echo.Context) error {
		queryParams := c.QueryParams()
		tableParams := app.NewRecordTableParams(queryParams)
		total, records := db.SelectRecords(tableParams)
		return render(c, http.StatusOK, templates.SearchPage(total, records, tableParams))
	})

	e.GET("/advanced-search", func(c echo.Context) error {
		sess, err := session.Get("session", c)
		if err == nil {
			println("GET found session")
			record, ok := sess.Values["advanced-search"].(app.StudentRecord)
			if ok {
				println("GET found session data")
				queryParams := c.QueryParams()
				tableParams := app.NewRecordTableParams(queryParams)
				total, records := db.SelectRecordsWithForm(tableParams, record)
				return render(c, http.StatusOK, templates.AdvancedSearchPage(record, total, records, tableParams, true))
			} else {
				println("GET could not find session data")
			}
		} else {
			c.SetCookie(&http.Cookie{Name: "session", Path: "/", MaxAge: -1})
			println("GET could not find session", err.Error())
		}
		return render(c, http.StatusOK, templates.AdvancedSearchPage(app.StudentRecord{}, 0, nil, app.RecordTableParams{}, false))
	})

	e.POST("/advanced-search", func(c echo.Context) error {
		formParams, _ := c.FormParams()
		record := app.NewAdvancedSearchForm(formParams)
		sess, err := session.Get("session", c)
		if err == nil {
			println("POST found session")
			if formParams.Has("reset") {
				sess.Options.MaxAge = -1
				saveErr := sess.Save(c.Request(), c.Response())
				if saveErr == nil {
					println("POST deleted session")
					return render(c, http.StatusOK, templates.AdvancedSearchPage(app.StudentRecord{}, 0, nil, app.RecordTableParams{}, false))
				} else {
					println("POST could not delete session", saveErr.Error())
				}
			} else {
				sess.Values["advanced-search"] = record
				saveErr := sess.Save(c.Request(), c.Response())
				if saveErr == nil {
					println("POST saved session")
				} else {
					println("POST could not save session", saveErr.Error())
				}
			}
		} else {
			println("POST could not find session", err.Error())
		}
		queryParams := c.QueryParams()
		tableParams := app.NewRecordTableParams(queryParams)
		total, records := db.SelectRecordsWithForm(tableParams, record)
		return render(c, http.StatusOK, templates.AdvancedSearchPage(record, total, records, tableParams, true))
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

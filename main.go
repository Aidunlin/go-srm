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
	"github.com/gorilla/securecookie"
	"github.com/gorilla/sessions"
	"github.com/labstack/echo-contrib/session"
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
	e.Use(middleware.Logger(), session.Middleware(sessions.NewCookieStore(securecookie.GenerateRandomKey(32))))
	e.Static("/css", "css")

	e.GET("/", getIndex)
	e.GET("/search", getSearch)
	e.GET("/advanced-search", getAdvancedSearch)

	e.GET("/create", getCreate)
	e.POST("/create", postCreate)
	e.GET("/update/:id", getUpdate)
	e.POST("/update/:id", postUpdate)
	e.GET("/delete/:id", getDelete)

	e.GET("/register", getRegister)
	e.POST("/register", postRegister)
	e.GET("/login", getLogin)
	e.POST("/login", postLogin)
	e.GET("/logout", getLogout)

	e.Logger.Fatal((e.Start(":3000")))
}

func getIndex(c echo.Context) error {
	queryParams := c.QueryParams()
	tableParams := app.NewRecordTableParams(queryParams)
	total, records := db.SelectRecords(tableParams)
	params := map[any]any{
		"table":   tableParams,
		"message": app.NewMessageParams(queryParams),
	}
	return render(c, http.StatusOK, templates.IndexPage(total, records), params)
}

func getSearch(c echo.Context) error {
	queryParams := c.QueryParams()
	tableParams := app.NewRecordTableParams(queryParams)
	total, records := db.SelectRecords(tableParams)
	params := map[any]any{
		"table": tableParams,
	}
	return render(c, http.StatusOK, templates.SearchPage(total, records), params)
}

func getAdvancedSearch(c echo.Context) error {
	formParams, _ := c.FormParams()
	form := app.AdvancedSearchParams(formParams)
	if len(form) == 0 {
		return render(c, http.StatusOK, templates.AdvancedSearchPage(app.StudentRecord{}, 0, nil, false), nil)
	}
	record := app.NewStudentFromAdvancedSearchForm(formParams)
	queryParams := c.QueryParams()
	tableParams := app.NewRecordTableParams(queryParams)
	total, records := db.SelectRecordsWithForm(tableParams, record)
	params := map[any]any{
		"table": tableParams,
		"form":  form,
	}
	return render(c, http.StatusOK, templates.AdvancedSearchPage(record, total, records, true), params)
}

func getCreate(c echo.Context) error {
	sess, sessErr := session.Get("session", c)
	if sessErr != nil {
		return c.Redirect(http.StatusSeeOther, "/?error=Could not establish a session.")
	}
	_, adminOk := sess.Values["adminId"]
	if !adminOk {
		sess.Options.MaxAge = -1
		if err := sess.Save(c.Request(), c.Response()); err != nil {
			return c.Redirect(http.StatusSeeOther, "/?error=Could not reset the session.")
		}
		return c.Redirect(http.StatusSeeOther, "/login")
	}

	return render(c, http.StatusOK, templates.CreatePage(app.StudentRecord{}, nil), nil)
}

func postCreate(c echo.Context) error {
	sess, sessErr := session.Get("session", c)
	if sessErr != nil {
		return c.Redirect(http.StatusSeeOther, "/?error=Could not establish a session.")
	}
	_, adminOk := sess.Values["adminId"]
	if !adminOk {
		sess.Options.MaxAge = -1
		if err := sess.Save(c.Request(), c.Response()); err != nil {
			return c.Redirect(http.StatusSeeOther, "/?error=Could not reset the session.")
		}
		return c.Redirect(http.StatusSeeOther, "/login")
	}

	formParams, _ := c.FormParams()
	record, errors := app.NewStudentFromCreateForm(formParams)
	if len(errors) > 0 {
		return render(c, http.StatusOK, templates.CreatePage(record, errors), nil)
	}
	success := db.InsertRecord(record)
	if !success {
		return render(c, http.StatusOK, templates.CreatePage(record, []string{"Could not save that record!"}), nil)
	}
	return c.Redirect(http.StatusSeeOther, fmt.Sprintf("/?success=%v created!", record.FirstName))
}

func getUpdate(c echo.Context) error {
	sess, sessErr := session.Get("session", c)
	if sessErr != nil {
		return c.Redirect(http.StatusSeeOther, "/?error=Could not establish a session.")
	}
	_, adminOk := sess.Values["adminId"]
	if !adminOk {
		sess.Options.MaxAge = -1
		if err := sess.Save(c.Request(), c.Response()); err != nil {
			return c.Redirect(http.StatusSeeOther, "/?error=Could not reset the session.")
		}
		return c.Redirect(http.StatusSeeOther, "/login")
	}

	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return c.Redirect(http.StatusSeeOther, "/?error=Invalid id.")
	}
	success, record := db.SelectRecord(id)
	if !success {
		return c.Redirect(http.StatusSeeOther, "/?error=Could not get record!")
	}
	return render(c, http.StatusOK, templates.UpdatePage(record, nil), nil)
}

func postUpdate(c echo.Context) error {
	sess, sessErr := session.Get("session", c)
	if sessErr != nil {
		return c.Redirect(http.StatusSeeOther, "/?error=Could not establish a session.")
	}
	_, adminOk := sess.Values["adminId"]
	if !adminOk {
		sess.Options.MaxAge = -1
		if err := sess.Save(c.Request(), c.Response()); err != nil {
			return c.Redirect(http.StatusSeeOther, "/?error=Could not reset the session.")
		}
		return c.Redirect(http.StatusSeeOther, "/login")
	}

	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return c.Redirect(http.StatusSeeOther, "/?error=Invalid id.")
	}
	formParams, _ := c.FormParams()
	record, errors := app.NewStudentFromUpdateForm(formParams)
	if len(errors) > 0 {
		return render(c, http.StatusOK, templates.UpdatePage(record, errors), nil)
	}
	success := db.UpdateRecord(id, record)
	if !success {
		return render(c, http.StatusOK, templates.UpdatePage(record, []string{"Could not update that record!"}), nil)
	}
	return c.Redirect(http.StatusSeeOther, fmt.Sprintf("/?success=%v updated!", record.FirstName))
}

func getDelete(c echo.Context) error {
	sess, sessErr := session.Get("session", c)
	if sessErr != nil {
		return c.Redirect(http.StatusSeeOther, "/?error=Could not establish a session.")
	}
	_, adminOk := sess.Values["adminId"]
	if !adminOk {
		sess.Options.MaxAge = -1
		if err := sess.Save(c.Request(), c.Response()); err != nil {
			return c.Redirect(http.StatusSeeOther, "/?error=Could not reset the session.")
		}
		return c.Redirect(http.StatusSeeOther, "/login")
	}

	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return c.Redirect(http.StatusSeeOther, "/?error=Invalid id.")
	}
	success := db.DeleteRecord(id)
	if !success {
		return c.Redirect(http.StatusSeeOther, "/?error=Could not delete record!")
	}
	return c.Redirect(http.StatusSeeOther, "/?success=Deleted record.")
}

func getRegister(c echo.Context) error {
	return render(c, http.StatusOK, templates.RegisterPage(app.AdminRecord{}, nil), nil)
}

func postRegister(c echo.Context) error {
	formParams, _ := c.FormParams()
	admin, errors := app.NewAdminFromRegisterForm(formParams)
	if len(errors) > 0 {
		return render(c, http.StatusOK, templates.RegisterPage(admin, errors), nil)
	}
	return render(c, http.StatusOK, templates.RegisterPage(admin, errors), nil)
}

func getLogin(c echo.Context) error {
	return render(c, http.StatusOK, templates.LoginPage(app.AdminRecord{}, nil), nil)
}

func postLogin(c echo.Context) error {
	formParams, _ := c.FormParams()
	admin, errors := app.NewAdminFromLoginForm(formParams)
	if len(errors) > 0 {
		return render(c, http.StatusOK, templates.LoginPage(admin, errors), nil)
	}
	return render(c, http.StatusOK, templates.LoginPage(admin, nil), nil)
}

func getLogout(c echo.Context) error {
	sess, err := session.Get("session", c)
	if err != nil {
		return c.Redirect(http.StatusSeeOther, "/?error=Could not establish a session.")
	}
	sess.Options.MaxAge = -1
	if err := sess.Save(c.Request(), c.Response()); err != nil {
		return c.Redirect(http.StatusSeeOther, "?error=Could not log out.")
	}
	return c.Redirect(http.StatusSeeOther, "/?success=Logged out!")
}

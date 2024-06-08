package routes

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/Aidunlin/go-srm/db"
	"github.com/Aidunlin/go-srm/model"
	"github.com/Aidunlin/go-srm/templates"
	"github.com/labstack/echo-contrib/session"
	"github.com/labstack/echo/v4"
)

func AddListsRoutes(e *echo.Echo) {
	e.GET("/create", getCreate)
	e.POST("/create", postCreate)
	e.GET("/update/:id", getUpdate)
	e.POST("/update/:id", postUpdate)
	e.GET("/delete/:id", getDelete)
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

	return render(c, http.StatusOK, templates.CreatePage(model.StudentRecord{}, nil), nil)
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
	student, errors := model.NewStudentRecord(formParams)
	if len(errors) > 0 {
		return render(c, http.StatusOK, templates.CreatePage(student, errors), nil)
	}
	success := db.InsertStudent(student)
	if !success {
		return render(c, http.StatusOK, templates.CreatePage(student, []string{"Could not create student!"}), nil)
	}
	return c.Redirect(http.StatusSeeOther, fmt.Sprintf("/?success=%v created!", student.FirstName))
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
	student, success := db.SelectStudent(id)
	if !success {
		return c.Redirect(http.StatusSeeOther, "/?error=Could not get student!")
	}
	return render(c, http.StatusOK, templates.UpdatePage(student, nil), nil)
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
	student, errors := model.NewStudentRecord(formParams)
	if len(errors) > 0 {
		return render(c, http.StatusOK, templates.UpdatePage(student, errors), nil)
	}
	success := db.UpdateStudent(id, student)
	if !success {
		return render(c, http.StatusOK, templates.UpdatePage(student, []string{"Could not update student!"}), nil)
	}
	return c.Redirect(http.StatusSeeOther, fmt.Sprintf("/?success=%v updated!", student.FirstName))
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
	success := db.DeleteStudent(id)
	if !success {
		return c.Redirect(http.StatusSeeOther, "/?error=Could not delete student!")
	}
	return c.Redirect(http.StatusSeeOther, "/?success=Deleted student.")
}

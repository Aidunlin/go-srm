package routes

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/Aidunlin/go-srm/db"
	"github.com/Aidunlin/go-srm/model"
	"github.com/Aidunlin/go-srm/templates"
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
	if !isAuthenticated(c) {
		return c.Redirect(http.StatusSeeOther, "/?error=Not authorized!")
	}

	return render(c, http.StatusOK, templates.CreatePage(model.StudentRecord{}, nil), nil)
}

func postCreate(c echo.Context) error {
	if !isAuthenticated(c) {
		return c.Redirect(http.StatusSeeOther, "/?error=Not authorized!")
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
	if !isAuthenticated(c) {
		return c.Redirect(http.StatusSeeOther, "/?error=Not authorized!")
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
	if !isAuthenticated(c) {
		return c.Redirect(http.StatusSeeOther, "/?error=Not authorized!")
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
	if !isAuthenticated(c) {
		return c.Redirect(http.StatusSeeOther, "/?error=Not authorized!")
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

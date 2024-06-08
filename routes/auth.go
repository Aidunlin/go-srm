package routes

import (
	"net/http"

	"github.com/Aidunlin/go-srm/model"
	"github.com/Aidunlin/go-srm/templates"
	"github.com/labstack/echo-contrib/session"
	"github.com/labstack/echo/v4"
)

func AddAuthRoutes(e *echo.Echo) {
	e.GET("/register", getRegister)
	e.POST("/register", postRegister)
	e.GET("/login", getLogin)
	e.POST("/login", postLogin)
	e.GET("/logout", getLogout)
}

func getRegister(c echo.Context) error {
	return render(c, http.StatusOK, templates.RegisterPage(model.AdminRecord{}, nil), nil)
}

func postRegister(c echo.Context) error {
	formParams, _ := c.FormParams()
	admin, errors := model.NewAdminRecordFromRegisterForm(formParams)
	if len(errors) > 0 {
		return render(c, http.StatusOK, templates.RegisterPage(admin, errors), nil)
	}
	return render(c, http.StatusOK, templates.RegisterPage(admin, errors), nil)
}

func getLogin(c echo.Context) error {
	return render(c, http.StatusOK, templates.LoginPage(model.AdminRecord{}, nil), nil)
}

func postLogin(c echo.Context) error {
	formParams, _ := c.FormParams()
	admin, errors := model.NewAdminRecordFromLoginForm(formParams)
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

package routes

import (
	"net/http"

	"github.com/Aidunlin/go-srm/db"
	"github.com/Aidunlin/go-srm/model"
	"github.com/Aidunlin/go-srm/templates"
	"github.com/labstack/echo-contrib/session"
	"github.com/labstack/echo/v4"
	"golang.org/x/crypto/bcrypt"
)

func AddAuthRoutes(e *echo.Echo) {
	e.GET("/register", getRegister)
	e.POST("/register", postRegister)
	e.GET("/login", getLogin)
	e.POST("/login", postLogin)
	e.GET("/logout", getLogout)
}

func getRegister(c echo.Context) error {
	_, err := session.Get("session", c)
	if err != nil {
		c.SetCookie(&http.Cookie{Name: "session", MaxAge: -1})
	}

	return render(c, http.StatusOK, templates.RegisterPage(model.AdminRecord{}, nil), nil)
}

func postRegister(c echo.Context) error {
	formParams, _ := c.FormParams()
	admin, errors := model.NewAdminRecordFromRegisterForm(formParams)
	if len(errors) > 0 {
		return render(c, http.StatusOK, templates.RegisterPage(admin, errors), nil)
	}

	success := db.InsertAdmin(admin)
	if !success {
		return render(c, http.StatusOK, templates.RegisterPage(admin, []string{"Could not register!"}), nil)
	}

	sess, err := session.Get("session", c)
	if err != nil {
		c.SetCookie(&http.Cookie{Name: "session", MaxAge: -1})
		return render(c, http.StatusOK, templates.LoginPage(admin, []string{"Something weird happened, try again"}), nil)
	}

	sess.Values["adminId"] = admin.Id
	if err := sess.Save(c.Request(), c.Response()); err != nil {
		return c.Redirect(http.StatusSeeOther, "/?error=Could not log in.")
	}

	return c.Redirect(http.StatusSeeOther, "/?success=Registered and logged in!")
}

func getLogin(c echo.Context) error {
	_, err := session.Get("session", c)
	if err != nil {
		c.SetCookie(&http.Cookie{Name: "session", MaxAge: -1})
	}

	return render(c, http.StatusOK, templates.LoginPage(model.AdminRecord{}, nil), nil)
}

func postLogin(c echo.Context) error {
	formParams, _ := c.FormParams()
	admin, errors := model.NewAdminRecordFromLoginForm(formParams)
	if len(errors) > 0 {
		return render(c, http.StatusOK, templates.LoginPage(admin, errors), nil)
	}

	inputtedPassword := formParams.Get("password")
	admin, success := db.SelectAdminWithEmail(admin.Email)
	if !success || bcrypt.CompareHashAndPassword([]byte(admin.Password), []byte(inputtedPassword)) != nil {
		errors := []string{"Invalid <strong>email</strong> or <strong>password</strong>!"}
		return render(c, http.StatusOK, templates.LoginPage(admin, errors), nil)
	}

	sess, err := session.Get("session", c)
	if err != nil {
		c.SetCookie(&http.Cookie{Name: "session", MaxAge: -1})
		return render(c, http.StatusOK, templates.LoginPage(admin, []string{"Something weird happened, try again"}), nil)
	}

	sess.Values["adminId"] = admin.Id
	if err := sess.Save(c.Request(), c.Response()); err != nil {
		return c.Redirect(http.StatusSeeOther, "/?error=Could not log in.")
	}

	return c.Redirect(http.StatusSeeOther, "/?success=Logged in!")
}

func getLogout(c echo.Context) error {
	c.SetCookie(&http.Cookie{Name: "session", MaxAge: -1})
	return c.Redirect(http.StatusSeeOther, "/?success=Logged out!")
}

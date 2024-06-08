package main

import (
	"github.com/Aidunlin/go-srm/routes"
	"github.com/gorilla/securecookie"
	"github.com/gorilla/sessions"
	"github.com/labstack/echo-contrib/session"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

// Entry point for the web server.
func main() {
	e := echo.New()
	e.Use(middleware.Logger(), session.Middleware(sessions.NewCookieStore(securecookie.GenerateRandomKey(32))))
	e.Static("/css", "css")

	routes.AddMainRoutes(e)
	routes.AddListsRoutes(e)
	routes.AddAuthRoutes(e)

	e.Logger.Fatal((e.Start(":3000")))
}

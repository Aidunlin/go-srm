package routes

import (
	"context"

	"github.com/a-h/templ"
	"github.com/labstack/echo-contrib/session"
	"github.com/labstack/echo/v4"
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

func isAuthenticated(c echo.Context) bool {
	isAuthed := false
	if sess, sessErr := session.Get("session", c); sessErr == nil {
		_, isAuthed = sess.Values["adminId"]
	}
	return isAuthed
}

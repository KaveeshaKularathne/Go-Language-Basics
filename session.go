package main

import (
	"github.com/gorilla/sessions"
	"github.com/labstack/echo-contrib/session"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"net/http"
)

func main()  {
	e := echo.New()
	e.Use(session.Middleware(sessions.NewCookieStore([]byte("secret"))))

	e.GET("/", func(c echo.Context) error {
		sess, _ := session.Get("session", c)
		sess.Options = &sessions.Options{
			Path:     "/",
			MaxAge:   86400 * 7,
			HttpOnly: true,
		}
		sess.Values["foo"] = "bar"
		sess.Save(c.Request(), c.Response())
		return c.NoContent(http.StatusOK)
	})
e.Use(middleware.Static("/static"))
e.Use(middleware.StaticWithConfig(middleware.StaticConfig{
	Root: "static",
	Browse: true,
}))

e.Pre(middleware.AddTrailingSlashWithConfig(middleware.TrailingSlashConfig{
	RedirectCode: http.StatusMovedPermanently,
}))


}

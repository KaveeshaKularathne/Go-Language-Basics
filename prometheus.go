package main
import (
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo-contrib/prometheus"
	"github.com/labstack/echo/v4/middleware"
	"github.com/labstack/gommon/log"
	"strings"
)



func urlSkipper(c echo.Context) bool {
	if strings.HasPrefix(c.Path(), "/testurl") {
		return true
	}
	return false
}

func main() {
	e := echo.New()
	// Enable metrics middleware
	p := prometheus.NewPrometheus("echo", urlSkipper)
	p.Use(e)

	e.Logger.Fatal(e.Start(":8080"))

	e.Use(middleware.RecoverWithConfig(middleware.RecoverConfig{
		StackSize: 1 << 10, // 1 KB
		LogLevel:  log.ERROR,
	}))
	e.Use(middleware.RequestIDWithConfig(middleware.RequestIDConfig{
		Generator: func() string {
			return customGenerator()
		},
	}))


	e.Pre(middleware.Rewrite(map[string]string{
		"/old":              "/new",
		"/api/*":            "/$1",
		"/js/*":             "/public/javascripts/$1",
		"/users/*/orders/*": "/user/$1/order/$2",
	}))
	e.Pre(middleware.RecoverWithConfig(middleware.RewriteConfig{}))
}


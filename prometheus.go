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
}


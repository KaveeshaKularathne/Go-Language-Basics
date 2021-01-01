package main

import (

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/labstack/echo-contrib/jaegertracing"
	"net/http"
	"strings"
	"time"
)
func urlSkipper(c echo.Context) bool {
	if strings.HasPrefix(c.Path(), "/testurl") {
		return true
	}
	return false
}

func main()  {
	e := echo.New()
	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins: []string{"https://labstack.com", "https://labstack.net"},
		AllowHeaders: []string{echo.HeaderOrigin, echo.HeaderContentType, echo.HeaderAccept},
	}))
	e.Use(middleware.GzipWithConfig(middleware.GzipConfig{
		Level: 5,
	}))
	c := jaegertracing.New(e, nil)
	defer c.Close()
	e.GET("/", func(c echo.Context) error {
		// Wrap slowFunc on a new span to trace it's execution passing the function arguments
		jaegertracing.TraceFunction(c, slowFunc, "Test String")
		return c.String(http.StatusOK, "Hello, World!")
	})



	e.Logger.Fatal(e.Start(":1323"))


}
func slowFunc(s string) {
	time.Sleep(200 * time.Millisecond)
	return
}


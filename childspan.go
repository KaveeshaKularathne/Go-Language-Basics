package main

import (
	"github.com/labstack/echo-contrib/jaegertracing"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"net/http"
	"time"
)

func main()  {
	e := echo.New()

	c := jaegertracing.New(e, nil)
	defer c.Close()
	e.GET("/", func(c echo.Context) error {
		time.Sleep(40 * time.Millisecond)
		sp := jaegertracing.CreateChildSpan(c, "Child span for additional processing")
		defer sp.Finish()
		sp.LogEvent("Test log")
		sp.SetBaggageItem("Test baggage", "baggage")
		sp.SetTag("Test tag", "New Tag")
		time.Sleep(100 * time.Millisecond)
		return c.String(http.StatusOK, "Hello, World!")
	})
	e.Logger.Fatal(e.Start(":8088"))

	e.Use(middleware.KeyAuth(func(key string, c echo.Context) (bool, error) {
		return key == "valid-key", nil
	}))
	e.Use(middleware.LoggerWithConfig(middleware.LoggerConfig{
		Format: "method=${method}, uri=${uri}, status=${status}\n",
	}))
	e.Pre(middleware.MethodOverrideWithConfig(middleware.MethodOverrideConfig{
		Getter: middleware.MethodFromForm("_method"),
	}))
}




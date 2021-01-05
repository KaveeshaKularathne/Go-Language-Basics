package main

import (
	"net/http"

	"github.com/GeertJohan/go.rice"
	"github.com/labstack/echo/v4"
)

func main() {
	e := echo.New()


	assetHandler := http.FileServer(rice.MustFindBox("app").HTTPBox())

	e.GET("/", echo.WrapHandler(assetHandler))


	e.GET("/static/*", echo.WrapHandler(http.StripPrefix("/static/", assetHandler)))

	e.Logger.Fatal(e.Start(":1323"))
}


package main


import (
	"github.com/labstack/echo/v4/middleware"
	"net/http"

	"github.com/labstack/echo/v4"
)

func hello(c echo.Context) error  {
	return c.String(http.StatusOK,"Hello,World")

}

func main() {
	e := echo.New()

	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "Hello, World!")
	})
	e.Logger.Fatal(e.Start(":1323"))

	e.GET("/hello",hello)


	e.GET("/users/:id", func(c echo.Context) error {
		return c.String(http.StatusOK, "/users/:id")
	})

	e.GET("/users/new", func(c echo.Context) error {
		return c.String(http.StatusOK, "/users/new")
	})

	e.GET("/users/1/files/*", func(c echo.Context) error {
		return c.String(http.StatusOK, "/users/1/files/*")
	})

	g:=e.Group("/admin")
	g.Use(middleware.BasicAuth(func(username,password string,c echo.Context)(bool,error){
		if username=="joe"&&password=="secret" {
			return true,nil

		}
		return false,nil
	}))

	route:=e.POST("/users",func(c echo.Context)error{


	})
	route.Name="create-user"

	e.GET("/users/:id", func(c echo.Context) error {

	}).Name="get-user"

	e.Static("/static","assets")
	e.File("/", "public/index.html")


}
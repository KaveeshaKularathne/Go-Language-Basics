package main


import (
	"github.com/labstack/echo/v4/middleware"
	"io"
	"net/http"
	"github.com/labstack/echo/v4"
	"os"
)
func getUser(c echo.Context )error{
	id:=c.Param("id")
	return  c.String(http.StatusOK,id)

}
func show(c echo.Context) error {
	team := c.QueryParam("team")
	member := c.QueryParam("member")
	return c.String(http.StatusOK, "team:" + team + ", member:" + member)
}
func save(c echo.Context) error {

	name := c.FormValue("name")
	avatar,err:=c.FormFile("avatarar")
	if err!=nil{
		return err
	}
	src,err:=avatar.Open()
	if err!=nil{
		return err
	}
	defer src.Close()

	dst,err:=os.Create(avatar.Filename)
	if err!=nil{
		return err
	}
	defer  dst.Close()

	if _, err=io.Copy(dst,src);
	err!=nil{
		return err
	}
	return c.HTML(http.StatusOK, "<b>Thank you! " + name + "</b>")
	email := c.FormValue("email")
	return c.String(http.StatusOK, "name:" + name + ", email:" + email)
}

func main() {



	e := echo.New()
	e.GET("/users/:id", getUser)

	e.GET("/", func(c echo.Context) error {

		return c.String(http.StatusOK, "Hello, World!")
	})


	e.Logger.Fatal(e.Start(":8088"))

	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	// Group level middleware
	g := e.Group("/admin")
	g.Use(middleware.BasicAuth(func(username, password string, c echo.Context) (bool, error) {
		if username == "joe" && password == "secret" {
			return true, nil
		}
		return false, nil
	})) 

	// Route level middleware
	track := func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			println("request to /users")
			return next(c)
		}
	}
	e.GET("/users", func(c echo.Context) error {
		return c.String(http.StatusOK, "/users")
	}, track)
}
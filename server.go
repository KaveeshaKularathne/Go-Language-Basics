package main


import (
	"fmt"
	"github.com/labstack/echo/v4/middleware"
	"io"
	"net/http"
	"github.com/labstack/echo/v4"
	"os"
	"github.com/labstack/gommon/log"
	"golang.org/x/net/http2"
	"time"
)


type User struct {
	Name string `json:"name" form:"name" query:"name"`
	Email string `json:"email" form:"email" query:"email"`


}
type CustomContext struct {
	echo.Context
}




func (c *CustomContext) Foo() {
	println("foo")
}

func (c *CustomContext) Bar() {
	println("bar")
}

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

func writeCookie(c echo.Context)error{
	cookie:=new(http.Cookie)
	cookie.Name="username"
	cookie.Value="join"
	cookie.Expires=time.Now().Add(24*time.Hour)
	c.SetCookie(cookie)
	return c.String(http.StatusOK,"write a cookie")
}

func readCookie(c echo.Context)error  {
	cookie,err:=c.Cookie("username")
	if err !=nil{
		return err
	}
	fmt.Println(cookie.Name)
	fmt.Println(cookie.Value)
	return  c.String(http.StatusOK,"read a cookie")

}

func readAllCookies(c echo.Context)error  {
	for _,cookie:=range c.Cookies() {
		fmt.Println(cookie.Name)
		fmt.Println(cookie.Value)

	}
	return  c.String(http.StatusOK,"read all tha cookies")

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


	g := e.Group("/admin")
	g.Use(middleware.BasicAuth(func(username, password string, c echo.Context) (bool, error) {
		if username == "joe" && password == "secret" {
			return true, nil
		}
		return false, nil
	}))


	track := func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			println("request to /users")
			return next(c)
		}
	}
	e.GET("/users", func(c echo.Context) error {
		return c.String(http.StatusOK, "/users")
	}, track)
	if l, ok := e.Logger.(*log.Logger); ok {
		l.SetHeader("${time_rfc3339} ${level}")
	}

	s := &http2.Server{
		MaxConcurrentStreams: 250,
		MaxReadFrameSize:     1048576,
		IdleTimeout:          10 * time.Second,
	}
	e.Logger.Fatal(e.StartH2CServer(":8088", s))

	e.Use(func(next echo.HandlerFunc) echo.HandlerFunc{
		return func(c echo.Context) error {
			cc:=&CustomContext{c}
			return  next(cc)

		}

	})
	e.GET("/", func(c echo.Context) error {
		cc:=c.(*CustomContext)
		cc.Foo()
		cc.Bar()
		return cc.String(200,"OK")

	})
	e.Use(func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			return echo.NewHTTPError(http.StatusUnauthorized,"Please provide valid credentials")
		}

	})
}
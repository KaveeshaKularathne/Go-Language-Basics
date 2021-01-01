package main

import (
	"crypto/subtle"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"net/http"
)

type (
	user struct {
		Name string `json:"name" from :"name"`
		Email string `json:"email" from:"email"`
	}
	handler struct {
		db map[string]*user
	}
)

func (h *handler)createUser(c echo.Context)error  {
	u:=new(User)
	if err:=c.Bind(u);err!=nil{
		return err
	}
	return c.JSON(http.StatusCreated,u)

}
func (h*handler)getUser(c echo.Context)error {
	email := c.Param("email")
	user := h.db[email]
	if user == nil {
		return echo.NewHTTPError(http.StatusNotFound, "user not found")

	}
	return c.JSON(http.StatusOK, user)
}
func main()  {
	e := echo.New()
	e.Use(middleware.BasicAuth(func(username, password string, c echo.Context) (bool, error) {
		// Be careful to use constant time comparison to prevent timing attacks
		if subtle.ConstantTimeCompare([]byte(username), []byte("joe")) == 1 &&
			subtle.ConstantTimeCompare([]byte(password), []byte("secret")) == 1 {
			return true, nil
		}
		return false, nil
	}))
	e.Use(middleware.BodyDump(func(c echo.Context, reqBody, resBody []byte) {
	}))
	e.Use(middleware.BodyLimit("2M"))



}





package api

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

func SetBindHandler(g *echo.Group) {
	g.GET("/binder", useBinder)
	g.POST("/basic", useBind)
}

// POST /bind/basic
func useBind(c echo.Context) error {
	form := &recieveUser{}
	if err := c.Bind(form); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	u := &user{
		Name:    form.Name,
		Email:   form.Email,
		IsAdmin: false,
	}

	doSomeBusinessLogic(u)

	return c.JSON(http.StatusOK, echo.Map{"data": form})
}

func doSomeBusinessLogic(u *user) {
	// do some business logic
}

// url =  "/bind/binder?active=true&id=1&id=2&id=3&length=25"
func useBinder(c echo.Context) error {
	var opts struct {
		IDs    []int64
		Active bool
		Length int64
	}

	// creates query params binder that stops binding at first error
	err := echo.QueryParamsBinder(c).
		Int64("length", &opts.Length).
		Int64s("id", &opts.IDs).
		Bool("active", &opts.Active).
		BindError() // returns first binding error

	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	return c.JSON(http.StatusOK, echo.Map{"data": &opts})
}

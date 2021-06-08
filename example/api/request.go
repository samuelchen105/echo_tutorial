package api

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

func SetRequestHandler(g *echo.Group) {
	g.POST("/form", getForm)
	g.GET("/query", getQueryParam)
	g.GET("/path/:name", getPathParam)
	g.POST("/validate", validateForm)
}

func getForm(c echo.Context) error {
	name := c.FormValue("name")
	reply := echo.Map{
		"data": echo.Map{"name": name},
	}
	return c.JSON(http.StatusOK, reply)
}

func getQueryParam(c echo.Context) error {
	name := c.QueryParam("name")
	reply := echo.Map{
		"data": echo.Map{"name": name},
	}
	return c.JSON(http.StatusOK, reply)
}

func getPathParam(c echo.Context) error {
	name := c.Param("name")
	reply := echo.Map{
		"data": echo.Map{"name": name},
	}
	return c.JSON(http.StatusOK, reply)
}

func validateForm(c echo.Context) (err error) {
	form := &recieveUser{}
	if err = c.Bind(form); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	if err = c.Validate(form); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	return c.JSON(http.StatusOK, echo.Map{"data": form})
}

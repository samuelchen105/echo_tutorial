package api

import (
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func SetCsrfHandler(g *echo.Group) {
	g.Use(middleware.CSRFWithConfig(middleware.CSRFConfig{
		TokenLookup: "form:csrf",
	}))
	g.GET("", showCsrfExample)
	g.POST("", doCsrfExample)
}

func showCsrfExample(c echo.Context) error {
	csrf, ok := c.Get(middleware.DefaultCSRFConfig.ContextKey).(string)
	if !ok {
		return echo.NewHTTPError(http.StatusInternalServerError, "convert not expected type")
	}
	data := struct {
		Csrf string
	}{
		Csrf: csrf,
	}
	return c.Render(http.StatusOK, "csrf", data)
}

func doCsrfExample(c echo.Context) error {
	email := c.FormValue("email")
	reply := echo.Map{
		"data": echo.Map{"email": email},
	}
	return c.JSON(http.StatusOK, reply)
}

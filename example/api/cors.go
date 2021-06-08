package api

import (
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func SetCorsHandler(g *echo.Group) {
	g.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins: []string{"https://example.com"},
		AllowHeaders: []string{echo.HeaderOrigin},
	}))
	g.GET("", corsExample)
}

func corsExample(c echo.Context) error {
	return c.String(http.StatusOK, "CORS")
}

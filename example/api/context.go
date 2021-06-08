package api

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

func SetContextHandler(g *echo.Group) {
	g.Use(ccMiddleware)
	g.GET("", contextExample)
}

type CustomContext struct {
	echo.Context
}

func (c *CustomContext) SendHello() error {
	return c.JSON(http.StatusOK, echo.Map{"msg": "hello, I am custom context"})
}

func ccMiddleware(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		cc := &CustomContext{c}
		return next(cc)
	}
}

func contextExample(c echo.Context) error {
	cc := c.(*CustomContext)
	return cc.SendHello()
}

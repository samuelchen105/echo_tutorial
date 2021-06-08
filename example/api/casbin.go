package api

import (
	"net/http"

	"github.com/casbin/casbin/v2"
	"github.com/labstack/echo/v4"
)

func SetCasbinHandler(g *echo.Group) error {
	enforcer, err := casbin.NewEnforcer(casbinModelPath, casbinConfigPath)
	if err != nil {
		return err
	}
	ce := Enforcer{enforcer: enforcer}
	g.Use(ce.Enforce)
	g.GET("/resource", casbinExample)
	return nil
}

var (
	casbinModelPath  = "configs/casbin_model.conf"
	casbinConfigPath = "configs/casbin_policy.csv"
)

type Enforcer struct {
	enforcer *casbin.Enforcer
}

func (e *Enforcer) Enforce(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		// this is a simple example, there are better ways to examine login
		cookie, err := c.Cookie("user")
		if err != nil {
			return echo.ErrForbidden
		}
		user := cookie.Value
		path := c.Request().URL.Path
		method := c.Request().Method

		isAllowed, err := e.enforcer.Enforce(user, path, method)
		if err != nil {
			return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
		}

		if isAllowed {
			return next(c)
		}
		return echo.ErrForbidden
	}
}

func casbinExample(c echo.Context) error {
	return c.JSON(http.StatusOK, echo.Map{"data": "resource"})
}

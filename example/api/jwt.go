package api

import (
	"fmt"
	"net/http"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func SetJwtHandler(g *echo.Group) {
	g.POST("/login", jwtLogin)
	sg := g.Group("/restricted")
	sg.Use(middleware.JWTWithConfig(middleware.JWTConfig{
		SigningKey: []byte("secret"),
	}))
	sg.GET("", jwtGetData)
}

func jwtLogin(c echo.Context) error {
	username := c.FormValue("username")
	password := c.FormValue("password")

	if username != "admin" || password != "pwd" {
		return echo.ErrUnauthorized
	}

	token := jwt.New(jwt.SigningMethodHS256)

	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok {
		return echo.NewHTTPError(http.StatusInternalServerError, "convert not expected type")
	}
	claims["name"] = "admin"
	claims["exp"] = time.Now().Add(time.Hour * 1).Unix()

	t, err := token.SignedString([]byte("secret"))
	if err != nil {
		return err
	}

	return c.JSON(http.StatusOK, map[string]string{
		"token": t,
	})
}

func jwtGetData(c echo.Context) error {
	user, ok := c.Get(middleware.DefaultJWTConfig.ContextKey).(*jwt.Token)
	if !ok {
		return echo.NewHTTPError(http.StatusInternalServerError, "convert not expected type")
	}
	claims, ok := user.Claims.(jwt.MapClaims)
	if !ok {
		return echo.NewHTTPError(http.StatusInternalServerError, "convert not expected type")
	}
	msg := fmt.Sprintf("Welcome %s", claims["name"])
	return c.String(http.StatusOK, msg)
}

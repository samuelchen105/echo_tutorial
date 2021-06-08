package api

import (
	"log"
	"net/http"
	"time"

	"github.com/labstack/echo/v4"
)

func SetCookieHandler(g *echo.Group) {
	g.POST("/create", createCookie)
	g.GET("/read", readCookie)
	g.GET("/readall", readAllCookies)
}

func createCookie(c echo.Context) error {
	cookie := &http.Cookie{
		Name:    "username",
		Value:   "Jack",
		Expires: time.Now().Add(24 * time.Hour),
	}
	c.SetCookie(cookie)
	log.Println("set cookie: ", cookie)
	return c.JSON(http.StatusOK, echo.Map{"msg": "write a cookie"})
}

func readCookie(c echo.Context) error {
	cookie, err := c.Cookie("username")
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err)
	}
	log.Println(cookie)
	return c.JSON(http.StatusOK, echo.Map{"msg": "read a cookie"})
}

func readAllCookies(c echo.Context) error {
	for _, cookie := range c.Cookies() {
		log.Println(cookie)
	}
	return c.JSON(http.StatusOK, echo.Map{"msg": "read all the cookies"})
}

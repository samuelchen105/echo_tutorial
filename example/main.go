package main

import (
	"net/http"

	"github.com/yuhsuan105/echo_tutorial/example/api"

	"github.com/labstack/echo/v4"
)

func main() {
	e := echo.New()
	if err := api.Setup(e); err != nil {
		panic(err)
	}

	e.GET("/", hello)
	e.Logger.Fatal(e.Start(":8080"))
}

func hello(c echo.Context) error {
	return c.String(http.StatusOK, "Hello, World!")
}

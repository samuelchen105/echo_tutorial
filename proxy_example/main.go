package main

import (
	"net/url"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

var (
	port       = "8081"
	serviceUrl = "http://localhost:8080"
)

func main() {
	e := echo.New()

	// Setup proxy
	url1, err := url.Parse(serviceUrl)
	if err != nil {
		e.Logger.Fatal(err)
	}

	e.Use(middleware.Proxy(middleware.NewRoundRobinBalancer([]*middleware.ProxyTarget{
		{
			URL: url1,
		},
	})))

	e.Logger.Fatal(e.Start(":" + port))
}

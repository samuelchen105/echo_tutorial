package api

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

func SetResponseHandler(g *echo.Group) {
	g.GET("/string", sendString)
	g.GET("/html", sendHTML)
	g.GET("/template", renderTemplate)
	g.GET("/json", sendJSON)
	g.GET("/xml", sendXML)
}

func sendString(c echo.Context) error {
	return c.String(http.StatusOK, "Send String")
}

func sendHTML(c echo.Context) error {
	return c.HTML(http.StatusOK, "<strong>Send HTML</strong>")
}

func renderTemplate(c echo.Context) error {
	data := "I'm data"
	return c.Render(http.StatusOK, "response", data)
}

func sendJSON(c echo.Context) error {
	u := &user{
		Name:  "Jon",
		Email: "jon@labstack.com",
	}
	return c.JSON(http.StatusOK, u)
}

func sendXML(c echo.Context) error {
	u := &user{
		Name:  "Jon",
		Email: "jon@labstack.com",
	}
	return c.XML(http.StatusOK, u)
}

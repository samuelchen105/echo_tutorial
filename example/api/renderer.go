package api

import (
	"io"
	"text/template"

	"github.com/labstack/echo/v4"
)

var templatePath = "templates"

type Template struct {
	templates *template.Template
}

func (t *Template) Render(w io.Writer, name string, data interface{}, c echo.Context) error {
	return t.templates.ExecuteTemplate(w, name, data)
}

func SetRenderer(e *echo.Echo) {
	t := &Template{
		templates: template.Must(template.ParseGlob(templatePath + "/*.html")),
	}
	e.Renderer = t
}

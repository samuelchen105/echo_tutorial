# 06 Response

## Send String

`example/api/response.go`

```go
func sendString(c echo.Context) error {
    return c.String(http.StatusOK, "Send String")
}
```

**Try it:**

```powershell
curl -X GET "http://localhost:8080/response/string"
```

```text
Send String
```

## Send HTML

`example/api/response.go`

```go
func sendHTML(c echo.Context) error {
    return c.HTML(http.StatusOK, "<strong>Send HTML</strong>")
}
```

**Try it:**

```powershell
curl -X GET "http://localhost:8080/response/html"
```

```text
<strong>Send HTML</strong>
```

### Render Template

If you are looking to send dynamically generate HTML, you can register custom renderer.

The example below uses Go `html/template`

Implement `echo.Render` interface:
`example/api/renderer.go`

```go
import (
    "io"
    "text/template"

    "github.com/labstack/echo/v4"
)

var templatePath = "example/templates"

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
```

Handler:
`example/api/response.go`

```go
func renderTemplate(c echo.Context) error {
    data := "I'm data"
    return c.Render(http.StatusOK, "layout", data)
}
```

**Try it:**

```powershell
curl -X GET "http://localhost:8080/response/template"
```

```text
<h2>This is data:</h2>
<p>I'm data</p>
```

For more details: [official document](https://echo.labstack.com/guide/templates/)

## Send JSON

`example/api/response.go`

```go
func sendJSON(c echo.Context) error {
    u := &user{
        Name:  "Jon",
        Email: "jon@labstack.com",
    }
    return c.JSON(http.StatusOK, u)
}
```

**Try it:**

```powershell
curl -X GET "http://localhost:8080/response/json"
```

```text
{"name":"Jon","email":"jon@labstack.com","is_admin":false}
```

## Send XML

`example/api/response.go`

```go
func sendXML(c echo.Context) error {
    u := &user{
        Name:  "Jon",
        Email: "jon@labstack.com",
    }
    return c.XML(http.StatusOK, u)
}
```

**Try it:**

```powershell
curl -X GET "http://localhost:8080/response/xml"
```

```text
<?xml version="1.0" encoding="UTF-8"?>
<user><name>Jon</name><email>jon@labstack.com</email><is_admin>false</is_admin></user>
```

## See Also

For more method, see [document](https://echo.labstack.com/guide/response/)

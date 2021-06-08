# 05 Request & Validate

## Retrieve Data

### Form Data

`example/api/request.go`

```go
func getForm(c echo.Context) error {
    name := c.FormValue("name")
    reply := echo.Map{
        "data": echo.Map{"name": name},
    }
    return c.JSON(http.StatusOK, reply)
}
```

**Try it:**

```powershell
curl -X POST "http://localhost:8080/request/form" `
-d "name=Joe"
```

```text
{"data":{"name":"Joe"}}
```

### Query Parameters

`example/api/request.go`

```go
func getQueryParam(c echo.Context) error {
    name := c.QueryParam("name")
    reply := echo.Map{
        "data": echo.Map{"name": name},
    }
    return c.JSON(http.StatusOK, reply)
}
```

**Try it:**

```powershell
curl -X GET "http://localhost:8080/request/query?name=Joe"
```

```text
{"data":{"name":"Joe"}}
```

### Path Parameters

`example/api/request.go`

```go
func getPathParam(c echo.Context) error {
    name := c.Param("name")
    reply := echo.Map{
        "data": echo.Map{"name": name},
    }
    return c.JSON(http.StatusOK, reply)
}
```

**Try it:**

```powershell
curl -X GET "http://localhost:8080/request/path/Joe"
```

```text
{"data":{"name":"Joe"}}
```

You can see the difference of handler registering among them.

`example/api/request.go`

```go
func SetRequestHandler(g *echo.Group) {
    g.POST("/form", getForm)
    g.GET("/query", getQueryParam)
    g.GET("/path/:name", getPathParam)
    g.POST("/validate", validateForm)
}
```

### See also

[01 Binding](./01binding.md)

## Validation

Echo doesn’t have built-in data validation capabilities, however, you can register a custom validator by implement the validator interface.

Example uses [github.com/go-playground/validator](https://github.com/go-playground/validator) validation framework

`example/api/validator.go`

```go
import (
    "net/http"

    "github.com/go-playground/validator"
    "github.com/labstack/echo/v4"
)

type CustomValidator struct {
    validator *validator.Validate
}

func (cv *CustomValidator) Validate(i interface{}) error {
    if err := cv.validator.Struct(i); err != nil {
        return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
    }
    return nil
}

func SetValidator(e *echo.Echo) {
    e.Validator = &CustomValidator{validator: validator.New()}
}
```

Handler:

`example/api/request.go`

```go
func validateForm(c echo.Context) (err error) {
    form := &recieveUser{}
    if err = c.Bind(form); err != nil {
        return echo.NewHTTPError(http.StatusBadRequest, err.Error())
    }
    if err = c.Validate(form); err != nil {
        return echo.NewHTTPError(http.StatusBadRequest, err.Error())
    }
    return c.JSON(http.StatusOK, echo.Map{"data": form})
}
```

**Try it:**

Valid Request:

```powershell
curl -X POST "http://localhost:8080/request/validate" `
-d "name=Joe" `
-d "email=joe@labstack.com"
```

```text
{"data":{"name":"Joe","email":"joe@labstack.com"}}
```

Bad Request:

```powershell
curl -X POST "http://localhost:8080/request/validate" `
-d "email=joe@labstack.com"
```

```text
{"message":"code=500, message=Key: 'recieveUser.Name' Error:Field validation for 'Name' failed on the 'required' tag"}
```

```powershell
curl -X POST "http://localhost:8080/request/validate" `
-d "name=Joe" `
-d "email=something_not_email"
```

```text
{"message":"code=500, message=Key: 'recieveUser.Email' Error:Field validation for 'Email' failed on the 'email' tag"}
```

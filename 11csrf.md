# 11 CSRF(Cross Site Request Forgery)

What is CSRF?

- [讓我們來談談 CSRF](https://blog.techbridge.cc/2017/02/25/csrf-introduction/)
- [OWASP: CSRF](https://owasp.org/www-community/attacks/csrf)

## Example

Use csrf middleware, the `TokenLookup` is the way to examine csrf token, here is form.

`example/api/csrf.go`

```go
func SetCsrfHandler(g *echo.Group) {
    g.Use(middleware.CSRFWithConfig(middleware.CSRFConfig{
        TokenLookup: "form:csrf",
    }))
    g.GET("", showCsrfExample)
    g.POST("", doCsrfExample)
}
```

Show the form when method is `GET`:

`example/api/csrf.go`

```go
func showCsrfExample(c echo.Context) error {
    csrf, ok := c.Get(middleware.DefaultCSRFConfig.ContextKey).(string)
    if !ok {
        return echo.NewHTTPError(http.StatusInternalServerError, "convert not expected type")
    }
    data := struct {
        Csrf string
    }{
        Csrf: csrf,
    }
    return c.Render(http.StatusOK, "csrf", data)
}
```

The template, note that the form must have csrf hidden field, we will fill it when rendering template.

`example/templates/csrf.html`

```html
{{ define "csrf" }}
<form method="POST" action="/csrf" >
    <input type="hidden" name="csrf" value="{{ .Csrf }}">
    <input type="email" id="email" name="email" required>
    <input type="submit" value="submit"/>
</form>
{{ end }}
```

The handler that get form data:

`example/api/csrf.go`

```go
func doCsrfExample(c echo.Context) error {
    email := c.FormValue("email")
    reply := echo.Map{
        "data": echo.Map{"email": email},
    }
    return c.JSON(http.StatusOK, reply)
}
```

**Try it:**

[Use Browser](http://localhost:8080/csrf)

And if using curl without giving form data, you will see error message

```powershell
curl -X POST "http://localhost:8080/csrf"
```

```text
{"message":"missing csrf token in the form parameter"}
```

## Reference

- [讓我們來談談 CSRF](https://blog.techbridge.cc/2017/02/25/csrf-introduction/)
- [echo doc](https://echo.labstack.com/middleware/csrf/)

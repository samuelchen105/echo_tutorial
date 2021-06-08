# 02 Custom Context

`echo.Context` represents the context of the current HTTP request. It contain all the information of the request including what standard `http.ResponseWriter` and `http.Request` would have.

As `echo.Context` is a interface, so we can extend it easily

**Example:**

Define a custom context

`example/api/context.go`

```go
type CustomContext struct {
    echo.Context
}

func (c *CustomContext) SendHello() error {
    return c.JSON(http.StatusOK, echo.Map{"msg": "hello, I am custom context"})
}

func ccMiddleware(next echo.HandlerFunc) echo.HandlerFunc {
    return func(c echo.Context) error {
        cc := &CustomContext{c}
        return next(cc)
    }
}
```

Handler:

`example/api/context.go`

```go
func SetContextHandler(g *echo.Group) {
    g.Use(ccMiddleware)
    g.GET("", contextExample)
}
```

```go
func contextExample(c echo.Context) error {
    cc := c.(*CustomContext)
    return cc.SendHello()
}
```

**Try it:**

```powershell
curl "http://localhost:8080/context"
```

```text
{"msg":"hello, I am custom context"}
```

> In the example, we only use the custom context in /context group

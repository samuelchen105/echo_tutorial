# 07 Route

- Routes can be registered by specifying HTTP method, path and a matching handler ex:`e.GET("/hello", hello)`
- Handlers in Echo must be define as `func(echo.Context) error`

## Match-any

Matches zero or more characters in the path. For example, pattern `/users/*` will match:

- `/users/`
- `/users/1`
- `/users/1/files/1`
- `/users/anything...`

## Path Matching Order

- Static
- Param
- Match any

```go
e.GET("/users/:id", func(c echo.Context) error {
    return c.String(http.StatusOK, "/users/:id")
})

e.GET("/users/new", func(c echo.Context) error {
    return c.String(http.StatusOK, "/users/new")
})

e.GET("/users/1/files/*", func(c echo.Context) error {
    return c.String(http.StatusOK, "/users/1/files/*")
})
```

Above routes would resolve in the following order:

1. `/users/new`
1. `/users/:id`
1. `/users/1/files/*`

## Group

`Echo#Group(prefix string, m ...Middleware) *Group`

- Group function will return a sub-router with optional middleware
- Sub-router also inherits parant middleware

Let's take 02 Custom Context as a example

`example/api/handler.go`

```go
SetContextHandler(e.Group("/context"))
```

`example/api/context.go`

```go
func SetContextHandler(g *echo.Group) {
    g.Use(ccMiddleware)
    g.GET("", contextExample)
}
```

1. Define a sub-router with prefix `"/context"`
2. Use the middleware to make custom context work, it will only be used in `"/context"` sub-router
3. Register handler `contextExample`, its entry point is `/context` because the path it registered is `""`

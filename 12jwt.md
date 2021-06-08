# 12 JWT(JSON Web Token)

- JWT is an open standard for securely transmitting information between parties as a JSON object.
- JWT is usually used for **Authorization** and **Information Exchange**

For more information:

- [jwt.io](https://jwt.io/introduction)
- [通俗易懂版講解JWT和OAuth2，以及他倆的區別和聯絡](https://www.itread01.com/content/1542396010.html)

## Echo JWT Middleware

- For valid token, it sets the user in context and calls next handler.
- For invalid token, it sends “401 - Unauthorized” response.
- For missing or invalid Authorization header, it sends “400 - Bad Request”.

## Example

The login handler function, create a JWT token and response.

`example/api/jwt.go`

```go
func jwtLogin(c echo.Context) error {
    username := c.FormValue("username")
    password := c.FormValue("password")

    //in real case, it shound search database
    if username != "admin" || password != "pwd" {
        return echo.ErrUnauthorized
    }

    token := jwt.New(jwt.SigningMethodHS256)

    claims, ok := token.Claims.(jwt.MapClaims)
    if !ok {
        return echo.NewHTTPError(http.StatusInternalServerError, "convert not expected type")
    }
    claims["name"] = "admin"
    claims["exp"] = time.Now().Add(time.Hour * 1).Unix()

    t, err := token.SignedString([]byte("secret"))
    if err != nil {
        return err
    }

    return c.JSON(http.StatusOK, map[string]string{
        "token": t,
    })
}
```

Get JWT token from context and response data.

`example/api/jwt.go`

```go
func jwtGetData(c echo.Context) error {
    user, ok := c.Get(middleware.DefaultJWTConfig.ContextKey).(*jwt.Token)
    if !ok {
        return echo.NewHTTPError(http.StatusInternalServerError, "convert not expected type")
    }
    claims, ok := user.Claims.(jwt.MapClaims)
    if !ok {
        return echo.NewHTTPError(http.StatusInternalServerError, "convert not expected type")
    }
    msg := fmt.Sprintf("Welcome %s", claims["name"])
    return c.String(http.StatusOK, msg)
}
```

Register JWT middleware and handler

`example/api/jwt.go`

```go
func SetJwtHandler(g *echo.Group) {
    g.POST("/login", jwtLogin)
    sg := g.Group("/restricted")
    sg.Use(middleware.JWTWithConfig(middleware.JWTConfig{
        SigningKey: []byte("secret"),
    }))
    sg.GET("", jwtGetData)
}
```

Note that we only use jwt middleware in `"/jwt/restricted"` sub-router, so the path `/jwt/restricted` need token to auth and `/jwt/login` doesn't.

**Try it:**

Login to get JWT token:

```powershell
curl -X POST "http://localhost:8080/jwt/login" `
-d "username=admin" `
-d "password=pwd"
```

```text
{"token":"eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE2MjMzMTc1MzYsIm5hbWUiOiJ0ZXN0bWFuIn0.Gcx6y2pPAn_IO2hnjQYoMI7VIUstinAAleABfhSkY0A"}
```

Use the token to get data:

```powershell
curl "http://localhost:8080/jwt/restricted" -H "Authorization: Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE2MjMzMTc1MzYsIm5hbWUiOiJ0ZXN0bWFuIn0.Gcx6y2pPAn_IO2hnjQYoMI7VIUstinAAleABfhSkY0A"
# format: change {TOKEN} to what you get
curl "http://localhost:8080/jwt/restricted" -H "Authorization: Bearer {TOKEN}"
```

```text
Welcome admin
```

Bad request:

```powershell
curl "http://localhost:8080/jwt/restricted"
```

```text
{"message":"missing or malformed jwt"}
```

## Reference

- [echo cookbook: jwt](https://echo.labstack.com/cookbook/jwt/)
- [echo doc: jwt](https://echo.labstack.com/middleware/jwt/)

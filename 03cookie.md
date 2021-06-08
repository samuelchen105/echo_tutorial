# 03 Cookie

Echo uses go standard `http.Cookie` object to add/retrieve cookies from the context which received in the handler function.

## Create Cookie

`example/api/cookie.go`

```go
func createCookie(c echo.Context) error {
    cookie := &http.Cookie{
        Name:    "username",
        Value:   "Jack",
        Expires: time.Now().Add(24 * time.Hour),
    }
    c.SetCookie(cookie)
    log.Println("set cookie: ", cookie)
    return c.JSON(http.StatusOK, echo.Map{"msg": "write a cookie"})
}
```

- `c.SetCookie(cookie)` will add a `Set-Cookie` header in HTTP response.

**Try it:**

```powershell
curl -X POST "http://localhost:8080/cookie/create"
```

```text
{"msg":"write a cookie"}
```

In server, you can see log that print the cookie

```text
2021/06/08 15:02:12 set cookie:  username=Jack; Expires=Wed, 09 Jun 2021 07:02:12 GMT
```

## Read Cookie

`example/api/cookie.go`

```go
func readCookie(c echo.Context) error {
    cookie, err := c.Cookie("username")
    if err != nil {
        return echo.NewHTTPError(http.StatusInternalServerError, err)
    }
    log.Println(cookie)
    return c.JSON(http.StatusOK, echo.Map{"msg": "read a cookie"})
}
```

**Try it:**

```powershell
curl -X GET "http://localhost:8080/cookie/read" `
--cookie "username=Jack"
```

```text
{"msg":"read a cookie"}
```

The log

```text
2021/06/08 15:04:22 username=Jack
```

## Read All Cookies

`example/api/cookie.go`

```go
func readAllCookies(c echo.Context) error {
    for _, cookie := range c.Cookies() {
        log.Println(cookie)
    }
    return c.JSON(http.StatusOK, echo.Map{"msg": "read all the cookies"})
}
```

**Try it:**

```powershell
curl -X GET "http://localhost:8080/cookie/readall" `
--cookie "username=Jack;email=Jack@gmail.com"
```

```text
{"msg":"read all the cookies"}
```

The log

```text
2021/06/08 15:11:11 username=Jack
2021/06/08 15:11:11 email=Jack@gmail.com
```

## See Also

- `http.Cookie` [definition](https://golang.org/pkg/net/http/#Cookie)

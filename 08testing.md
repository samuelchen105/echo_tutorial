# 08 Testing

Take 05 Request & Validation code as a example

`example/api/request.go`

```go
func SetRequestHandler(g *echo.Group) {
    g.POST("/form", getForm)
    g.GET("/query", getQueryParam)
    g.GET("/path/:name", getPathParam)
    g.POST("/validate", validateForm)
}
```

`example/api/request.go`

```go
func getForm(c echo.Context) error {
    name := c.FormValue("name")
    reply := echo.Map{
        "data": echo.Map{"name": name},
    }
    return c.JSON(http.StatusOK, reply)
}

func getQueryParam(c echo.Context) error {
    name := c.QueryParam("name")
    reply := echo.Map{
        "data": echo.Map{"name": name},
    }
    return c.JSON(http.StatusOK, reply)
}

func getPathParam(c echo.Context) error {
    name := c.Param("name")
    reply := echo.Map{
        "data": echo.Map{"name": name},
    }
    return c.JSON(http.StatusOK, reply)
}
```

Tests:

`example/api/request_test.go`

`TestGetForm`

```go
func TestGetForm(t *testing.T) {
    assert := assert.New(t)
    e := echo.New()
    var err error
    //test data
    name := "Jack"
    expected := map[string]interface{}{
        "data": map[string]interface{}{"name": name},
    }
    //setup form
    form := &url.Values{}
    form.Set("name", name)
    //setup request
    req := httptest.NewRequest(http.MethodPost, "/", strings.NewReader(form.Encode()))
    req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationForm)
    //setup context
    rec := httptest.NewRecorder()
    ctx := e.NewContext(req, rec)
    //assert
    assert.NoError(getForm(ctx))
    assert.Equal(http.StatusOK, rec.Code)
    var reply map[string]interface{}
    err = json.Unmarshal(rec.Body.Bytes(), &reply)
    assert.NoError(err)
    assert.Equal(expected, reply)
}
```

`TestGetQueryParam`

```go
//setup query
query := &url.Values{}
query.Set("name", name)
//setup request
req := httptest.NewRequest(http.MethodGet, "/?"+query.Encode(), nil)
req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationForm)
//setup context
rec := httptest.NewRecorder()
ctx := e.NewContext(req, rec)
//assert
assert.NoError(getQueryParam(ctx))
```

`TestGetPathParam`

```go
//setup request
req := httptest.NewRequest(http.MethodGet, "/", nil)
req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationForm)
//setup context
rec := httptest.NewRecorder()
ctx := e.NewContext(req, rec)
ctx.SetParamNames("name")
ctx.SetParamValues(name)
//assert
assert.NoError(getPathParam(ctx))
```

You can see the difference among them.

**Run all tests:**

```powershell
go test ./...
```

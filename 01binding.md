# 01 Binding

This part is about binding data from source to structure, if you want to retrieve single data by name, see [05 Request & Validation](./05request_and_validation.md).

The struct fields for binding data should have **tags**

```go
type User struct {
    ID string `param:"id" query:"id" form:"id" json:"id" xml:"id"`
}
```

**Notes:**

- For `query`, `param`, `form` **only fields with tags** are bound.
- For `json` and `xml` can bind to public fields without tags but this is by their standard library implementation.
- Each step can overwrite bound fields from the previous step.
- To avoid security flaws, don't use bound struct directly

**Example:**
`example/api/schema.go`

```go
type recieveUser struct {
    Name  string `json:"name" form:"name" validate:"required"`
    Email string `json:"email" form:"email" validate:"required,email"`
}

type user struct {
    Name    string `json:"name" xml:"name"`
    Email   string `json:"email" xml:"email"`
    IsAdmin bool   `json:"is_admin" xml:"is_admin"`
}
```

> Ignore the validate tag, we will mention it in other part.

`example/api/binding.go`

```go
// POST /bind/basic
func useBind(c echo.Context) error {
    form := &recieveUser{}
    if err := c.Bind(form); err != nil {
        return echo.NewHTTPError(http.StatusBadRequest, err.Error())
    }

    u := &user{
        Name:    form.Name,
        Email:   form.Email,
        IsAdmin: false,
    }

    doSomeBusinessLogic(u)

    return c.JSON(http.StatusOK, echo.Map{"data": form})
}
```

**Try it:**

```powershell
curl -X POST "http://localhost:8080/bind/basic" `
    -d "name=Joe" `
    -d "email=joe@labstack.com"
```

```text
{"data":{"name":"Joe","email":"joe@labstack.com"}}
```

## Use Binder

Rcho provides some helper functions for binding

- `echo.QueryParamsBinder(c)` - binds query parameters (source URL)
- `echo.PathParamsBinder(c)` - binds path parameters (source URL)
- `echo.FormFieldBinder(c)` - binds form fields (source URL + body). See also Request.ParseForm.

methods:

- `<Type>("param", &destination)` - if parameter value exists then binds it to given destination of that type.
- `Must<Type>("param", &destination)` - parameter value is required to exist.
- `<Type>s("param", &destination)` - for slices
- `Must<Type>s("param", &destination)` - for slices
- `BindError()` returns the first bind error from binder and resets all errors in this binder.
- `BindErrors()` returns all bind errors from binder and resets errors in binder.

**Example:**

`example/api/binding.go`

```go
// url =  "/bind/binder?active=true&id=1&id=2&id=3&length=25"
func useBinder(c echo.Context) error {
    var opts struct {
        IDs    []int64
        Active bool
        Length int64
    }

    // creates query params binder that stops binding at first error
    err := echo.QueryParamsBinder(c).
        Int64("length", &opts.Length).
        Int64s("id", &opts.IDs).
        Bool("active", &opts.Active).
        BindError() // returns first binding error

    if err != nil {
        return echo.NewHTTPError(http.StatusBadRequest, err.Error())
    }
    return c.JSON(http.StatusOK, echo.Map{"data": &opts})
}
```

**Try it:**

```powershell
curl "http://localhost:8080/bind/binder?active=true&id=1&id=2&id=3&length=25"
```

```text
{"data":{"IDs":[1,2,3],"Active":true,"Length":25}}
```

# 09 Casbin Auth

Casbin is a open-source access control library. It provides support for enforcing authorization based on various models.

Such as:

- ACL(Access Control List)
- RBAC (Role-Based Access Control)
- ABAC (Attribute-Based Access Control)

The example below uses the ACL

The casbin enforcer need two config files:

- `model.conf`: define how to authorize
- `policy.csv`: list rules

`example/configs/casbin_policy.csv`

```csv
p, admin, /casbin/resource, GET
```

This file will contain multiple rows, and each row will be `p, sub, obj, act` to define a rule. The above rule means that `admin` can access `/casbin/resource` by `GET` method.

`example/configs/casbin_model.conf`

```conf
[request_definition]
r = sub, obj, act

[policy_definition]
p = sub, obj, act

[policy_effect]
e = some(where (p.eft == allow))

[matchers]
m = r.sub == p.sub && r.obj == p.obj && r.act == p.act
```

- Request Definition:
this defines the arguments to pass into the `.Enforce(...)` function.

- Policy definition:
this binds our policy rules from the policy.csv file.

- Policy effect:
this is how to evaluate overall result, `e = some(where (p.eft == allow))` means that if any of the rules is allowed, then the overall result is allowed.

- Policy matches:
this defines how to evaluate the result that check a rule is allowed or not.

## Example
Implement custom enforcer:
`example/api/casbin.go`

```go
type Enforcer struct {
    enforcer *casbin.Enforcer
}

func (e *Enforcer) Enforce(next echo.HandlerFunc) echo.HandlerFunc {
    return func(c echo.Context) error {
        // this is a simple example, there are better ways to examine login
        cookie, err := c.Cookie("user")
        if err != nil {
            return echo.ErrForbidden
        }
        user := cookie.Value
        path := c.Request().URL.Path
        method := c.Request().Method

        isAllowed, err := e.enforcer.Enforce(user, path, method)
        if err != nil {
            return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
        }

        if isAllowed {
            return next(c)
        }
        return echo.ErrForbidden
    }
}
```

Handler:
`example/api/casbin.go`

```go
func SetCasbinHandler(g *echo.Group) error {
    enforcer, err := casbin.NewEnforcer(casbinModelPath, casbinConfigPath)
    if err != nil {
        return err
    }
    ce := Enforcer{enforcer: enforcer}
    g.Use(ce.Enforce)
    g.GET("/resource", casbinExample)
    return nil
}
```

```go
func casbinExample(c echo.Context) error {
    return c.JSON(http.StatusOK, echo.Map{"data": "resource"})
}
```

**Try it:**

Allowed:

```powershell
curl -X GET "http://localhost:8080/casbin/resource" `
--cookie "user=admin"\
```

```text
{"data":"resource"}
```

Forbidden:

```powershell
curl -X GET "http://localhost:8080/casbin/resource"
```

```text
{"message":"Forbidden"}
```

## Reference

- [casbin github](https://github.com/casbin/casbin)
- [casbin doc](https://casbin.org/docs/en/tutorials)

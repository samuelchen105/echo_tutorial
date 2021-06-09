# Echo Tutorial

## Installation

```powershell
go mod init echo/example
go get -u github.com/labstack/echo/v4
go mod tidy
```

Or you can just import it, initial module, and use `go mod tidy`. It will be downloaded automatically.

### Use the example

```powershell
git clone https://github.com/yuhsuan105/echo_tutorial.git
cd echo_tutorial/example
go run main.go
````

After server start, you can use these tools to send request and see the result.

- `curl`

```powershell
curl -X GET "http://localhost:8080/"
```

- browser

- postman

In most example, I will use `curl`

## Basic

- [01 Binding](01binding.md)
- [02 Custom Context](02custom_context.md)
- [03 Cookie](03cookie.md)
- [04 IP Address](04ip_address.md)
- [05 Request & Validation](05request_and_validation.md)
- [06 Response](06response.md)
- [07 Route](07route.md)
- [08 Testing](08testing.md)

## Advanced

- [09 Casbin Auth](09casbin_auth.md)
- [10 CORS](10cors.md)
- [11 CSRF](11csrf.md)
- [12 JWT](12jwt.md)
- [13 Proxy](13proxy.md)

## Reference

- [offical document](https://echo.labstack.com/guide/)

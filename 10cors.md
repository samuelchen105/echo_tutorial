# 10 CORS(Cross-Origin Resource Sharing)

For the secure reason, browsers will follow the same-origin policy and restrict cross-origin HTTP request. Unless the cross-origin request and response meet the CORS policy.

## Same-origin Policy

Two URLs have the same origin if the following are the same for both

- protocol (`http`, `https`)
- port (`:8080`)
- host (`localhost`)

## Simple Request

A simple request is one that meets all the following conditions:

- One of the allowed method
  - `GET`
  - `HEAD`
  - `POST`
- One of the allowed header
  - `Accept`
  - `Accept-Language`
  - `Content-Language`
  - `Content-Type` (partial)
- The only allowed values for the `Content-Type`
  - `application/x-www-form-urlencoded`
  - `multipart/form-data`
  - `text/plain`

The simple request doesn't need to "preflight".

The exchange:

1. Client send a request with `Origin` header
2. Server receive the request and determine whether the request is allowed
3. If allowed, server send a response with `Access-Control-Allow-Origin` header

The `Access-Control-Allow-Origin` header can also be `*` which means that the server can be accessed by **any** origin

## Preflighted requests

The not-simple request need to **preflight** before sending the request.

The exchange:

1. Client send a preflight request:
    - Method is `OPTIONS`
    - `Origin`: the origin of main request
    - `Access-Control-Allow-Methods`: the method of main request
    - `Access-Control-Allow-Headers`: the header of main request
2. If the method and header of the main request is allowed, server will response with:
    - `Access-Control-Allow-Origin`: allowed orgins
    - `Access-Control-Allow-Methods`: allowed methods
    - `Access-Control-Allow-Headers`: allowed headers
    - `Access-Control-Max-Age`: how long the response to the preflight request can be cached
3. Client send main request

## Example

Set option and use CORS middleware

`example/api/cors.go`

```go
func SetCorsHandler(g *echo.Group) {
  g.Use(middleware.CORSWithConfig(middleware.CORSConfig{
    AllowOrigins: []string{"https://example.com"},
    AllowHeaders: []string{echo.HeaderOrigin},
  }))
  g.GET("", corsExample)
}
```

> `corsExample` is just response a string

**Try it:**

```powershell
curl -H "Origin: https://example.com" `
--verbose `
"http://localhost:8080/cors"
```

The response will look like this:

```text
HTTP/1.1 200 OK
Access-Control-Allow-Origin: https://example.com
Content-Type: text/plain; charset=UTF-8
```

## Reference

- [MDN: CORS](https://developer.mozilla.org/en-US/docs/Web/HTTP/CORS)
- [MDN: Same-origin policy](https://developer.mozilla.org/en-US/docs/Web/Security/Same-origin_policy)

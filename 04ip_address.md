# 04 IP Address

Echo provides the `Echo#IPExtractor` method to define how retrieving ip that request from

- In default behavior, Echo sees all of first XFF header, X-Real-IP header and IP from network layer.
- Any HTTP header is untrustable because the clients have full control what headers to be set.
- The default behavior is not good

To decide what method to use, you need to know whether the application uses proxies or not

## 1. With no proxy

```go
e.IPExtractor = echo.ExtractIPDirect()
```

## 2. With proxies using `X-Forwarded-For` header

```go
e.IPExtractor = echo.ExtractIPFromXFFHeader()
```

See also: [TrustOptions](https://godoc.org/github.com/labstack/echo#TrustOption)

## With proxies using `X-Real-IP` header

```go
e.IPExtractor = echo.ExtractIPFromRealIPHeader()
```

See also: [TrustOptions](https://godoc.org/github.com/labstack/echo#TrustOption)

## Reference

- [offical doc](https://echo.labstack.com/guide/ip-address/)

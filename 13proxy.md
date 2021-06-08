# Proxy

Echo `middleware.Proxy` provides an HTTP/WebSocket reverse proxy middleware.

## Reverse Proxy

- Can provide single entry point to multiple backend servers
- Hide the backend servers, client can only know the proxy
- Protect against common web-based attacks, it can filter out unwanted requests before forwarding
- Load banlancing
- [Url rewrite](https://zh.wikipedia.org/wiki/URL%E9%87%8D%E5%AF%AB)
- De-coupling
For example, in microservice architecture, if all the service need authorization, the proxy can access the authorization service first and then determine whether forwarding the request or not, instead of using middleware in every services. So that seperating authorization from other services.

## Example

`proxy_example/main.go`

```go
var (
    port       = "8081"
    serviceUrl = "http://localhost:8080"
)

func main() {
    e := echo.New()

    // Setup proxy
    url1, err := url.Parse(serviceUrl)
    if err != nil {
        e.Logger.Fatal(err)
    }

    e.Use(middleware.Proxy(middleware.NewRoundRobinBalancer([]*middleware.ProxyTarget{
        {
            URL: url1,
        },
    })))

    e.Logger.Fatal(e.Start(":" + port))
}
```


**Try it:**

After run example, open another terminal

```powershell
# enter proxy_example
cd proxy_example
# run it
go run main.go
```

```powershell
curl "http://localhost:8081/"
```

And you can see the helloworld string what `http://localhost:8080/` will show.

```text
Hello, World!
```

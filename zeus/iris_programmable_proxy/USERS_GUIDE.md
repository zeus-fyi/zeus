Using the Programmable Proxy (Detailed Overview)

```go
package iris_programmable_proxy

const IrisEndpoint = "https://iris.zeus.fyi"

/*
    You'll use the API bearer token that you generate from the Access panel to authenticate with the load balancer.
 */

IrisClientProd = Iris{
    resty_base.GetBaseRestyClient("https://iris.zeus.fyi", tc.Bearer),
}

/*
    You then use the name of your route table group as a query parameter like the below,
    and it will round-robin the requests between the endpoints in that group table. 
 */

routeGroup := "quicknode-mainnet"

Add HEADER "X-Route-Group" with value "quicknode-mainnet"
path := "https://iris.zeus.fyi/v1/router"
```

```text
Supported HTTP Methods

GET, POST, PUT, DELETE

The proxy will pass through & respond according to these rules:
 
 1. The HTTP method that you use in your request
    GET requests -> GET
    POST requests -> POST
    PUT requests -> PUT
    DELETE requests -> DELETE
    
 2. Your request headers (minus your iris service & auth headers).
    
    You will need to pass a header called: "X-Route-Group", which matches at least one of your routing groups.
    
 3. Your request body (if you have one).
 4. Your request query parameters (if you have any).

       Eg. GET: https://iris.zeus.fyi/v1/router/items/?skip=0&limit=10
       
       Iris will translate this and send the request as
       
       https://{selected-route}/items/?skip=0&limit=10

        ...the query parameters are:
        
        skip: with a value of 0
        limit: with a value of 10

 5. Your extended route path
    a. Here's an example of a request path that you might send to the proxy:
        
       HEADER: X-Route-Group: avalanche-mainnet
       https://iris.zeus.fyi/v1/router/
        
       Let's say this routing group has these two endpoints:
       https://avalanche-avalanche-mainnet.sandbox.quiknode.net/2f568e4df78544629ce9af64bbe3cef9145895f5/
       https://avalanche-avalanche-mainnet.sandbox.quiknode.net/8d568e4df78544629ce9af64bbe3cef9145895f8/
       
       Avalanche has three route paths: P-Chain, X-Chain, and C-Chain.
       
       /ext/bc/P
       /ext/bc/C/rpc
       /ext/bc/X
       
       To use an extended route you would append it to the end of the proxy path as if it were the base node url:
       
       Eg. POST: https://iris.zeus.fyi/v1/router/ext/bc/P
       will send the request to the P-Chain route for the node-url selected by the proxy.
       
       Iris will translate & send the request to this path
       
       POST: https://avalanche-avalanche-mainnet.sandbox.quiknode.net/2f568e4df78544629ce9af64bbe3cef9145895f5/ext/bc/P
       
 6. The proxy will return the response from the node to you.
 
 7. The proxy will return the response headers from the node to you.
    
    Additionally, Iris will return a response header: "X-Selected-Route" with the value of the route that was used for the request.
 
 8. If a response is received from the route it will use the response status code from the node as the response status code to you.
    
    E.g. If the node returns a 404, the proxy will return a 404 to you & the raw response body from the request.
```




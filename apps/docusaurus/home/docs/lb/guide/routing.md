---
sidebar_position: 2
---

# Routing

## Using the Programmable Proxy (Detailed Overview)

### Supported HTTP Methods

- **GET**
- **POST**
- **PUT**
- **DELETE**

### Proxy Rules

The proxy will respond based on the following rules:

1. **HTTP Method**: It will pass the request and respond according to the method used.
    - GET requests -> GET
    - POST requests -> POST
    - PUT requests -> PUT
    - DELETE requests -> DELETE


2. **Request Headers**: You will need to pass a header called "X-Route-Group" that matches at least one of your routing
   groups.

3. **Request Body**: If you have one.

4. **Request Query Parameters**: If you have any.
    - Example: `GET: https://iris.zeus.fyi/v1/router/items/?skip=0&limit=10`
        - Iris will translate this and send the request as `https://{selected-route}/items/?skip=0&limit=10`
        - The query parameters shown here are:
            - `skip`: with a value of 0
            - `limit`: with a value of 10

5. **Extended Route Path**: Here's an example of a request path that you might send to the proxy:
    - HEADER: X-Route-Group: avalanche-mainnet
    - URL: `https://iris.zeus.fyi/v1/router/`

   Let's say this routing group has these two endpoints:
    - `https://avalanche-avalanche-mainnet.sandbox.quiknode.net/2f568e4df78544629ce9af64bbe3cef9145895f5/`
    - `https://avalanche-avalanche-mainnet.sandbox.quiknode.net/8d568e4df78544629ce9af64bbe3cef9145895f8/`

   Avalanche has three route paths: P-Chain, X-Chain, and C-Chain.
    - `/ext/bc/P`
    - `/ext/bc/C/rpc`
    - `/ext/bc/X`

   To use an extended route you would append it to the end of the proxy path as if it were the base node url:
    - Eg. POST: `https://iris.zeus.fyi/v1/router/ext/bc/P` will send the request to the P-Chain route for the node-url
      selected by the proxy.
    - Iris will translate & send the request to this
      path: `POST: https://avalanche-avalanche-mainnet.sandbox.quiknode.net/2f568e4df78544629ce9af64bbe3cef9145895f5/ext/bc/P`

6. **Proxy Response**: It will return the response from the route.

7. **Response Headers**: It will return a response header called "X-Selected-Route" with the value of the route used.

8. **Response Status Code**: If a response is received from the route, the proxy will use the response status code as
   the response status code to you.
    - E.g. If the node returns a 404, the proxy will return a 404 to you & the raw response body from the request.


---
sidebar_position: 1
displayed_sidebar: zK8s
---

# Client #

The zK8s client is used for interacting with our cloud platform's core infrastructure apis and services.

## Quickstart ##

You can override the values in the test files we have and set to your own following these steps

1. In /test/configs -> create a config.yaml using the sample-config.yaml as a reference, the config.yaml should be in
   .gitignore by default so it doesn't commit your tokens
2. Add your bearer token to this config, otherwise config it directly in the client
3. Override the zeus_client_test variables that are used to point to your desired chart and kubernetes location
4. Then run the test, the zeus_client_integrated_test will upload and then query for the uploaded chart and then print
   it.

## Table of Contents

1. [Imports](#imports)
2. [Structures](#structures)
    - [ZeusClient](#zeusclient)
3. [Functions](#functions)
    - [NewZeusClient](#newzeusclient)
    - [NewDefaultZeusClient](#newdefaultzeusclient)

---

## Imports

```go
package zeus_client

import resty_base "github.com/zeus-fyi/zeus/zeus/z_client/base"

type ZeusClient struct {
   resty_base.Resty
}

func NewZeusClient(baseURL, bearer string) ZeusClient {
   z := ZeusClient{}
   z.Resty = resty_base.GetBaseRestyClient(baseURL, bearer)

   return z
}

const ZeusEndpoint = "https://api.zeus.fyi"

func NewDefaultZeusClient(bearer string) ZeusClient {
   return NewZeusClient(ZeusEndpoint, bearer)
}
```

---

## Structures

### `ZeusClient`

Provides a client for interacting with Zeus, wrapping the base resty client.

**Fields:**

- `Resty`: The underlying REST client.

---

## Functions

### `NewZeusClient(baseURL, bearer string) ZeusClient`

Creates a new Zeus client.

**Parameters:**

- `baseURL`: The base URL for the Zeus service.
- `bearer`: The bearer token for authentication.

**Returns:**

- A new `ZeusClient` instance.

### `NewDefaultZeusClient(bearer string) ZeusClient`

Creates a new default Zeus client with the predefined Zeus endpoint.

**Parameters:**

- `bearer`: The bearer token for authentication.

**Returns:**

- A new `ZeusClient` instance.

---


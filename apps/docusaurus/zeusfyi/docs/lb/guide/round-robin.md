---
sidebar_position: 3
---

# Round Robin

Our Redis powered load balancer will robin round requests which improves load distribution and reliability,
delivering a more consistent performance and resilience against traffic spikes, while also letting you easily
mix in your desired mix of clients for your nodes and RPC providers.

Curl example:

```sh
curl --location 'https://iris.zeus.fyi/v1/router' \
--header 'Content-Type: application/json' \
--header 'Authorization: Bearer YOUR-BEARER-TOKEN' \
--header ‘X-Load-Balancing-Strategy: RoundRobin’ \
--header 'X-Route-Group: ethereum-mainnet' \
--data '{
    "jsonrpc": "2.0",
    "method": "eth_getBlockByNumber",
    "params": ["latest", true],
    "id": 1
}'
```
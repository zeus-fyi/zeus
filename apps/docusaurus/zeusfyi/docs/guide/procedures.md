---
sidebar_position: 4

---

# Procedures

To use embedded protocol procedures you only need to add the key value to your payload. In this example, to use the max
block procedure for Ethereum, which polls your routing table for the current block number, and then forwards your
request to the endpoints returning the highest block number seen and then returns the first successful response.

"procedure": "eth_maxBlockAggReduce"

```
curl --location 'https://iris.zeus.fyi/v1/router' \
--header 'Content-Type: application/json' \
--header 'Authorization: Bearer YOUR-BEARER-TOKEN' \
--header 'X-Route-Group: ethereum-mainnet' \
--data '{
    "jsonrpc": "2.0",
    "procedure": "eth_maxBlockAggReduce",
    "method": "eth_getBlockByNumber",
    "params": ["latest", true],
    "id": 1
}'
```
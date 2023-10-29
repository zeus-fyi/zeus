---
sidebar_position: 4
---

# Adaptive

Once you have ~20 or so request samples for the same method, the load balancer will start to use the adaptive strategy
automatically and manage the routing group table for you based on the best predicted performing endpoint for that method
that's available.

Stats will only persist for one hour since the last API call for that method, so you'll need to keep making requests to
keep the stats. It doesn't take long, only ~20 samples per metric to trend towards a near optimal routing group table
from scratch, so it's really not a big deal to reset the stats.

You can tune the adaptive scale factors for your application, but the defaults are pretty good for most applications.

![Screenshot 2023-10-28 at 8 23 04 PM](https://github.com/zeus-fyi/olympus/assets/17446735/a4a152e6-201a-4863-9702-0aeaa9a10052)

```text
Add HEADER "X-Route-Group" with value "quicknode-mainnet"
Add HEADER "X-Load-Balancing-Strategy" with value "Adaptive"
Add HEADER "X-Adaptive-Metrics-Key" with value "JSON-RPC" (or "Other Metric Keys...")
```

JSON-RPC is a reserved value for json-rpc based POST api metrics, it collects stats by the method value in the json rpc
POST request

### Curl Example:

```sh
curl --location ‘https://iris.zeus.fyi/v1/router’ \
--header ‘Content-Type: application/json’ \
--header ‘Authorization: Bearer YOUR-BEARER-TOKEN’ \
--header ‘X-Route-Group: ethereum-mainnet’ \
--header ‘X-Load-Balancing-Strategy: Adaptive’ \
--header ‘X-Adaptive-Metrics-Key: JSON-RPC’ \
--data ‘{“jsonrpc”:“2.0”,“method”:“eth_getBlockByNumber”,“params”:[“latest”, true],“id”:1}’
```

You can also check out our round-robin load_balancing_test.go for an example of how to use the programmable proxy to
query the block number from a routing group of ethereum node urls endpoints.

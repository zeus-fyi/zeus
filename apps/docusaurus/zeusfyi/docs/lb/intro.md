---
sidebar_position: 1
displayed_sidebar: lb
---

# Setup

#### Exclusively running Iris services through the QuickNode Marketplace until our v1 release later this year.

## Overview

We use ZU to denote Zeus compute units.
These units are derived from server bandwidth, cpu, memory usage and average and peak traffic usage, and operational
costs for development & maintenance.

<b>4 ZU</b> per request & response.<br/>
<b>1 ZU</b> per 1 kB<br/><br/>
<b>10 ZU minimum for round trip request </b><br/><br/>

Each user can store up to 1000 endpoints for free, you can then use these endpoints to create routing group tables via
API or UI Dashboard

## QuickNode Marketplace Users

QuickNode marketplace customers will use an SSO link from QuickNode and be directed to their load balancing stored
endpoints
table, you'll then be able to generate an API key from the access panel.

### Additional Benefits:

QuickNode users will have their endpoints automatically registered with the load balancer, and will have automatically
generated routing group tables based on the network-chain type for that endpoint. E.g. ethereum mainnet endpoints will
be
automatically registered to the ethereum-mainnet routing group table. These won't count against your
routing table limits unless you make any manual changes to them.

Need more? Send us a message at support@zeus.fyi

## API Key Setup

This is a guide to help you set up your own programmable proxy for the Iris Load Balancer.
Prerequisites: You'll need to generate an API key from the access panel if you don't have an existing one.

![APIkeypage](https://github.com/zeus-fyi/zeus/assets/17446735/7352892d-49ad-4a72-add1-5b212a90b914)

### Using the Load Balancer

You'll use the API bearer token that you generate from the Access panel to authenticate with the load balancer.
You then use the name of your route table group as a query parameter like the below,
and it will default to round-robin the requests between the endpoints in that group table if you have a lite plan, and
adaptive for standard+ plans.

### Curl Example:

```shell
curl --location 'https://iris.zeus.fyi/v1/router' \
--header 'Content-Type: application/json' \
--header 'Authorization: Bearer YOUR-BEARER-TOKEN' \
--header 'X-Route-Group: quicknode-mainnet' \
--data '{"jsonrpc":"2.0","method":"eth_getBlockByNumber","params":["latest", true],"id":1}'
```
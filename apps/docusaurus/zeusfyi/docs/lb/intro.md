---
sidebar_position: 1
displayed_sidebar: lb
---

# Setup

## Top Reason to Use the Adaptive Load Balancer?

There's many, but the top one? People are sick of hitting 429 rate limiting and 5xx errors. What's
double worse is paying for the request that failed. So we set out to solve this problem.

Now there's a solution, backed by extensive studies. Turn many endpoints into one super endpoint that handles the
scale you need without the errors. It can handle Nx more requests with N being the number of
requests/sec than any single endpoint can handle. That's in addition to the other significant proven performance gains
and error rate reductions you can expect from our adaptive load balancing technology, see our benchmarking section for
details.

## Web3 Users Free Mempool Access

We're offering unlimited free access to our mempool service at a level which is comparable to:

- BloXroute - Enterprise/Enterprise-Elite ($1250-5000/mo)
- Blocknative - Growth1-Growth2 ($1250-5000/mo)


# Sign Up

You can use the free tier to get started, and then upgrade to a paid plan at any time, by setting a credit card on your
account and paying as you go, or by prepaying for a plan.
https://cloud.zeus.fyi/billing


## Overview

We use ZU to denote Zeus compute units.
These units are derived from server bandwidth, cpu, memory usage and average and peak traffic usage, and operational
costs for development & maintenance.

<b>4 ZU</b> per request & response.<br/>
<b>1 ZU</b> per 1 kB<br/><br/>
<b>10 ZU minimum for round trip request </b><br/><br/>

Each user can store up to 1000 endpoints for free, you can then use these endpoints to create routing group tables via
API or UI Dashboard

## 1-on-1 Onboarding

Want one-on-one help getting started? Schedule a Google meet with an expert
at [https://calendly.com/zeusfyi/solutions-engineering](https://calendly.com/zeusfyi/solutions-engineering)

## API Key Setup

This is a guide to help you set up your own programmable proxy for the Iris Load Balancer.
Prerequisites: You'll need to generate an API key from the access panel if you don't have an existing one.

![APIkeypage](https://github.com/zeus-fyi/zeus/assets/17446735/7352892d-49ad-4a72-add1-5b212a90b914)

### Using the Load Balancer

You'll use the API bearer token that you generate from the Access panel to authenticate with the load balancer.
You then use the name of your route table group as a query parameter like the below,
and it will default to adaptive, or you can specify round-robin with the `X-Load-Balancing-Strategy` header set
to `RoundRobin`.

```shell

### Curl Example:

```shell
curl --location 'https://iris.zeus.fyi/v1/router' \
--header 'Content-Type: application/json' \
--header 'Authorization: Bearer YOUR-BEARER-TOKEN' \
--header 'X-Route-Group: quicknode-mainnet' \
--data '{"jsonrpc":"2.0","method":"eth_getBlockByNumber","params":["latest", true],"id":1}'
```

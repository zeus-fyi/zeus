---
sidebar_position: 5

---

# Benchmarks

We ran a several weeks long test using a real production workload, and were delighted to see how much of an impact we
make on performance.

## Methodology

Study was conducted using a real production workload used for monitoring Uniswap prices

- We used 6 Enterprise level QuickNode Ethereum Mainnet endpoints for this test.
- We took 190k samples ran over several weeks using our Adaptive algorithm, using T-digest to calculate the median.
- We used a 20 sample round-robin sampling using T-digest.
- We used our max-block procedure to ensure that we were always using the most up-to-tip-of-chain endpoints for our
samples.

Read more about how T-Digest works: https://www.softwareimpacts.com/article/S2665-9638(20)30040-3/fulltext

### Assumptions

- We assumed that the endpoints are of equal quality
- We focused on highlighting eth_getBlockByNumber
  - Since it is least likely to be improved by vendor caching and thus more representative of the true performance
  - We used caching on ethGetBlockByNumber to ensure that we were not calling it more than once per block
- We use 20 samples for our Round Robin comparison
    - Since that is the minimum recommended for statistically significant accuracy
- T-digest has slightly less accurate medians with lower samples
    - We think that the error variance is captured within the 3% difference between our simulated results & actual

### Post-Analysis Findings

We found that if eth_getBlockByNumber was called an the block wasn't cached that it would take 400-600ms in most
cases, but if it was cached it would take 50-100ms. We also noticed that if a block value was called multiple times
even if it was a historical one, it appeared to cache the first request, so all subsequent calls were
significantly faster.

```json
{
  "jsonrpc": "2.0",
  "method": "eth_getBlockByNumber",
  "params": [
    "latest",
    true
  ],
  "id": 1
}
```

For this method the latency from using our load balancer which proxies the request adds ~100-200ms RTT to the total
latency.
This sample is using a direct connection to a QuickNode endpoint

```text
time taken:  396
time taken:  112
time taken:  107
time taken:  101
```

This sample is using our load balancer with a round-robin connection without server network latency

```text
time taken:  299
time taken:  284
time taken:  288
time taken:  281
```

## Adaptive Scale Factor Settings Used

- Priority Score: 0.52
- Error Score: 3
- Decay Rate: 0.95

We lowered the priority score to 0.52 from 0.6, since the adaptive tuning starts to create a dominant median, so
lowering the growth rate slightly keeps the Adaptive scores within a cyclical range instead of growing too fast.

### Round Robin

![small](https://github.com/zeus-fyi/zeus/assets/17446735/efccf2b0-ecc8-4bef-a966-e7fe994370a2)

### Adaptive

![weekly](https://github.com/zeus-fyi/zeus/assets/17446735/9919f53c-7b6a-46ba-9780-7fbbc0aa9da0)

## Results

```eth_getBlockByNumber```

- p50 median improved from 505ms -> 426ms
- p50 reduced by ~80ms or about 18% improvement.

Matching our initial prediction closely of ~15%

We also saw a significant reduction in our overall api requests needed for the same workload,
and thus consumed significantly less QuickNode compute units needed for the same workload.

Still think you don't need a load balancer?

## Next steps

- Using a better control group, and more endpoints to test with.
- Better understanding of how error rate impacts the adaptive algorithm.
- Testing more request types, and more workloads over archive, near tip, and tip of chain data.
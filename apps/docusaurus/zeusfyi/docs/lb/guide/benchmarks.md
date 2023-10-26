---
sidebar_position: 5

---

# Benchmarks

We ran a several weeks long test using a real production workload, and were delighted to see how much of an impact we
make on performance.

## Methodology

Study was conducted using a real production workload used for monitoring Uniswap prices

- We used 6 Enterprise level QuickNode Ethereum Mainnet endpoints for this test.
- We took 190k samples ran over several weeks using our Adaptive algorithm, using T-Digest to calculate the median.
- We used a 20 sample round-robin sampling using T-Digest.
- We used our max-block procedure to ensure that we were always using the most up-to-tip-of-chain endpoints for our
samples.

Read more about how T-Digest works: https://www.softwareimpacts.com/article/S2665-9638(20)30040-3/fulltext

### Assumptions

- We assumed that the endpoints are of equal quality
- We focused on highlighting eth_getBlockByNumber
  - Since it is least likely to be improved by vendor caching and thus more representative of the true performance
  - We used caching on ethGetBlockByNumber to reduce the impact of vendor caching downstream
- We use 20 samples for our Round Robin comparison
    - Since that is the minimum recommended for statistically significant accuracy
- T-Digest has slightly less accurate medians with lower samples
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

This sample is using a direct connection to a QuickNode endpoint, each subsequent request is faster due to caching,
so the first request is the most representative of the true latency.
```text
time taken:  396ms
time taken:  112ms
time taken:  107ms
time taken:  101ms
```

We determined that the load balancer adds about 20ms of latency to the request, for round-robin and adaptive

This sample is using our load balancer with a round-robin connection without server network latency
```text
time taken:  282ms
    - 26ms total load balancer latency
    - 256ms for raw request RTT
time taken:  256ms
    - 19ms total load balancer latency
    - 237ms for raw request RTT
time taken:  282ms
    - 21ms total load balancer latency
    - 261ms for raw request RTT
time taken: 255ms
    - 17ms total load balancer latency
    - 238ms for raw request RTT
```

This sample is using our load balancer inclusive of all total latency (raw request RTT + load balancer latency + server
proxy RTT)

![Adaptive](https://github.com/zeus-fyi/zeus/assets/17446735/d583ca5e-e742-4dfb-aab3-b305ef648798)

```text
total time taken:  659ms
    - 24ms load balancer latency
    - 443ms raw request RTT
    - 192ms server proxy RTT
```

We calculated the networking latency from proxying the request adds ~100-300ms RTT to the total latency.

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

P25 Improvement:

- Initial P25: 478ms
- New P25 after load balancer: 384ms
- Saving 478ms−384ms = <b>94ms</b>

P50 (Median) Improvement:

- Initial P50: 505ms
- New P50 after load balancer: 426ms
- Improvement: 505ms−426ms = <b>79ms</b>

P75 Improvement:

- Initial P75: 602ms
- New P75 after load balancer: 486ms
- Improvement: 602ms−486ms = <b>116ms</b>

P95 Improvement:

- Initial P95: 647ms
- New P95 after load balancer: 596ms
- Improvement: 647ms−596ms = <b>51ms</b>

Total time saved:

- 200k requests -> 4.58 hours saved
- 1M requests -> 24 hours time saved

We also saw a meaningful reduction in our overall api requests needed for the same workload,
and thus consumed less QuickNode compute units needed for the same workload. We still need to
do more analysis on this with a clean control group, but we think it is about 10-30% less requests needed.

Still think you don't need a load balancer?

## Note

We return this response header, which has the raw request RTT time, which you can also use for your own analysis.

- X-Response-Latency-Milliseconds

## Next steps

- Using a better control group, and more endpoints to test with.
- Better understanding of how error rate impacts the adaptive algorithm.
- Testing more request types, and more workloads over archive, near tip, and tip of chain data.
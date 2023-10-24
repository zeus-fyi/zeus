---
sidebar_position: 5

---

# Benchmarks

We ran a several weeks long test using a real production workload, and were delighted to see how much of an impact we
make on performance.

## Methodology

Study was conducted using a real production workload
We used 6 Enterprise level QuickNode Ethereum Mainnet endpoints for this test.
We took 190k samples ran over several weeks using our Adaptive algorithm, using t-digest to calculate the median.
We used a 20 sample round-robin sampling using t-digest.
We used our max-block procedure to ensure that we were always using the most up-to-tip-of-chain endpoints for our
samples.

### Assumptions

- We assumed that the endpoints are of equal quality
- We use 20 samples for our Round Robin comparison
    - Since that is the minimum recommended for statistically significant accuracy
- T-digest has slightly less accurate medians with lower samples
    - We think that the error variance is captured within the 3% difference between our simulated results & actual

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
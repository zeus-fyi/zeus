---
sidebar_position: 4
---

# Adaptive

### Overview

Once you have ~20 or so request samples for the same method, the load balancer will start to use the adaptive strategy
automatically and manage the routing group table for you based on the best predicted performing endpoint for that method
that's available.

Stats will only persist for one hour since the last API call for that method, so you'll need to keep making requests to
keep the stats. It doesn't take long, only ~20 samples per metric to trend towards a near optimal routing group table
from scratch, so it's really not a big deal to reset the stats.

You can tune the adaptive scale factors for your application, but the defaults are pretty good for most applications.

![Screenshot 2023-10-28 at 8 53 11 PM](https://github.com/zeus-fyi/zeus/assets/17446735/24d22cfb-c91a-4adf-a062-8e3dda2d8583)

### Usage Example

#### JSON-RPC

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

```sh
curl --location ‘https://iris.zeus.fyi/v1/router’ \
--header ‘Content-Type: application/json’ \
--header ‘Authorization: Bearer YOUR-BEARER-TOKEN’ \
--header ‘X-Route-Group: ethereum-mainnet’ \
--data ‘{“jsonrpc”:“2.0”,“method”:“eth_getBlockByNumber”,“params”:[“latest”, true],“id”:1}’
```

## Adaptive Score Simulator

We’re sharing this python script which we used to model simulations so you can copy, paste, and replicate our results or
tune your own workloads. Use an online IDE such as https://www.boot.dev/playground/py and you can start experimenting
right away. Try changing the random distribution function, injecting more routes to see how it performs across more
dynamic environments.

See results here:
https://docs.google.com/spreadsheets/d/1edd9iUSgmGraRqh3O1i8j1Xb7osjpUVD-e5UdkKrYII/edit?usp=sharing

Online Python IDE:
https://programiz.pro/ide/python/KTJ73ZZQX5?utm_source=python_playground-save-button

### Adaptive Tuning Parameters

```python
# These are your adaptive scoring parameters that you can tune
latency_scale_factor = 0.6
decay_scale_factor = 0.95
# This is error rate for each route
error_scale_factor = 3
```

### Adaptive Scenario Parameters

```python
# This is how many routes are in your table
table_routes_count = 2
# This is error rate for each route
route_a_error_rate = 0.01
route_b_error_rate = 0.01


# Used to set the latency distribution
def random_latency_percentile():
    return random.uniform(0, 1)
```

### How to Simulate A | B Testing

To simulate latency skew per endpoint you can modify the random_latency_percentile function and use

```python
def random_a_latency_percentile():
    return random.uniform(0, 1)


def random_b_latency_percentile():
    return random.uniform(0.1, 1)
````

Then you can select the route specific latency based on the endpoint selected

```python
if score_a < score_b:
    selected_endpoint = "A"
    latency = random_a_latency_percentile()
else:
    selected_endpoint = "B"
    latency = random_b_latency_percentile()
```

### Python Simulation Code

```python
import random

num_simulations = 1000
num_requests = 1000

latency_scale_factor = 0.6
decay_scale_factor = 0.95
error_scale_factor = 3
table_routes_count = 2

route_a_error_rate = 0.01
route_b_error_rate = 0.01


def random_latency_percentile():
    return random.uniform(0, 1)


def simulate(num_simulations=1000, num_requests=1000):
    results_a = []
    results_b = []

    for _ in range(num_simulations):
        score_a, score_b = 1.0, 1.0
        requests_a, requests_b = 0, 0

        for _ in range(num_requests):
            if score_a < score_b:
                selected_endpoint = "A"
            else:
                selected_endpoint = "B"

            latency = random_latency_percentile()

            if selected_endpoint == "A":
                score_a *= (latency + latency_scale_factor)
                if random.random() < route_a_error_rate:
                    score_a *= error_scale_factor
                requests_a += 1
                if requests_a % table_routes_count == 0:
                    score_a *= decay_scale_factor
            else:
                score_b *= (latency + latency_scale_factor)
                if random.random() < route_b_error_rate:
                    score_b *= error_scale_factor
                requests_b += 1
                if requests_b % table_routes_count == 0:
                    score_b *= decay_scale_factor

        results_a.append(requests_a)
        results_b.append(requests_b)

    # Calculate average and other statistics if needed
    avg_a = sum(results_a) / num_simulations
    avg_b = sum(results_b) / num_simulations

    return avg_a, avg_b


# After all simulations, you can calculate statistics on results_a and results_b
print(simulate())
```
## Adaptive Score Simulator

You can use the adaptive_load_sim.py script to simulate how your adaptive scoring will impact the
requests per endpoint.

See results here:
https://docs.google.com/spreadsheets/d/1edd9iUSgmGraRqh3O1i8j1Xb7osjpUVD-e5UdkKrYII/edit?usp=sharing

Online Python IDE:
https://programiz.pro/ide/python/KTJ73ZZQX5?utm_source=python_playground-save-button

```python
# These are your adaptive scoring parameters that you can tune
latency_scale_factor = 0.6
decay_scale_factor = 0.95
error_scale_factor = 3

# This is how many routes are in your table
table_routes_count = 2

# This is error rate for each route
route_a_error_rate = 0.01
route_b_error_rate = 0.01
```

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



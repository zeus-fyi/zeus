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


# After all simulations, you can calculate statistics on results_A and results_B
print(simulate())

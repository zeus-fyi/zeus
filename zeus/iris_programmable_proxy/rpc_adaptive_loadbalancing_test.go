package iris_programmable_proxy

import (
	"fmt"
	"strconv"
	"sync"

	iris_programmable_proxy_v1_beta "github.com/zeus-fyi/zeus/zeus/iris_programmable_proxy/v1beta"
)

func (t *IrisConfigTestSuite) TestAdaptiveRPCLoadBalancing() {
	groupName := "load-sim"
	t.IrisClientProd.SetRoutingGroupHeader(groupName)
	t.IrisClientProd.SetDebug(false)
	t.IrisClientProd.SetHeader(iris_programmable_proxy_v1_beta.LoadBalancingStrategy, iris_programmable_proxy_v1_beta.Adaptive)
	t.IrisClientProd.SetHeader(iris_programmable_proxy_v1_beta.AdaptiveLoadBalancingKey, iris_programmable_proxy_v1_beta.JsonRpcAdaptiveMetrics)
	t.IrisClientProd.SetHeader("X-Sim-Response-Format", "json")
	t.IrisClientProd.SetHeader("X-Sim-Response-Size", "1")

	payload := `{
		"jsonrpc": "2.0",
		"method": "eth_blockNumber",
		"params": [],
		"id": 1
	}`

	reqCountInParallel := 10
	// Define a channel for controlling the number of concurrent requests
	sem := make(chan bool, reqCountInParallel)

	// Define an error channel to catch errors from goroutines
	errCh := make(chan error, reqCountInParallel)

	var wg sync.WaitGroup

	offset := 0
	latencySums := make(map[string]int64)   // Sum of latencies for each endpoint
	samplesCounts := make(map[string]int64) // Count of samples for each endpoint
	errorCounts := make(map[string]int64)   // Count of samples for each endpoint

	mu := &sync.Mutex{} // Mutex to synchronize access to shared variables

	for i := 0; i < reqCountInParallel*100; i++ {
		// Acquire a semaphore
		sem <- true

		wg.Add(1)
		go func(offset int, latencySums map[string]int64, samplesCounts map[string]int64, errorCounts map[string]int64) {
			defer func() {
				// Release the semaphore
				<-sem
				wg.Done()
			}()

			resp, err := t.IrisClientProd.R().
				SetBody(payload).
				SetHeader("Content-Type", "application/json").
				Post("/v1/router/v1/load/simulate")

			t.Require().NoError(err)

			endpoint := resp.Header().Get(SelectedRouteResponseHeader)
			latencyStr := resp.Header().Get("X-Response-Latency-Milliseconds")
			latency, err := strconv.ParseInt(latencyStr, 10, 64)
			if err == nil {
				mu.Lock()
				if resp.StatusCode() >= 400 {
					errorCounts[endpoint]++
				}
				latencySums[endpoint] += latency // Increment total latency for the endpoint
				samplesCounts[endpoint]++        // Increment count of samples for the endpoint
				mu.Unlock()
			}
		}(offset, latencySums, samplesCounts, errorCounts)
	}

	offset++
	// Wait for all the goroutines to complete
	wg.Wait()
	close(errCh)

	// Check for any errors in the error channel
	for err := range errCh {
		t.NoError(err)
	}

	for endpoint, totalLatency := range latencySums {
		samplesCount := samplesCounts[endpoint]
		averageLatency := float64(totalLatency) / float64(samplesCount)
		errorCount := errorCounts[endpoint]
		fmt.Printf("Endpoint: %s | Average Latency: %.2f ms | Total Samples Count: %d | Total Errors Count: %d\n", endpoint, averageLatency, samplesCount, errorCount)
	}
}

/*
Test One: Using Load-Sim:

endpoint A was set to fail 12.5% of the time
adaptive LB -> reduced 62.5 expected failures to 42
			-> 6.25% round robin lb failure -> 4.2% adaptive lb failure
Endpoint A: Average Latency: 90.02 ms | Total Samples Count: 378 | Total Errors Count: 42
Endpoint B: Average Latency: 94.96 ms | Total Samples Count: 622 | Total Errors Count: 0
*/

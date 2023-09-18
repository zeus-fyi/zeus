# Load Simulator

The load simulator allows you to simulate various response sizes, formats, latencies, and success/failure scenarios by
sending specific HTTP headers with your request. The simulator supports GET, POST, PUT, and DELETE methods on
the `/v1/load/simulate` endpoint.

## Headers

You can control the simulator's behavior using the following HTTP headers:

### Response Size

- `X-Sim-Response-Size`: Specifies the response size. Defaults to `0` if not set.
- `X-Sim-Response-Size-Units`: Defines the units for response size. Only supports `KiB` and `MiB` for now. Defaults
  to `KiB` if not set.

### Response Format

- `X-Sim-Response-Format`: Defines the response format. Supports `bytes`, `string`, and `json`. Defaults to `string` if
  not set.

### Response Status Codes

- `X-Sim-Response-Success-Status-Code`: Specifies the success status code. Defaults to `200` if not set.
- `X-Sim-Response-Failure-Status-Code`: Specifies the failure status code. Defaults to `500` if not set.

### Failure Simulation

- `X-Sim-Failure-Percentage`: Defines the percentage of requests that should fail, in the range `[0-100]`. If a random
  number is less than or equal to this value, the failure status code is returned.

### Latency Simulation

- `X-Sim-Latency-ms`: Specifies the latency delay in milliseconds.

## Routes

- Health check: `GET /health` or `GET /healthz`
- Load Simulation: `GET|POST|PUT|DELETE /v1/load/simulate`
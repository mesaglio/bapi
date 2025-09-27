# bapi

A benchmarking project comparing performance across different programming languages and frameworks.

## Documentation

The API specification is available in the [swagger definition file](swagger-3.yaml).

## Testing

API tests can be run using:

```bash
pytest api_test.py
```

These tests validate HTTP status codes and response bodies by making requests to localhost.

## Benchmarking

We use [k6](https://k6.io/) for performance testing. For local dashboard visualization, we recommend using [k6 dashboard](https://github.com/grafana/xk6).

To set up and run the benchmark with dashboard:

```bash
# Install xk6
go install go.k6.io/xk6/cmd/xk6@latest

# Build k6 with dashboard extension
xk6 build --with github.com/grafana/xk6-dashboard@latest

# Run benchmark with dashboard output
./k6 run --out dashboard bench/bench.js
```

## Benchmark Results

<!-- BENCHMARK_RESULTS_START -->
| Framework | Reqs/sec | Latency (p95) | Latency (avg) | Checks OK | Failures |
| :--- | :--- | :--- | :--- | :--- | :--- |
| Rust | 18837.74 | 3.09ms | 1.15ms | 100.00% | 0.00% |
| Go (Native) | 18369.93 | 3.22ms | 1.18ms | 100.00% | 0.00% |
| Go (Mux) | 15992.40 | 3.93ms | 1.36ms | 100.00% | 0.00% |
| Go (Gin) | 15671.24 | 4.11ms | 1.39ms | 100.00% | 0.00% |
| Deno | 15254.52 | 3.85ms | 1.43ms | 100.00% | 0.00% |
| Node (Fastify) | 8009.92 | 8.02ms | 2.76ms | 66.67% | 33.33% |
| Ruby (Sinatra) | 2587.04 | 33.38ms | 8.63ms | 100.00% | 0.00% |
| Python (FastAPI) | 2273.47 | 24.85ms | 9.81ms | 100.00% | 0.00% |
| Python (Flask) | 1189.62 | 11.17ms | 3.03ms | 100.00% | 0.00% |
| Python (Falcon) | 508.75 | 9.12ms | 3.91ms | 100.00% | 0.00% |
<!-- BENCHMARK_RESULTS_END -->

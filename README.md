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
| Rust | 29660.53 | 1.56ms | 0.73ms | 100.00% | 0.00% |
| Go (Native) | 27479.16 | 1.71ms | 0.79ms | 100.00% | 0.00% |
| Go (Mux) | 24939.99 | 2.23ms | 0.87ms | 100.00% | 0.00% |
| Go (Gin) | 24770.47 | 2.28ms | 0.88ms | 100.00% | 0.00% |
| Deno | 21384.11 | 2.55ms | 1.02ms | 100.00% | 0.00% |
| Node (Fastify) | 10978.28 | 5.15ms | 2.02ms | 66.67% | 33.33% |
| Ruby (Sinatra) | 3644.63 | 7.85ms | 6.13ms | 100.00% | 0.00% |
| Python (FastAPI) | 2992.10 | 18.76ms | 7.46ms | 100.00% | 0.00% |
| Python (Flask) | 1598.07 | 5.72ms | 1.61ms | 81.25% | 18.75% |
| Python (Falcon) | 515.25 | 5.98ms | 1.86ms | 99.97% | 0.03% |
<!-- BENCHMARK_RESULTS_END -->

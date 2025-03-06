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

Below are the performance results for each implementation:

- Rust

![alt text](rust-server/image.png "rust")

- Go - Native

![alt text](go-native/image.png "Go - Native")

- Go - Gin

![alt text](go-gin-server/image.png "Go - Gin")

- Go - Mux

![alt text](go-mux-server/image.png "Go - Mux")

- Deno

![alt text](deno-server/image.png "Deno")

- Node - Fastify

![alt text](node-fastify-server/image.png "Node - Fastify")

- Ruby - Sinatra

![alt text](ruby-sinatra-server/image.png "Ruby - Sinatra")

- Python - Flask

![alt text](python-flask-server/image.png "Python - Flask")

- Python - Fastapi

![alt text](python-fastapi-server/image.png "Python - Fastapi")

- Python - Falcon

![alt text](python-falcon-server/image.png "Python - Falcon")

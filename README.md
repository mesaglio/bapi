# bapi

This is a project to test benchmarking for differents lenguage/frameworks.

The swagger definition is [here](https://github.com/mesaglio/bapi/blob/main/swagger-3.yml).

You can run api tests with `pytest api_test.py`, here we have http request to localhost validation http status and response bodys.

---

To run benchamark we use the [k6 script](https://github.com/mesaglio/bapi/blob/main/bench.js). To get a dashboard locally you can use [kibana xk6](https://github.com/grafana/xk6).

```terminal
go install go.k6.io/xk6/cmd/xk6@latest
xk6 build --with github.com/grafana/xk6-dashboard@latest
./k6 run --out dashboard bench/bench.js
```

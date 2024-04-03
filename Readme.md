# Benchmark Web Framework

This Benchmark are benchmark how many request of server this response `GET` request with Hello string

```bash
wrk -t12 -c400 -d10s http://127.0.0.1:{PORT}
```
Running 10s test @ http://127.0.0.1:{PORT}
  12 threads and 400 connections


> Device

* `CPU`: AMD Rysen 7 5800H
* `MEMORY`: 24 GB

> Golang

* Golang version: 1.22

|Frame Work| req/sec | Tranfers/sec |
|--|--|--|
|net/http|531642.67|61.35MB|
|Gin|502816.22|58.02MB|
|Chi|487678.34|56.28MB|

> Javascript

In this report are not test with worker cluster or pm2 yet that make disadvantage concurrency model compare to golang 

* Node: 20.12.0
* bun: 1.0.25


|Frame Work| req/sec | Tranfers/sec |
|--|--|--|
|NodeJS-native|26439.01|4.34MB|
|Node(Express)|8095.41|1.78MB|
|BUN-native|108436.85|12.41MB|
|Bun(Elysia)|78636.33|9.00MB|
.PHONY:
test_http:
	wrk -t12 -c400 -d10s http://127.0.0.1:8080
test_gin:
	wrk -t12 -c400 -d10s http://127.0.0.1:8081
test_chi:
	wrk -t12 -c400 -d10s http://127.0.0.1:8082

server:
	go run ./cmp/main.go

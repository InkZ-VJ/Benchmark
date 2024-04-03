.PHONY:
test: 
	wrk -t12 -c400 -d10s http://127.0.0.1:8080


.PHONY:
node:
	node ./cmp/node.js
express:
	node ./cmp/express.js
bun:
	bun run ./cmp/bun.js
elysia:
	bun run ./cmp/elysia.js

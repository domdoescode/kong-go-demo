.PHONY: all

build:
	docker run --rm -v $(PWD):/plugins kong/go-plugin-tool:2.0.4-alpine-latest build go-demo.go

run:
	docker build -t kong-go-demo . && docker run --rm -p 8000:8000 kong-go-demo

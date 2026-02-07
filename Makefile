.PHONY: build run inspect

build:
	go build -o mcp-server ./cmd/server/main.go

run:
	go run ./cmd/server/main.go

inspect: build
	npx -y @modelcontextprotocol/inspector ./mcp-server

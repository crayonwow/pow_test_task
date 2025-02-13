CLIENT_CONFIG?=config/client.json
SERVER_CONFIG?=config/server.json

lint:
	golangci-lint -v run ./...

build_server:
	go build -o server -v ./cmd/server/main.go
	
build_client:
	go build -o client -v ./cmd/client/main.go

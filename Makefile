CLIENT_CONFIG?=config/client.json
SERVER_CONFIG?=config/server.json

lint:
	golangci-lint -v run ./...

build_server:
	go build -o pow_server -v ./cmd/server/main.go
	
build_client:
	go build -o pow_client -v ./cmd/client/main.go

run_server:
	POW_SERVER_DEBUG=1 POW_SERVER_CONFIG_PATH=$(SERVER_CONFIG) ./pow_server
	
run_client:
	POW_CLIENT_CONFIG_PATH=$(CLIENT_CONFIG) ./pow_client

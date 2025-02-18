CLIENT_CONFIG?=config/client.json
SERVER_CONFIG?=config/server.json
POW_DEBUG?=0

DOCKER_SERVER_BUILD_ARGS=--build-arg POW_DEBUG=$(POW_DEBUG) --build-arg SERVER_CONFIG=$(SERVER_CONFIG)
DOCKER_CLIENT_BUILD_ARGS=--build-arg POW_DEBUG=$(POW_DEBUG) --build-arg SERVER_CONFIG=$(CLIENT_CONFIG)

lint:
	golangci-lint -v run ./...

build_server:
	go build -o pow_server -v ./cmd/server/main.go
	
build_client:
	go build -o pow_client -v ./cmd/client/main.go

run_server:
	POW_CONFIG_PATH=$(SERVER_CONFIG) ./pow_server
	
run_client:
	POW_CONFIG_PATH=$(CLIENT_CONFIG) ./pow_client

run_compose:
	docker-compose up --build

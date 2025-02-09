prepare:
	git submodule add https://github.com/tevador/equix.git
	
lint:
	golangci-lint -v run ./...

build_server:
	go build -o server ./cmd/server/main.go

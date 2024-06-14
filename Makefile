server: test build-server execute-server

test:
	go test ./...

build-server:
	go build -o build/server ./cmd/server

execute-server:
	./build/server

build:
    go build -o bin/gateway cmd/app/main.go

run: build
    ./bin/gateway

test:
    go test ./internal/...

.PHONY: run build test fmt vet lint

run:
	go run ./cmd/server

build:
	go build -o build/lms ./cmd/server

test:
	go test ./...

fmt:
	gofmt -w .

vet:
	go vet ./...

tidy:
	go mod tidy

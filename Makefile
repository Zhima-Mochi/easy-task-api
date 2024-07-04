lint:
	go mod tidy
	golangci-lint run ./... --timeout=10m

test:
	go mod tidy
	go generate ./...
	go test ./... -v -coverprofile .testCoverage.txt

build:
	go mod tidy
	swag init -g main.go
	go build -o build/app/server main.go
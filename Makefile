lint:
	golangci-lint run ./... --timeout=10m

test:
	go mod tidy
	go generate ./...
	go test ./... -v -coverprofile .testCoverage.txt

prepare:
	go mod tidy
	swag init -g main.go

run:
	go run main.go

docker-build:
	docker build -t easy-task-api .

docker-run:
	docker run --rm -it --name easy-task-api -p 8080:8080 easy-task-api

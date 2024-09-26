.PHONY: build
build:
	go build -o bin/service cmd/service/main.go
	
.PHONY: start
start:
	go run cmd/service/main.go

.PHONY: test
test:
	go test -v ./...

.PHONY: test-coverage-html
test-coverage-html:
	go test -coverprofile=coverage.out ./...
	go tool cover -html=coverage.out
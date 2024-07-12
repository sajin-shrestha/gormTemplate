build:
	@go build -o bin/gormTemplate cmd/main.go

test:
	@go test -v ./...

run: build
	@./bin/gormTemplate
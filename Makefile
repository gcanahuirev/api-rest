build:
	@go build -o bin/main cmd/main.go

run: build
	@./bin/main

test:
	@go test -v ./...

hurl:
	@hurl test.hurl --variable host=http://localhost:3000 --report-html ./bin

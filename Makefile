build:
	@go build -o bin/article-go ./cmd/article-go/main.go

run: build
	@./bin/article-go

test:
	@go test -v ./..
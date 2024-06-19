build:
	@go build -o bin/bh

run: build
	@./bin/fs

test:
	@go test ./...

safety: 
	@go vet ./...
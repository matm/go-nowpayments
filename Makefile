.PHONY: np

all: 
	@go build ./cmd/np

mocks:
	@mockery --all --with-expecter --dir .

test: mocks
	@go test ./...

clean:
	@rm -f np

.PHONY: np

all: 
	@go build ./cmd/np

mocks:
	@mockery --all --with-expecter --dir pkg

test: mocks
	@go test ./...

clean:
	@rm -f np

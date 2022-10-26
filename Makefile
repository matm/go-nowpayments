.PHONY: np

all: 
	@go build ./cmd/np

clean:
	@rm -f np

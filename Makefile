
all: fmt test

.PHONY: fmt

fmt:
	@go fmt .


.PHONY: test

test:
	go test .

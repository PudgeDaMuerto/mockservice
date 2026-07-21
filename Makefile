.PHONY: build
build:
	go build -o mockservice ./cmd/cli

.PHONY: run
run:
	go run ./cmd/cli

.PHONY: build_release
build_release:
	go build \
	-ldflags="-s -w" \
	-o "mockservice$(BINARY_EXT)" \
	./cmd/cli

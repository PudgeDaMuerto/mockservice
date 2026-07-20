build:
	go build -o mockservice ./cmd/cli

run:
	go run ./cmd/cli

build_release:
	CGO_ENABLED=0 GOOS=linux go build \
	-ldflags="-s -w" \
	-o mockservice \
	./cmd/cli

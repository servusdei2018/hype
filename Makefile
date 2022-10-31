GO=go
BUILDFLAGS=-ldflags="-s -w"

all:
	$(GO) build $(BUILDFLAGS) ./cmd/hype
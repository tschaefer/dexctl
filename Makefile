.PHONY: all
all: fmt lint build

.PHONY: fmt
fmt:
	gofmt -l .

.PHONY: lint
lint:
	golangci-lint run

.PHONY: build
build:
	goreleaser build --clean --snapshot --single-target

.PHONY: test
test:
	go test -count=1 -v ./pkg/...

.PHONY: test-coverage
test-coverage:
	go test -count=1 -v -coverprofile=coverage.out ./pkg/...

.PHONY: clean
clean:
	rm -rf dist/ coverage/ coverage.out

.PHONY: start-dex
start-dex:
	docker compose --file hack/docker-compose.yml up --detach

.PHONY: stop-dex
stop-dex:
	docker compose --file hack/docker-compose.yml down

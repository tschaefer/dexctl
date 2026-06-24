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

.PHONY: start-dex-container
start-dex-container:
	docker compose --file hack/docker-compose.yml up --detach

.PHONY: stop-dex-container
stop-dex-container:
	docker compose --file hack/docker-compose.yml down

.PHONY: start-dex-daemon
start-dex-daemon:
	daemonize -c /tmp -p /tmp/dex.pid ${PWD}/hack/dex/bin/dex serve ${PWD}/hack/dex/etc/config.yaml

.PHONY: stop-dex-daemon
stop-dex-daemon:
	kill -9 `cat /tmp/dex.pid`

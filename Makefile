DEX_TLS := $(if $(DEX_TLS),-tls,)

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
	docker compose --file hack/docker-compose$(DEX_TLS).yml up --detach

.PHONY: stop-dex-container
stop-dex-container:
	docker compose --file hack/docker-compose$(DEX_TLS).yml down

.PHONY: start-dex-daemon
start-dex-daemon:
	daemonize -c . -p /tmp/dex.pid -o /dev/stdout -e /dev/stderr \
		${PWD}/hack/dex/bin/dex serve hack/dex/etc/dex-daemon$(DEX_TLS).yaml

.PHONY: stop-dex-daemon
stop-dex-daemon:
	kill `cat /tmp/dex.pid`

.PHONY: gen-tls-assets
gen-tls-assets:
	hack/dex/etc/tls/cert-gen

.PHONY: clean-tls-assets
clean-tls-assets:
	hack/dex/etc/tls/cert-clean

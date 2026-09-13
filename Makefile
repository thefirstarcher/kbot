APP      ?= $(shell basename -s .git $(shell git remote get-url origin))
REGISTRY ?= ghcr.io/thefirstarcher
VERSION  ?= $(shell git describe --tags --abbrev=0)-$(shell git rev-parse --short HEAD)

TARGETOS    ?= $(shell go env GOOS)
TARGETARCH  ?= $(shell go env GOARCH)
CGO_ENABLED ?= 0

IMAGE_TAG ?= $(REGISTRY)/$(APP):$(VERSION)-$(TARGETARCH)

.PHONY: help format get build linux arm macos windows image push clean

help:
	@echo "Targets:"
	@echo "  build    - build for \$$(TARGETOS)/\$$(TARGETARCH), defaults to host"
	@echo "  linux    - build for linux/amd64"
	@echo "  arm      - build for linux/arm64"
	@echo "  macos    - build for darwin/arm64"
	@echo "  windows  - build for windows/amd64"
	@echo "  image    - build container image for the host platform"
	@echo "  push     - push image to \$$(REGISTRY)"
	@echo "  clean    - remove binary and image"
	@echo ""
	@echo "Current: TARGETOS=$(TARGETOS) TARGETARCH=$(TARGETARCH)"
	@echo "Image:   $(IMAGE_TAG)"

format:
	gofmt -s -w ./

get:
	go mod tidy
	go mod download

build: format get
	CGO_ENABLED=$(CGO_ENABLED) GOOS=$(TARGETOS) GOARCH=$(TARGETARCH) \
		go build -v -o kbot -ldflags "-X=kbot/cmd.appVersion=$(VERSION) -w -s"

linux:
	$(MAKE) build TARGETOS=linux TARGETARCH=amd64

arm:
	$(MAKE) build TARGETOS=linux TARGETARCH=arm64

macos:
	$(MAKE) build TARGETOS=darwin TARGETARCH=arm64

windows:
	$(MAKE) build TARGETOS=windows TARGETARCH=amd64

image:
	docker build . -t $(IMAGE_TAG) \
		--build-arg TARGETARCH=$(TARGETARCH) \
		--build-arg VERSION=$(VERSION)

push:
	docker push $(IMAGE_TAG)

clean:
	rm -f kbot
	docker rmi $(IMAGE_TAG) || true

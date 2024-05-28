GOARCH ?= amd64
PKG_VERSION != git describe --tags
PKG_VERSION ?= 0.0.0
GCFLAGS ?=
LDFLAGS ?= -w
COMMIT != git rev-parse --short HEAD
APP_NAME ?= words-of-wisdom
PACKAGE_NAME ?= github.com/kormiltsev/proofofwork

build:
	go build -o $(APP_NAME)-$(GOARCH) $(GCFLAGS) -a --ldflags "$(LDFLAGS) \
	-X $(PACKAGE_NAME)/version.Version=$(PKG_VERSION) \
	-X $(PACKAGE_NAME)/version.GitCommit=$(COMMIT)" .
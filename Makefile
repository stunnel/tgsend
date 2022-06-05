GOCMD=go
CGO_ENABLED=0
VERSION=$(shell grep 'VERSION\s=\s' main.go | cut -d'"' -f2)
GOBUILD=$(GOCMD) build
GOCLEAN=$(GOCMD) clean
GOTEST=$(GOCMD) test
BINARY=tgsend
DIR=bin
LDFLAGS+="-s -w"

# Builds the project
build:
		$(GOBUILD) -ldflags=${LDFLAGS} -o $(DIR)/$(BINARY)-$(GOOS)-$(GOARCH)-$(VERSION)
		ln -s -r $(DIR)/$(BINARY)-$(GOOS)-$(GOARCH)-$(VERSION) $(DIR)/$(BINARY)
		tar -Jcvf $(DIR)/$(BINARY)-$(GOOS)-$(GOARCH)-${VERSION}.tar.xz -C $(DIR) $(BINARY)-$(GOOS)-$(GOARCH)-$(VERSION) $(BINARY)
		rm -f $(DIR)/$(BINARY)-$(GOOS)-$(GOARCH)-$(VERSION) $(DIR)/$(BINARY)

build-windows:
		$(GOBUILD) -ldflags=${LDFLAGS} -o $(DIR)/$(BINARY)-$(GOOS)-$(GOARCH)-$(VERSION).exe
		tar -Jcvf $(DIR)/$(BINARY)-$(GOOS)-$(GOARCH)-$(VERSION).tar.xz -C $(DIR) $(BINARY)-$(GOOS)-$(GOARCH)-$(VERSION).exe
		rm -f $(DIR)/$(BINARY)-$(GOOS)-$(GOARCH)-$(VERSION).exe

linux-amd64:
		$(GOCLEAN)
		GOOS=linux GOARCH=amd64 make build

linux-arm64:
		$(GOCLEAN)
		GOOS=linux GOARCH=arm64 make build

linux-arm:
		$(GOCLEAN)
		GOOS=linux GOARCH=arm make build

linux-mips:
		$(GOCLEAN)
		GOOS=linux GOARCH=mips make build

linux-mipsle:
		$(GOCLEAN)
		GOOS=linux GOARCH=mipsle make build

linux-mips64:
		$(GOCLEAN)
		GOOS=linux GOARCH=mips64 make build

linux-mips64le:
		$(GOCLEAN)
		GOOS=linux GOARCH=mips64le make build

freebsd:
		$(GOCLEAN)
		GOOS=freebsd GOARCH=amd64 make build

windows:
		$(GOCLEAN)
		GOOS=windows GOARCH=amd64 make build-windows

mac:
		$(GOCLEAN)
		GOOS=darwin GOARCH=amd64 make build

mac-arm:
		$(GOCLEAN)
		GOOS=darwin GOARCH=arm64 make build

android-termux:
		$(GOCLEAN)
		GOOS=android GOARCH=arm64 make build

releases:
		$(GOCLEAN)

		make linux-amd64
		make linux-arm64
#		make linux-arm
#		make linux-mips
#		make linux-mipsle
#		make linux-mips64
#		make linux-mips64le
		make mac
		make mac-arm
		make freebsd
		make windows

		make clean

clean:
		$(GOCLEAN)

.PHONY:  clean build

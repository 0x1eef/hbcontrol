PREFIX := /usr/local
MANDIR := $(PREFIX)/share/man/man8
BINDIR := $(PREFIX)/bin
BIN := hbcontrol

all: build

fmt:
	go fmt ./...

build:
	go build -o bin/hbcontrol ./cmd/$(BIN)

release:
	go build -buildvcs=false -ldflags="-s -w" -o bin/$(BIN) ./cmd/$(BIN)

install: release
	install -d $(BINDIR)
	install -m 755 bin/$(BIN) $(BINDIR)/$(BIN)
	install -d $(MANDIR)
	install -m 644 share/man/man8/$(BIN).8 $(MANDIR)/$(BIN).8

uninstall:
	rm -f $(BINDIR)/$(BIN)
	rm -f $(MANDIR)/$(BIN).8

.PHONY: all fmt build release install uninstall
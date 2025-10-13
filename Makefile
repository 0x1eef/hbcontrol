PREFIX := /usr/local
MANDIR := $(PREFIX)/share/man/man8
BINDIR := $(PREFIX)/bin

all: build

fmt:
	go fmt ./...

build:
	go build -o bin/control ./cmd/control

release:
	go build -ldflags="-s -w" -o bin/control ./cmd/control

install: release
	install -d $(BINDIR)
	install -m 755 bin/control $(BINDIR)/control
	install -d $(MANDIR)
	install -m 644 share/man/man8/control.8 $(MANDIR)/control.8

uninstall:
	rm -f $(BINDIR)/control
	rm -f $(MANDIR)/control.8

.PHONY: all fmt build release install uninstall
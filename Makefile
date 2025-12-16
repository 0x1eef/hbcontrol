PREFIX := /usr/local
MANDIR := $(PREFIX)/share/man/man8
BINDIR := $(PREFIX)/bin
BIN := ctrl

all: test

fmt:
	for c in control/*.c control/*.h; do \
		clang-format --style="{BasedOnStyle: mozilla, IndentWidth: 4}" -i $$c; \
	done; \
	go fmt ./...

build:
	go build -o bin/ctrl ./cmd/$(BIN)

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

test:
	cd test/ && go test

example:
	go run ./examples/control/

.PHONY: all fmt test
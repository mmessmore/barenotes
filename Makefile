GO_FILES=$(shell find . -name '*.go' | tr '\n' ' ')
BINDIR:=$(HOME)/bin
SHELLCOMPDIR:=$(HOME)/.zsh_functions

messynotes: $(GO_FILES) .pretty
	go build -o $@

.pretty: $(GO_FILES)
	find . -name "*.go" -print0 | xargs -0 goimports -w
	go mod tidy
	touch .pretty

.PHONY: install
install: messynotes
	install -v -m 755 messynotes $(BINDIR)/
	messynotes completion zsh > $(SHELLCOMPDIR)/_messynotes

.PHONY: test
test:
	make -C test

.PHONY: update
update:
	go get -u

.PHONY: clean
clean:
	rm -f messynotes .pretty

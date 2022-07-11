GO_FILES=$(shell find . -name '*.go' | tr '\n' ' ')
BINDIR:=$(HOME)/bin
SHELLCOMPDIR:=$(HOME)/.zsh_functions

barenotes: $(GO_FILES) .pretty
	go build -o $@

.pretty: $(GO_FILES)
	find . -name "*.go" -print0 | xargs -0 goimports -w
	go mod tidy
	touch .pretty

.PHONY: install
install: barenotes
	install -v -m 755 barenotes $(BINDIR)/
	barenotes completion zsh > $(SHELLCOMPDIR)/_barenotes


.PHONY: clean
clean:
	rm -f barenotes .pretty

BIN := go-libra-cli

BUILDENV := CGO_ENABLED=0
BUILDENV += GO111MODULE=on

LINKFLAGS :=-s -extldflags "-static"

$(BIN): clean
	$(BUILDENV) go build -o $(BIN) -a -ldflags '$(LINKFLAGS)' .

.PHONY: clean
clean:
	@rm -f go-libra-cli

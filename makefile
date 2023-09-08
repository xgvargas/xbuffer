BIN := xbuffer
REPO := github.com/xgvargas/${BIN}
VERSION := v0.1.3

FLAGS := -ldflags "-s -w"

MAKEFLAGS += --always-make  # this will make ALL targets PHONY

########
#! podemos usar https://upx.github.io/ para comprimir o executavel se necessario
########

all: test

test:
#	go test -v -count 1 ./...
	go test -count 1 ./...

testdev:
	nodemon -e 'go' --exec 'go test -count 1 ./...' --signal SIGTERM

deps:
	go mod tidy

doc:
	@echo "Serving documentation on http://localhost:6060"
	godoc -http=:6060

publish: deps
#	git commit -m"???..."
#	git tag -a -m "some description..." ${VERSION}
#	git push && git push --tags
#	GOPROXY=proxy.golang.org go list -m ${REPO}@${VERSION}

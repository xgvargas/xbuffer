BIN := xbuffer
FLAGS := -ldflags "-s -w"
REPO := github.com/xgvargas/${BIN}
VERSION := v0.1.2

########
#! podemos usar https://upx.github.io/ para comprimir o executavel se necessario
########

all: build

.PHONY: run
run:
	go run .

.PHONY: build
build: deps
#	GOOS=linux GOARCH=386 go build ${FLAGS} -o ${BIN} ./main.go
#	GOOS=windows GOARCH=386 go build ${FLAGS} -o ${BIN}.exe ./main.go
	GOOS=linux go build ${FLAGS} -o ${BIN} ./main.go

.PHONY: clean
clean:
	rm -f ${BIN}
	rm -f ${BIN}.exe

.PHONY: test
test:
	go test -v -count 1 ./...

.PHONY: deps
deps:
	go generate
	go mod tidy
	go mod download

.PHONY: doc
doc:
	@echo "Serving documentation on http://localhost:6060"
	godoc -http=:6060

.PHONY: publish
publish:
#	git commit -m"???"
#	git tag ${VERSION}
#	git push --tags
#	GOPROXY=proxy.golang.org go list -m ${REPO}@${VERSION}

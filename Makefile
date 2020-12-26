# Miscellaneous make tasks.
sources := .

all: clean format lint protobuf test

clean:
	go clean

dependencies:
	go get -u github.com/onsi/ginkgo/ginkgo
	go get -u golang.org/x/lint/golint
	go get golang.org/x/tools/cmd/goimports
	go install google.golang.org/protobuf/cmd/protoc-gen-go

format:
	goimports -w $(sources)

lint:
	golint $(sources)
	golangci-lint run

protobuf:
	protoc -I=./ --go_out=./ ./event.proto

test: format
	ginkgo -r -race -randomizeSuites ./

upgrade:
	go get -u ./...

.PHONY: install-generate generate build run test clean

build: generate fmt vet
	go build -o bin/income-atlas .

build-dev: generate fmt vet
	go build -o tmp/main .

install-generate:
	go install github.com/atombender/go-jsonschema@latest

generate:
	go-jsonschema \
    -p main \
    schema/*.schema.json > generated.go
	gofmt -w generated.go

fmt:
	gofmt -w .

vet:
	go vet ./...

run: generate
	go run .

test: generate
	go test ./...

clean:
	rm -rf bin
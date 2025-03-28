.DEFAULT_GOAL := build
binary_name = tipcalc

.PHONEY: build fmt test vet clean 

fmt:
	go fmt ./...

vet: fmt
	go vet ./...
	staticcheck ./...

test: vet
	go test ./...

build: test
	go build

clean:
	go clean
	rm -f ${binary_name}

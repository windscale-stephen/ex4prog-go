.DEFAULT_GOAL := build
binary_name = sayhello

.PHONEY: build fmt test vet clean 

fmt:
	goimports -l -w .

vet: fmt
	go vet ./...
	staticcheck ./...
	govulncheck ./...

test: vet
	go test ./...

build: test
	go build -o ${binary_name}

clean:
	go clean
	rm -f ${binary_name}

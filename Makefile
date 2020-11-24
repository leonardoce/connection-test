bin/connection_test: fmt vet
	go generate cmd/connection_test/main.go
	go build -o bin/connection_test cmd/connection_test/main.go

.PHONY:fmt
fmt:
	go fmt ./...

.PHONY: vet
vet:
	go vet ./...

.PHONY: all
all: bin/connection_test

.PHONY: run
run: bin/connection_test
	bin/connection_test
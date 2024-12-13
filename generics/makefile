.DEFAULT_GOAL := run

.PHONY: fmt vet build run clean

fmt:
	go fmt ./main.go
vet:
	go vet ./main.go
clean:
	rm -f generics
build: clean
	go build -o generics ./main.go
run: build
	./generics
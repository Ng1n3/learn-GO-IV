.DEFAULT_GOAL := run

.PHONY: fmt vet build run clean

fmt:
	go fmt ./main.go
vet:
	go vet ./main.go
clean:
	rm -f pointers
build: clean
	go build -o pointers ./main.go
run: build
	./pointers
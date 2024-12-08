.DEFAULT_GOAL := run

.PHONY: fmt vet build run clean

fmt:
	go fmt ./main.go
vet:
	go vet ./main.go
clean:
	rm -f blocks
build: clean
	go build -o blocks ./main.go
run: build
	./blocks
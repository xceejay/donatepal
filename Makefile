.PHONY : build run

build:
	@echo "building program..."
	go mod tidy
	go build -o bin/donatepal main.go
	ln -sf bin/donatepal donatepal

run: build
	./donatepal

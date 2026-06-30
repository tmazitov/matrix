.PHONY: all test build clean

EX=ex01

all: build

build:
	go build -o $(EX) ./cmd 

test:
	go test ./...

test-v:
	go test -v ./...

clean:
	go clean ./...

fclean: clean
	rm -rf ./ex00

re: fclean all
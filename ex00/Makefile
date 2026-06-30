.PHONY: all test build clean

all: build

build:
	go build -o ex00

test:
	go test ./...

test-v:
	go test -v ./...

clean:
	go clean ./...

fclean: clean
	rm -rf ./ex00

re: fclean all

build:
	@go build  -o bin/five .

test2: build
	@./bin/five test1.txt test2.txt

part2: build
	@./bin/five part1.txt part2.txt

test1: build
	@./bin/five test1.txt

part1: build
	@./bin/five part1.txt

test:
	@go test -v ./...


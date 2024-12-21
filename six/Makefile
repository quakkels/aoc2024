
build:
	@go build  -o bin/six .

test2: build
	@./bin/six test1.txt test2.txt

part2: build
	@./bin/six part1.txt part2.txt

test1: build
	@./bin/six test1.txt

part1: build
	@./bin/six part1.txt

test:
	@go test -v ./...


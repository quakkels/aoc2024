
build:
	@go build  -o bin/four .

runtest: build
	@./bin/four test.txt

run: build
	@./bin/four input.txt

test:
	@go test -v ./...


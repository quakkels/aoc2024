
build:
	@go build  -o bin/three .

run-test: build
	@./bin/three test.txt

run: build
	@./bin/three input.txt

test:
	@go test -v ./...


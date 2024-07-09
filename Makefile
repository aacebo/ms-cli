clean:
	rm -rf ./bin

build:
	go build -o bin/main ./cmd

win_build:
	go build -o bin/main.exe ./cmd

clean_build: clean build

fmt:
	gofmt -w ./

test:
	go clean -testcache
	go test ./... -cover

run:
	go run ./cmd

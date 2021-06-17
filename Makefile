.PHONY:
test:
	go test -v ./...

build:
	go build -ldflags '-s -w' -o ./bin/main ./cmd

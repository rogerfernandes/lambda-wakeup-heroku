OS=$(shell uname -s)

build:
	GOOS=linux GOARCH=amd64 go build -o main main.go
	zip main.zip main
	rm main
.PHONY: build

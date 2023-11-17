.PHONY:
.SILENT:

clear:
	rm -rf ./bin/app
	go clean

build:
	go build -o ./bin/app ./cmd

run:
	./bin/app

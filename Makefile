

build-all: build

build:
	go build -o ./bin/server ./server

run-server: 
	./bin/server

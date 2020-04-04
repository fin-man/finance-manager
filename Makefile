

build-all: build build-filewatcher

build:
	go build -o ./bin/server ./server

build-filewatcher:
	go build -o ./bin/filewatcher ./filewatcher

run-server: 
	./bin/server

run-filewatcher:
	./bin/filewatcher
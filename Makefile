

build-all: build run-server

build:
	go build -o ./bin/server ./server

run-server: 
	./bin/server

db-exec:
	docker exec -it postgres_finman psql -U postgres -d finances
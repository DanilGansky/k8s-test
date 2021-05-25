deps:
	go mod download

build:
	go build -o ./bin/server

clean:
	rm -rf ./bin
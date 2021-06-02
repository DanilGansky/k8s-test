deps:
	go mod download

build:
	go build -o ./bin/server

build-release:
	CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o ./bin/server .

clean:
	rm -rf ./bin
build:
	go build -o bin/play-dog

run: build
	 ./bin/play-dog

test:
	go test ./...
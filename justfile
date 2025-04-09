build:
	go mod download
	go build -o lura

install:
	go install

run:
	go run .

build:
	go mod download
	go build -o lura ./cmd/lura/main.go

install:
	go install

run:
	go run ./cmd/lura/main.go

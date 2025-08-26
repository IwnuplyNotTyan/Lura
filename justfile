build:
	go mod download
	go build -o ./bin/lura ./cmd/lura/main.go

install:
	go install

run:
	go run ./cmd/lura/main.go

debug:
	go run ./cmd/lura/main.go -debug

verbose:
	go run ./cmd/lura/main.go -v

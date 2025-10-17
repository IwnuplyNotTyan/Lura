build:
	go mod download
	GOARCH=amd64 go build -o ./bin/lura ./cmd/lura/main.go
	GOOS=windows GOARCH=amd64 go build -o ./bin/lura.exe ./cmd/lura/main.go


install:
	go install

run:
	go run ./cmd/lura/main.go

debug:
	go run ./cmd/lura/main.go -debug

verbose:
	go run ./cmd/lura/main.go -v

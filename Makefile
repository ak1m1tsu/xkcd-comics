build:
	go build -o ./bin/xkcd ./cmd/xkcd/main.go

fetch:
	go run ./cmd/xkcd-data/main.go

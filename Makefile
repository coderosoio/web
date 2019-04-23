build:
	go build -a -i -o cmd/web/web cmd/web/main.go

run:
	go run cmd/web/*.go

.PHONY: build

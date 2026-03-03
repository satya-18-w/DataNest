build:
	go build -o datanest cmd/main.go

run:
	go run cmd/main.go

test:
	go test ./... -v

tidy:
	go mod tidy
clean:
	rm -f datanest
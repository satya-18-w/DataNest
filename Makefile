GO := $(shell which go 2>/dev/null || which go.exe 2>/dev/null || echo go)

build:
	"$(GO)" build -o datanest cmd/main.go

run:
	"$(GO)" run cmd/main.go

test:
	"$(GO)" test ./... -v

tidy:
	"$(GO)" mod tidy
clean:
	rm -f datanest
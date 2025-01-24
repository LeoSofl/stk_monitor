
.PHONY: build run test clean

build:
	go build -o bin/stk-monitor cmd/stk-monitor/main.go

run:
	go run cmd/stk-monitor/main.go

test:
	go test ./...

clean:
	rm -rf bin/
.PHONY: clean test build run

default: build

clean:
	@echo "*** Cleaning repo of unnecessary artifacts... ***"
	rm -rf bin

test: clean
	@echo "*** Testing the project... ***"
	go test internal/networking

build: test
	@echo "*** Building the project.. ***"
	go mod tidy
	go build -ldflags="-s -w" -o bin/bitcoin_node_metrics cmd/httpserver/main.go

run:
	@echo "*** Running the program.. ***"
	./bin/bitcoin_node_metrics

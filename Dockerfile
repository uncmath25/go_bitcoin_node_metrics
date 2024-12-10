FROM golang:1.23.1-bookworm

WORKDIR /app

COPY go.mod go.sum ./

RUN go mod download

COPY . .

RUN go mod tidy
RUN go build -ldflags="-s -w" -o bin/bitcoin_node_metrics cmd/httpserver/main.go

EXPOSE 8080

CMD ["./bin/bitcoin_node_metrics"]

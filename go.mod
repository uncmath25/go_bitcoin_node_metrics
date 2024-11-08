module go_bitcoin_node_metrics

go 1.23.1

replace internal/networking => ./internal/networking

require internal/networking v0.0.0-00010101000000-000000000000

require github.com/gorilla/mux v1.8.1 // indirect

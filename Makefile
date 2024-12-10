.PHONY: clean test build run stopdocker startdocker deploy

IMAGE_NAME="uncmath25/go_bitcoin_node_metrics"
CONTAINER_NAME="go_bitcoin_node_metrics"
REMOTE_SERVER_PROFILE="testinglab"
REMOTE_PARENT_WEBSITE_DIR="/home/player1/websites/go_bitcoin_node_metrics"

default: build

clean:
	@echo "*** Cleaning repo of unnecessary artifacts... ***"
	rm -rf bin

test: clean
	@echo "*** Testing the project... ***"
	go test ./internal/networking

build: test
	@echo "*** Building the project.. ***"
	go mod tidy
	go build -ldflags="-s -w" -o bin/bitcoin_node_metrics cmd/httpserver/main.go

run:
	@echo "*** Running the program.. ***"
	./bin/bitcoin_node_metrics

stopdocker: clean
	@echo "*** Stopping local Dockerized go server... ***"
	docker rm -f $(CONTAINER_NAME)

startdocker: stopdocker
	@echo "*** Starting local Dockerized go server... ***"
	docker build -t $(IMAGE_NAME) -f Dockerfile .
	docker run --rm -d --env-file=.env -p 8080:8080 --name $(CONTAINER_NAME) $(IMAGE_NAME)

deploy: clean
	@echo "*** Deploying Dockerized go server to DigitalOcean droplet... ***"
	docker build --platform=linux/x86_64 -t $(IMAGE_NAME) .
	docker save $(IMAGE_NAME) | ssh -C $(REMOTE_SERVER_PROFILE) docker load
	ssh $(REMOTE_SERVER_PROFILE) rm -rf $(REMOTE_PARENT_WEBSITE_DIR)
	scp -r ./* $(REMOTE_SERVER_PROFILE):$(REMOTE_PARENT_WEBSITE_DIR)
	scp -r .env $(REMOTE_SERVER_PROFILE):$(REMOTE_PARENT_WEBSITE_DIR)
	ssh $(REMOTE_SERVER_PROFILE) "\
		cd $(REMOTE_PARENT_WEBSITE_DIR); \
		docker rm -f $(CONTAINER_NAME); \
		docker run --rm -d --env-file=.env --network host --name $(CONTAINER_NAME) $(IMAGE_NAME); \
	"
	@echo "*** Restart the remote server with _restart_server.sh ***"

package client

import (
	"bytes"
	"encoding/json"
	"errors"
	"log"
	"net/http"
	"os"
)

type Client interface {
	GetBlockHeight() (int, error)
}

type client struct {
	logger *log.Logger
}

func BuildClient(logger *log.Logger) Client {
	return &client{
		logger: logger,
	}
}

func (c *client) GetBlockHeight() (int, error) {
	client := &http.Client{}
	url := os.Getenv("BITCOIN_RPC_URL")
	body := "{\"jsonrpc\": \"1.0\", \"id\": \"curltest\", \"method\": \"getblockcount\", \"params\": []}"
	req, err := http.NewRequest("POST", url, bytes.NewBufferString(body))
	if err != nil {
		return -1, err
	}
	username := os.Getenv("BITCOIN_RPC_USERNAME")
	password := os.Getenv("BITCOIN_RPC_PASSWORD")
	req.SetBasicAuth(username, password)
	resp, err := client.Do(req)
	if err != nil {
		return -1, err
	}
	var data map[string]interface{}
	json.NewDecoder(resp.Body).Decode(&data)
	var blockHeightRaw float64
	var ok bool
	if x, found := data["result"]; found {
		if blockHeightRaw, ok = x.(float64); !ok {
			return -1, errors.New("result value in GetBlockHeight not castable as int")
		}
	} else {
		return -1, errors.New("result key not found in GetBlockHeight response")
	}
	blockHeight := int(blockHeightRaw)
	c.logger.Println("Found blockHeight:", blockHeight)
	return blockHeight, nil
}

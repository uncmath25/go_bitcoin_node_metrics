package client

import "log"

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
	return 1, nil
}

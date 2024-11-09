package service

import (
	"go_bitcoin_node_metrics/internal/client"
	"log"
)

type Service interface {
	GetTestMessage() (string, error)
	BuildNodeSummary() (int, error)
}

type service struct {
	client client.Client
	logger *log.Logger
}

func BuildService(client client.Client, logger *log.Logger) Service {
	return &service{
		client: client,
		logger: logger,
	}
}

func (s *service) GetTestMessage() (string, error) {
	// return "Test", errors.New("fake error")
	return "Test", nil
}

func (s *service) BuildNodeSummary() (int, error) {
	blockHeight, err := s.client.GetBlockHeight()
	if err != nil {
		s.logger.Fatalln("Failed to build node summary:", err)
	}
	return blockHeight, nil
}

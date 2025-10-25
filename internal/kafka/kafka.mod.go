package kafka

import (
	dto "github.com/easy-attend-serviceV3/internal/kafka/dto"
)

type Module struct {
	Svc *Service
}

func New(conf *dto.Kafka) *Module {
	return &Module{
		Svc: newService(conf),
	}
}

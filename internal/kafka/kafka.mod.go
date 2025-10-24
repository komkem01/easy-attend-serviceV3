package kafka

import (
	dto "mcop/internal/kafka/dto"
)

type Module struct {
	Svc *Service
}

func New(conf *dto.Kafka) *Module {
	return &Module{
		Svc: newService(conf),
	}
}

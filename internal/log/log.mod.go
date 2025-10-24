package log

import (
	configdto "mcop/internal/config/dto"
)

type Module struct {
	Svc *Service
	Mid *Middleware
}

func New(conf *configdto.Config[Option]) *Module {
	svc := newService(conf)
	mid := NewMiddleware(conf, svc)
	return &Module{
		Svc: svc,
		Mid: mid,
	}
}

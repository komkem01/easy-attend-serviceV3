package log

import (
	configdto "mcop/internal/config/dto"
)

type Middleware struct {
	Config *configdto.Config[Option]
	Svc    *Service
}

func NewMiddleware(conf *configdto.Config[Option], svc *Service) *Middleware {
	return &Middleware{
		Config: conf,
		Svc:    svc,
	}
}

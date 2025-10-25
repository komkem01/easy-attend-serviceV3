package exampletwo

import (
	entitiesinf "github.com/easy-attend-serviceV3/app/modules/entities/inf"
	configDTO "github.com/easy-attend-serviceV3/internal/config/dto"
)

type (
	Module struct {
		Svc *Service
		Ctl *Controller
	}
	Service    struct{}
	Controller struct{}

	Config struct{}
)

func New(conf *configDTO.Config[Config], db entitiesinf.ExampleEntity) *Module {
	return &Module{
		Svc: &Service{},
		Ctl: &Controller{},
	}
}

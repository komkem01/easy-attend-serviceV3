package collector

import configdto "github.com/easy-attend-serviceV3/internal/config/dto"

type Module struct {
	Svc *Service
}

func New(conf *configdto.Config[Config]) *Module {
	return &Module{
		Svc: newService(conf),
	}
}

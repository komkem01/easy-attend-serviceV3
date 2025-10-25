package entities

import (
	"context"

	"github.com/easy-attend-serviceV3/app/modules/entities/ent"
	entitiesinf "github.com/easy-attend-serviceV3/app/modules/entities/inf"
	"github.com/google/uuid"
)

var _ entitiesinf.GenderEntity = (*Service)(nil)

func (s *Service) GetListGender(ctx context.Context) ([]*ent.GenderEntity, error) {
	var genders []*ent.GenderEntity
	err := s.db.NewSelect().Model(&genders).Scan(ctx)
	if err != nil {
		return nil, err
	}
	return genders, nil
}

func (s *Service) GetByIDGender(ctx context.Context, id uuid.UUID) (*ent.GenderEntity, error) {
	var gender ent.GenderEntity
	err := s.db.NewSelect().Model(&gender).Where("id = ?", id).Scan(ctx)
	if err != nil {
		return nil, err
	}
	return &gender, nil
}

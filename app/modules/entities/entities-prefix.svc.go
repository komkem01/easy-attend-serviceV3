package entities

import (
	"context"

	"github.com/easy-attend-serviceV3/app/modules/entities/ent"
	entitiesinf "github.com/easy-attend-serviceV3/app/modules/entities/inf"
	"github.com/google/uuid"
)

var _ entitiesinf.PrefixEntity = (*Service)(nil)

func (s *Service) GetListPrefix(ctx context.Context) ([]*ent.PrefixEntity, error) {
	var prefixes []*ent.PrefixEntity
	err := s.db.NewSelect().Model(&prefixes).Scan(ctx)
	if err != nil {
		return nil, err
	}
	return prefixes, nil
}

func (s *Service) GetByIDPrefix(ctx context.Context, id uuid.UUID) (*ent.PrefixEntity, error) {
	var prefix ent.PrefixEntity
	err := s.db.NewSelect().Model(&prefix).Where("id = ?", id).Scan(ctx)
	if err != nil {
		return nil, err
	}
	return &prefix, nil
}

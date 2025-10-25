package entities

import (
	"context"
	"fmt"
	"time"

	"github.com/easy-attend-serviceV3/app/modules/entities/ent"
	entitiesinf "github.com/easy-attend-serviceV3/app/modules/entities/inf"
	"github.com/google/uuid"
)

var _ entitiesinf.SchoolEntity = (*Service)(nil)

func (s *Service) GetListSchool(ctx context.Context) ([]*ent.SchoolEntity, error) {
	var schools []*ent.SchoolEntity
	err := s.db.NewSelect().Model(&schools).Scan(ctx)
	if err != nil {
		return nil, err
	}
	return schools, nil
}

func (s *Service) GetByIDSchool(ctx context.Context, id uuid.UUID) (*ent.SchoolEntity, error) {
	var school ent.SchoolEntity
	err := s.db.NewSelect().Model(&school).Where("id = ?", id).Scan(ctx)
	if err != nil {
		return nil, err
	}
	return &school, nil
}

func (s *Service) CreateSchool(ctx context.Context, name, address, phone string) (*ent.SchoolEntity, error) {
	school := &ent.SchoolEntity{
		ID:      uuid.New(),
		Name:    name,
		Address: address,
		Phone:   phone,
	}
	// Set creation and update timestamps
	school.CreatedAt = time.Now()
	school.UpdatedAt = time.Now()
	_, err := s.db.NewInsert().Model(school).Exec(ctx)
	if err != nil {
		return nil, err
	}
	return school, nil
}

func (s *Service) UpdateSchool(ctx context.Context, id uuid.UUID, name, address, phone string) (*ent.SchoolEntity, error) {
	school, err := s.GetByIDSchool(ctx, id)
	if err != nil {
		return nil, err
	}
	school.Name = name
	school.Address = address
	school.Phone = phone
	_, err = s.db.NewUpdate().Model(school).Where("id = ?", id).Exec(ctx)
	if err != nil {
		return nil, err
	}
	return school, nil
}

func (s *Service) DeleteSchool(ctx context.Context, id uuid.UUID) error {
	_, err := s.db.NewDelete().Model(&ent.SchoolEntity{}).Where("id = ?", id).Exec(ctx)
	return err
}

func (s *Service) CheckExistSchool(ctx context.Context, id uuid.UUID) (bool, error) {
	count, err := s.db.NewSelect().Model(&ent.SchoolEntity{}).Where("id = ?", id).Count(ctx)
	if err != nil {
		return false, err
	}
	if count == 0 {
		return false, fmt.Errorf("school with id %s does not exist", id)
	}
	return true, nil
}

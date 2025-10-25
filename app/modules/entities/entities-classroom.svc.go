package entities

import (
	"context"
	"fmt"
	"time"

	"github.com/easy-attend-serviceV3/app/modules/entities/ent"
	entitiesinf "github.com/easy-attend-serviceV3/app/modules/entities/inf"
	"github.com/google/uuid"
)

var _ entitiesinf.ClassroomEntity = (*Service)(nil)

func (s *Service) GetListClassroom(ctx context.Context) ([]*ent.ClassroomEntity, error) {
	var classrooms []*ent.ClassroomEntity
	err := s.db.NewSelect().Model(&classrooms).Scan(ctx)
	if err != nil {
		return nil, err
	}
	return classrooms, nil
}

func (s *Service) GetByIDClassroom(ctx context.Context, id uuid.UUID) (*ent.ClassroomEntity, error) {
	var classroom ent.ClassroomEntity
	err := s.db.NewSelect().Model(&classroom).Where("id = ?", id).Scan(ctx)
	if err != nil {
		return nil, err
	}
	return &classroom, nil
}

func (s *Service) CreateClassroom(ctx context.Context, schoolID uuid.UUID, name string) (*ent.ClassroomEntity, error) {
	classroom := &ent.ClassroomEntity{
		ID:       uuid.New(),
		SchoolID: schoolID,
		Name:     name,
	}
	classroom.CreatedAt = time.Now()
	classroom.UpdatedAt = time.Now()
	_, err := s.db.NewInsert().Model(classroom).Exec(ctx)
	if err != nil {
		return nil, err
	}
	return classroom, nil
}

func (s *Service) UpdateClassroom(ctx context.Context, id uuid.UUID, schoolID uuid.UUID, name string) (*ent.ClassroomEntity, error) {
	classroom, err := s.GetByIDClassroom(ctx, id)
	if err != nil {
		return nil, err
	}
	classroom.SchoolID = schoolID
	classroom.Name = name
	classroom.UpdatedAt = time.Now()
	_, err = s.db.NewUpdate().Model(classroom).Where("id = ?", id).Exec(ctx)
	if err != nil {
		return nil, err
	}
	return classroom, nil
}

func (s *Service) DeleteClassroom(ctx context.Context, id uuid.UUID) error {
	_, err := s.db.NewDelete().Model(&ent.ClassroomEntity{}).Where("id = ?", id).Exec(ctx)
	return err
}

func (s *Service) CheckExistClassroom(ctx context.Context, id uuid.UUID) (bool, error) {
	count, err := s.db.NewSelect().Model(&ent.ClassroomEntity{}).Where("id = ?", id).Count(ctx)
	if err != nil {
		return false, err
	}
	if count == 0 {
		return false, fmt.Errorf("classroom with id %s does not exist", id)
	}
	return true, nil
}

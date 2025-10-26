package entities

import (
	"context"
	"time"

	entitiesdto "github.com/easy-attend-serviceV3/app/modules/entities/dto"
	"github.com/easy-attend-serviceV3/app/modules/entities/ent"
	entitiesinf "github.com/easy-attend-serviceV3/app/modules/entities/inf"
	"github.com/google/uuid"
)

var _ entitiesinf.TeacherEntity = (*Service)(nil)

func (s *Service) CreateTeacher(ctx context.Context, req *entitiesdto.TeacherCreateRequest) (*ent.TeacherEntity, error) {
	// Generate new UUID for teacher
	teacherID := uuid.New()

	teacher := &ent.TeacherEntity{
		ID:          teacherID,
		SchoolID:    req.SchoolID,
		ClassroomID: req.ClassroomID,
		PrefixID:    req.PrefixID,
		GenderID:    req.GenderID,
		FirstName:   req.FirstName,
		LastName:    req.LastName,
		Email:       req.Email,
		Password:    req.Password,
		Phone:       req.Phone,
	}
	teacher.CreatedAt = time.Now()
	teacher.UpdatedAt = time.Now()

	_, err := s.db.NewInsert().Model(teacher).Exec(ctx)
	if err != nil {
		return nil, err
	}
	return teacher, nil
}

func (s *Service) GetListTeacher(ctx context.Context) ([]*ent.TeacherEntity, error) {
	var teachers []*ent.TeacherEntity
	err := s.db.NewSelect().Model(&teachers).Scan(ctx)
	if err != nil {
		return nil, err
	}
	return teachers, nil
}

func (s *Service) GetByIDTeacher(ctx context.Context, id uuid.UUID) (*ent.TeacherEntity, error) {
	var teacher ent.TeacherEntity
	err := s.db.NewSelect().Model(&teacher).Where("id = ?", id).Scan(ctx)
	if err != nil {
		return nil, err
	}
	return &teacher, nil
}

func (s *Service) UpdateTeacher(ctx context.Context, id uuid.UUID, req *entitiesdto.TeacherUpdateRequest) (*ent.TeacherEntity, error) {
	teacher := &ent.TeacherEntity{
		ID:          id,
		SchoolID:    req.SchoolID,
		ClassroomID: req.ClassroomID,
		PrefixID:    req.PrefixID,
		GenderID:    req.GenderID,
		FirstName:   req.FirstName,
		LastName:    req.LastName,
		Email:       req.Email,
		Phone:       req.Phone,
	}
	teacher.UpdatedAt = time.Now()

	_, err := s.db.NewUpdate().Model(teacher).Where("id = ?", id).Exec(ctx)
	if err != nil {
		return nil, err
	}
	return teacher, nil
}

func (s *Service) DeleteTeacher(ctx context.Context, id uuid.UUID) error {
	_, err := s.db.NewDelete().Model(&ent.TeacherEntity{}).Where("id = ?", id).Exec(ctx)
	return err
}

func (s *Service) CheckExistTeacher(ctx context.Context, id uuid.UUID) (bool, error) {
	count, err := s.db.NewSelect().Model(&ent.TeacherEntity{}).Where("id = ?", id).Count(ctx)
	if err != nil {
		return false, err
	}
	return count > 0, nil
}

// GetTeacherByEmail retrieves a teacher by email
func (s *Service) GetTeacherByEmail(ctx context.Context, email string) (*ent.TeacherEntity, error) {
	var teacher ent.TeacherEntity

	err := s.db.NewSelect().
		Model(&teacher).
		Where("email = ? AND deleted_at IS NULL", email).
		Scan(ctx)

	if err != nil {
		return nil, err
	}

	return &teacher, nil
}

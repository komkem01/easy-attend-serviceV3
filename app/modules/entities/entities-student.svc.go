package entities

import (
	"context"
	"time"

	entitiesdto "github.com/easy-attend-serviceV3/app/modules/entities/dto"
	"github.com/easy-attend-serviceV3/app/modules/entities/ent"
	entitiesinf "github.com/easy-attend-serviceV3/app/modules/entities/inf"
	"github.com/google/uuid"
)

var _ entitiesinf.StudentEntity = (*Service)(nil)

func (s *Service) CreateStudent(ctx context.Context, req *entitiesdto.StudentCreateRequest) (*ent.StudentEntity, error) {
	student := &ent.StudentEntity{
		ID:          uuid.New(),
		SchoolID:    req.School,
		ClassroomID: req.Classroom,
		PrefixID:    req.Prefix,
		GenderID:    req.Gender,
		StudentCode: req.StudentCode,
		FirstName:   req.FirstName,
		LastName:    req.LastName,
		Phone:       req.Phone,
	}
	student.CreatedAt = time.Now()
	student.UpdatedAt = time.Now()
	_, err := s.db.NewInsert().Model(student).Exec(ctx)
	if err != nil {
		return nil, err
	}
	return student, nil
}

func (s *Service) UpdateStudent(ctx context.Context, id uuid.UUID, req *entitiesdto.StudentUpdateRequest) (*ent.StudentEntity, error) {
	student := &ent.StudentEntity{}
	err := s.db.NewSelect().Model(student).Where("id = ?", id).Scan(ctx)
	if err != nil {
		return nil, err
	}
	student.SchoolID = req.School
	student.ClassroomID = req.Classroom
	student.PrefixID = req.Prefix
	student.GenderID = req.Gender
	student.StudentCode = req.StudentCode
	student.FirstName = req.FirstName
	student.LastName = req.LastName
	student.Phone = req.Phone
	student.UpdatedAt = time.Now()
	_, err = s.db.NewUpdate().Model(student).Where("id = ?", id).Exec(ctx)
	if err != nil {
		return nil, err
	}
	return student, nil
}

func (s *Service) GetListStudent(ctx context.Context, resp *entitiesdto.StudentListResponse) ([]*ent.StudentEntity, error) {
	var students []*ent.StudentEntity
	err := s.db.NewSelect().Model(&students).Scan(ctx)
	if err != nil {
		return nil, err
	}
	return students, nil
}

func (s *Service) GetStudentByID(ctx context.Context, id uuid.UUID, resp *entitiesdto.StudentInfoResponse) (*ent.StudentEntity, error) {
	student := &ent.StudentEntity{}
	err := s.db.NewSelect().Model(student).Where("id = ?", id).Scan(ctx)
	if err != nil {
		return nil, err
	}
	return student, nil
}

func (s *Service) DeleteStudent(ctx context.Context, id uuid.UUID) error {
	student := &ent.StudentEntity{}
	_, err := s.db.NewDelete().Model(student).Where("id = ?", id).Exec(ctx)
	if err != nil {
		return err
	}
	return nil
}

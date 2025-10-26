package entities

import (
	"context"
	"time"

	entitiesdto "github.com/easy-attend-serviceV3/app/modules/entities/dto"
	"github.com/easy-attend-serviceV3/app/modules/entities/ent"
	"github.com/google/uuid"
)

// CreateClassroomMember creates a new classroom member
func (s *Service) CreateClassroomMember(ctx context.Context, req *entitiesdto.ClassroomMemberCreateRequest) (*ent.ClassroomMemberEntity, error) {
	// Generate new UUID for classroom member
	memberID := uuid.New()

	member := &ent.ClassroomMemberEntity{
		ID:          memberID,
		ClassroomID: req.ClassroomID,
		StudentID:   req.StudentID,
		TeacherID:   req.TeacherID,
	}
	member.CreatedAt = time.Now()
	member.UpdatedAt = time.Now()

	_, err := s.db.NewInsert().Model(member).Exec(ctx)
	if err != nil {
		return nil, err
	}
	return member, nil
}

// GetListClassroomMember retrieves all members of a specific classroom
func (s *Service) GetListClassroomMember(ctx context.Context, classroomID uuid.UUID) ([]*ent.ClassroomMemberEntity, error) {
	var members []*ent.ClassroomMemberEntity
	err := s.db.NewSelect().
		Model(&members).
		Where("classroom_id = ?", classroomID).
		Scan(ctx)
	if err != nil {
		return nil, err
	}
	return members, nil
}

// GetClassroomMembersByTeacherID retrieves all classroom members for a specific teacher
func (s *Service) GetClassroomMembersByTeacherID(ctx context.Context, teacherID uuid.UUID) ([]*ent.ClassroomMemberEntity, error) {
	var members []*ent.ClassroomMemberEntity
	err := s.db.NewSelect().
		Model(&members).
		Where("teacher_id = ?", teacherID).
		Scan(ctx)
	if err != nil {
		return nil, err
	}
	return members, nil
}

// GetClassroomMemberByID retrieves a classroom member by ID
func (s *Service) GetClassroomMemberByID(ctx context.Context, id uuid.UUID) (*ent.ClassroomMemberEntity, error) {
	var member ent.ClassroomMemberEntity

	err := s.db.NewSelect().
		Model(&member).
		Where("id = ?", id).
		Scan(ctx)

	if err != nil {
		return nil, err
	}

	return &member, nil
}

// UpdateClassroomMember updates a classroom member
func (s *Service) UpdateClassroomMember(ctx context.Context, id uuid.UUID, req *entitiesdto.ClassroomMemberUpdateRequest) (*ent.ClassroomMemberEntity, error) {
	member := &ent.ClassroomMemberEntity{
		ID:          id,
		ClassroomID: req.ClassroomID,
		StudentID:   req.StudentID,
		TeacherID:   req.TeacherID,
	}
	member.UpdatedAt = time.Now()

	_, err := s.db.NewUpdate().
		Model(member).
		Column("classroom_id", "student_id", "teacher_id", "updated_at").
		Where("id = ?", id).
		Exec(ctx)

	if err != nil {
		return nil, err
	}

	return member, nil
}

// DeleteClassroomMember deletes a classroom member
func (s *Service) DeleteClassroomMember(ctx context.Context, id uuid.UUID) error {
	_, err := s.db.NewDelete().
		Model((*ent.ClassroomMemberEntity)(nil)).
		Where("id = ?", id).
		Exec(ctx)

	return err
}

// CheckExistClassroomMember checks if a classroom member exists
func (s *Service) CheckExistClassroomMember(ctx context.Context, id uuid.UUID) (bool, error) {
	count, err := s.db.NewSelect().
		Model((*ent.ClassroomMemberEntity)(nil)).
		Where("id = ?", id).
		Count(ctx)

	if err != nil {
		return false, err
	}
	return count > 0, nil
}

// GetClassroomMembersByStudentID retrieves all classroom memberships for a student
func (s *Service) GetClassroomMembersByStudentID(ctx context.Context, studentID uuid.UUID) ([]*ent.ClassroomMemberEntity, error) {
	var members []*ent.ClassroomMemberEntity
	err := s.db.NewSelect().
		Model(&members).
		Where("student_id = ?", studentID).
		Scan(ctx)
	if err != nil {
		return nil, err
	}
	return members, nil
}

// GetAllClassroomMembers retrieves all classroom members with limit
func (s *Service) GetAllClassroomMembers(ctx context.Context, limit int) ([]*ent.ClassroomMemberEntity, error) {
	var members []*ent.ClassroomMemberEntity
	err := s.db.NewSelect().
		Model(&members).
		Limit(limit).
		Scan(ctx)
	if err != nil {
		return nil, err
	}
	return members, nil
}

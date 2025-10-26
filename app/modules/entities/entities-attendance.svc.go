package entities

import (
	"context"
	"time"

	entitiesdto "github.com/easy-attend-serviceV3/app/modules/entities/dto"
	"github.com/easy-attend-serviceV3/app/modules/entities/ent"
	"github.com/google/uuid"
)

// CreateAttendance creates a new attendance record
func (s *Service) CreateAttendance(ctx context.Context, req *entitiesdto.AttendanceCreateRequest) (*ent.AttendanceEntity, error) {
	// Generate new UUID for attendance
	attendanceID := uuid.New()

	attendance := &ent.AttendanceEntity{
		ID:          attendanceID,
		ClassroomID: req.ClassroomID,
		TeacherID:   req.TeacherID,
		StudentID:   req.StudentID,
		Date:        req.Date,
		Time:        req.Time,
		Status:      req.Status,
	}
	attendance.CreatedAt = time.Now()
	attendance.UpdatedAt = time.Now()

	_, err := s.db.NewInsert().Model(attendance).Exec(ctx)
	if err != nil {
		return nil, err
	}
	return attendance, nil
}

// GetListAttendance retrieves attendance records by classroom and date
func (s *Service) GetListAttendance(ctx context.Context, classroomID uuid.UUID, date string) ([]*ent.AttendanceEntity, error) {
	var attendances []*ent.AttendanceEntity
	query := s.db.NewSelect().Model(&attendances).Where("classroom_id = ?", classroomID)

	if date != "" {
		query = query.Where("date = ?", date)
	}

	err := query.Scan(ctx)
	if err != nil {
		return nil, err
	}
	return attendances, nil
}

// GetAllAttendance retrieves all attendance records with limit
func (s *Service) GetAllAttendance(ctx context.Context, limit int) ([]*ent.AttendanceEntity, error) {
	var attendances []*ent.AttendanceEntity
	err := s.db.NewSelect().
		Model(&attendances).
		Limit(limit).
		Scan(ctx)
	if err != nil {
		return nil, err
	}
	return attendances, nil
}

// GetAttendanceByID retrieves an attendance record by ID
func (s *Service) GetAttendanceByID(ctx context.Context, id uuid.UUID) (*ent.AttendanceEntity, error) {
	var attendance ent.AttendanceEntity

	err := s.db.NewSelect().
		Model(&attendance).
		Where("id = ?", id).
		Scan(ctx)

	if err != nil {
		return nil, err
	}

	return &attendance, nil
}

// UpdateAttendance updates an attendance record
func (s *Service) UpdateAttendance(ctx context.Context, id uuid.UUID, req *entitiesdto.AttendanceUpdateRequest) (*ent.AttendanceEntity, error) {
	attendance := &ent.AttendanceEntity{
		ID:          id,
		ClassroomID: req.ClassroomID,
		TeacherID:   req.TeacherID,
		StudentID:   req.StudentID,
		Date:        req.Date,
		Time:        req.Time,
		Status:      req.Status,
	}
	attendance.UpdatedAt = time.Now()

	_, err := s.db.NewUpdate().
		Model(attendance).
		Column("classroom_id", "teacher_id", "student_id", "date", "time", "status", "updated_at").
		Where("id = ?", id).
		Exec(ctx)

	if err != nil {
		return nil, err
	}

	return attendance, nil
}

// DeleteAttendance deletes an attendance record
func (s *Service) DeleteAttendance(ctx context.Context, id uuid.UUID) error {
	_, err := s.db.NewDelete().
		Model((*ent.AttendanceEntity)(nil)).
		Where("id = ?", id).
		Exec(ctx)

	return err
}

// CheckExistAttendance checks if an attendance record exists
func (s *Service) CheckExistAttendance(ctx context.Context, id uuid.UUID) (bool, error) {
	count, err := s.db.NewSelect().
		Model((*ent.AttendanceEntity)(nil)).
		Where("id = ?", id).
		Count(ctx)

	if err != nil {
		return false, err
	}
	return count > 0, nil
}

// GetAttendanceByStudentID retrieves attendance records for a specific student on a specific date
func (s *Service) GetAttendanceByStudentID(ctx context.Context, studentID uuid.UUID, date string) (*ent.AttendanceEntity, error) {
	var attendance ent.AttendanceEntity

	query := s.db.NewSelect().Model(&attendance).Where("student_id = ?", studentID)
	if date != "" {
		query = query.Where("date = ?", date)
	}

	err := query.Scan(ctx)
	if err != nil {
		return nil, err
	}

	return &attendance, nil
}

// GetAttendanceByClassroomAndDate retrieves all attendance records for a classroom on a specific date
func (s *Service) GetAttendanceByClassroomAndDate(ctx context.Context, classroomID uuid.UUID, date string) ([]*ent.AttendanceEntity, error) {
	var attendances []*ent.AttendanceEntity
	err := s.db.NewSelect().
		Model(&attendances).
		Where("classroom_id = ? AND date = ?", classroomID, date).
		Scan(ctx)
	if err != nil {
		return nil, err
	}
	return attendances, nil
}

package attendance

import (
	"context"

	entitiesdto "github.com/easy-attend-serviceV3/app/modules/entities/dto"
	"github.com/easy-attend-serviceV3/app/utils"
	"github.com/google/uuid"
)

type UpdateServiceRequest struct {
	ID          uuid.UUID `json:"id"`
	ClassroomID uuid.UUID `json:"classroom_id" binding:"required,uuid"`
	TeacherID   uuid.UUID `json:"teacher_id" binding:"required,uuid"`
	StudentID   uuid.UUID `json:"student_id" binding:"required,uuid"`
	Date        string    `json:"date" binding:"required"`
	Time        string    `json:"time" binding:"required"`
	Status      string    `json:"status" binding:"required"`
}

type UpdateServiceResponse struct {
	ID          uuid.UUID `json:"id"`
	ClassroomID uuid.UUID `json:"classroom_id"`
	TeacherID   uuid.UUID `json:"teacher_id"`
	StudentID   uuid.UUID `json:"student_id"`
	Date        string    `json:"date"`
	Time        string    `json:"time"`
	Status      string    `json:"status"`
}

func (s *Service) UpdateService(ctx context.Context, req *UpdateServiceRequest) (*UpdateServiceResponse, error) {
	span, log := utils.LogSpanFromContext(ctx)
	span.AddEvent(`attendance.svc.update.start`)

	attendance, err := s.db.UpdateAttendance(ctx, req.ID, &entitiesdto.AttendanceUpdateRequest{
		ID:          req.ID,
		ClassroomID: req.ClassroomID,
		TeacherID:   req.TeacherID,
		StudentID:   req.StudentID,
		Date:        req.Date,
		Time:        req.Time,
		Status:      req.Status,
	})
	if err != nil {
		log.Error(err)
		return nil, err
	}

	response := &UpdateServiceResponse{
		ID:          attendance.ID,
		ClassroomID: attendance.ClassroomID,
		TeacherID:   attendance.TeacherID,
		StudentID:   attendance.StudentID,
		Date:        attendance.Date,
		Time:        attendance.Time,
		Status:      attendance.Status,
	}

	span.AddEvent(`attendance.svc.update.end`)
	return response, nil
}

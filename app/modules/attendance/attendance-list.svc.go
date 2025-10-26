package attendance

import (
	"context"

	"github.com/easy-attend-serviceV3/app/utils"
	"github.com/google/uuid"
)

type ListServiceRequest struct {
	ClassroomID *uuid.UUID `json:"classroom_id,omitempty"`
	StudentID   *uuid.UUID `json:"student_id,omitempty"`
	Date        *string    `json:"date,omitempty"`
	UserID      uuid.UUID  `json:"-"` // Teacher ID from token context
}

type ListServiceResponse struct {
	ID          uuid.UUID `json:"id"`
	ClassroomID uuid.UUID `json:"classroom_id"`
	TeacherID   uuid.UUID `json:"teacher_id"`
	StudentID   uuid.UUID `json:"student_id"`
	Date        string    `json:"date"`
	Time        string    `json:"time"`
	Status      string    `json:"status"`
}

func (s *Service) ListService(ctx context.Context, req *ListServiceRequest) ([]*ListServiceResponse, error) {
	span, log := utils.LogSpanFromContext(ctx)
	span.AddEvent(`attendance.svc.list.start`)

	var attendances []*ListServiceResponse

	// Get all attendance records for this teacher
	dbAttendances, err := s.db.GetAttendanceByTeacherID(ctx, req.UserID)
	if err != nil {
		log.Error(err)
		return nil, err
	}

	// Filter by additional criteria if provided
	for _, attendance := range dbAttendances {
		// Apply classroom filter if specified
		if req.ClassroomID != nil && attendance.ClassroomID != *req.ClassroomID {
			continue
		}

		// Apply student filter if specified
		if req.StudentID != nil && attendance.StudentID != *req.StudentID {
			continue
		}

		// Apply date filter if specified
		if req.Date != nil && attendance.Date != *req.Date {
			continue
		}

		attendances = append(attendances, &ListServiceResponse{
			ID:          attendance.ID,
			ClassroomID: attendance.ClassroomID,
			TeacherID:   attendance.TeacherID,
			StudentID:   attendance.StudentID,
			Date:        attendance.Date,
			Time:        attendance.Time,
			Status:      attendance.Status,
		})
	}

	span.AddEvent(`attendance.svc.list.end`)
	return attendances, nil
}

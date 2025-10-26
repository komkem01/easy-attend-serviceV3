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

	if req.ClassroomID != nil {
		// Get attendance by classroom and optional date
		date := ""
		if req.Date != nil {
			date = *req.Date
		}

		dbAttendances, err := s.db.GetListAttendance(ctx, *req.ClassroomID, date)
		if err != nil {
			log.Error(err)
			return nil, err
		}

		for _, attendance := range dbAttendances {
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
	} else if req.StudentID != nil {
		// Get attendance by student and optional date
		date := ""
		if req.Date != nil {
			date = *req.Date
		}

		dbAttendance, err := s.db.GetAttendanceByStudentID(ctx, *req.StudentID, date)
		if err != nil {
			log.Error(err)
			// If no records found, return empty array instead of error
			return []*ListServiceResponse{}, nil
		}

		attendances = append(attendances, &ListServiceResponse{
			ID:          dbAttendance.ID,
			ClassroomID: dbAttendance.ClassroomID,
			TeacherID:   dbAttendance.TeacherID,
			StudentID:   dbAttendance.StudentID,
			Date:        dbAttendance.Date,
			Time:        dbAttendance.Time,
			Status:      dbAttendance.Status,
		})
	} else {
		// Get all attendance with limit
		dbAttendances, err := s.db.GetAllAttendance(ctx, 100)
		if err != nil {
			log.Error(err)
			return nil, err
		}

		for _, attendance := range dbAttendances {
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
	}

	span.AddEvent(`attendance.svc.list.end`)
	return attendances, nil
}

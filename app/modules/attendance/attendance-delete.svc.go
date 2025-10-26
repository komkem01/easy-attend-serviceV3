package attendance

import (
	"context"

	"github.com/easy-attend-serviceV3/app/utils"
	"github.com/google/uuid"
)

type DeleteServiceRequest struct {
	ID uuid.UUID `json:"id" binding:"required,uuid"`
}

type DeleteServiceResponse struct {
	ID      uuid.UUID `json:"id"`
	Message string    `json:"message"`
}

func (s *Service) DeleteService(ctx context.Context, req *DeleteServiceRequest) (*DeleteServiceResponse, error) {
	span, log := utils.LogSpanFromContext(ctx)
	span.AddEvent(`attendance.svc.delete.start`)

	err := s.db.DeleteAttendance(ctx, req.ID)
	if err != nil {
		log.Error(err)
		return nil, err
	}

	response := &DeleteServiceResponse{
		ID:      req.ID,
		Message: "Attendance record deleted successfully",
	}

	span.AddEvent(`attendance.svc.delete.end`)
	return response, nil
}

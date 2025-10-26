package classroommember

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
	span.AddEvent(`classroom_member.svc.delete.start`)

	err := s.db.DeleteClassroomMember(ctx, req.ID)
	if err != nil {
		log.Error(err)
		return nil, err
	}

	response := &DeleteServiceResponse{
		ID:      req.ID,
		Message: "Classroom member deleted successfully",
	}

	span.AddEvent(`classroom_member.svc.delete.end`)
	return response, nil
}

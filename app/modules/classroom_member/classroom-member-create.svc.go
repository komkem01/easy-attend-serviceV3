package classroommember

import (
	"context"

	entitiesdto "github.com/easy-attend-serviceV3/app/modules/entities/dto"
	"github.com/easy-attend-serviceV3/app/utils"
	"github.com/google/uuid"
)

type CreateServiceRequest struct {
	ClassroomID uuid.UUID `json:"classroom_id" binding:"required,uuid"`
	TeacherID   uuid.UUID `json:"teacher_id" binding:"required,uuid"`
	StudentID   uuid.UUID `json:"student_id" binding:"required,uuid"`
}

type CreateServiceResponse struct {
	ID          uuid.UUID `json:"id"`
	ClassroomID uuid.UUID `json:"classroom_id"`
	TeacherID   uuid.UUID `json:"teacher_id"`
	StudentID   uuid.UUID `json:"student_id"`
}

func (s *Service) CreateService(ctx context.Context, req *CreateServiceRequest) (*CreateServiceResponse, error) {
	span, log := utils.LogSpanFromContext(ctx)
	span.AddEvent(`classroom_member.svc.create.start`)

	member, err := s.db.CreateClassroomMember(ctx, &entitiesdto.ClassroomMemberCreateRequest{
		ClassroomID: req.ClassroomID,
		TeacherID:   req.TeacherID,
		StudentID:   req.StudentID,
	})
	if err != nil {
		log.Error(err)
		return nil, err
	}

	response := &CreateServiceResponse{
		ID:          member.ID,
		ClassroomID: member.ClassroomID,
		TeacherID:   member.TeacherID,
		StudentID:   member.StudentID,
	}

	span.AddEvent(`classroom_member.svc.create.end`)
	return response, nil
}

package classroommember

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
}

type UpdateServiceResponse struct {
	ID          uuid.UUID `json:"id"`
	ClassroomID uuid.UUID `json:"classroom_id"`
	TeacherID   uuid.UUID `json:"teacher_id"`
	StudentID   uuid.UUID `json:"student_id"`
}

func (s *Service) UpdateService(ctx context.Context, req *UpdateServiceRequest) (*UpdateServiceResponse, error) {
	span, log := utils.LogSpanFromContext(ctx)
	span.AddEvent(`classroom_member.svc.update.start`)

	member, err := s.db.UpdateClassroomMember(ctx, req.ID, &entitiesdto.ClassroomMemberUpdateRequest{
		ID:          req.ID,
		ClassroomID: req.ClassroomID,
		TeacherID:   req.TeacherID,
		StudentID:   req.StudentID,
	})
	if err != nil {
		log.Error(err)
		return nil, err
	}

	response := &UpdateServiceResponse{
		ID:          member.ID,
		ClassroomID: member.ClassroomID,
		TeacherID:   member.TeacherID,
		StudentID:   member.StudentID,
	}

	span.AddEvent(`classroom_member.svc.update.end`)
	return response, nil
}

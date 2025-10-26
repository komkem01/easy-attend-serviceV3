package classroommember

import (
	"context"

	"github.com/easy-attend-serviceV3/app/utils"
	"github.com/google/uuid"
)

type ListServiceRequest struct {
	ClassroomID *uuid.UUID `json:"classroom_id,omitempty"`
	StudentID   *uuid.UUID `json:"student_id,omitempty"`
}

type ListServiceResponse struct {
	ID          uuid.UUID `json:"id"`
	ClassroomID uuid.UUID `json:"classroom_id"`
	TeacherID   uuid.UUID `json:"teacher_id"`
	StudentID   uuid.UUID `json:"student_id"`
}

func (s *Service) ListService(ctx context.Context, req *ListServiceRequest) ([]*ListServiceResponse, error) {
	span, log := utils.LogSpanFromContext(ctx)
	span.AddEvent(`classroom_member.svc.list.start`)

	var members []*ListServiceResponse

	// If classroom ID is provided, get members by classroom
	if req.ClassroomID != nil {
		dbMembers, dbErr := s.db.GetListClassroomMember(ctx, *req.ClassroomID)
		if dbErr != nil {
			log.Error(dbErr)
			return nil, dbErr
		}

		for _, member := range dbMembers {
			members = append(members, &ListServiceResponse{
				ID:          member.ID,
				ClassroomID: member.ClassroomID,
				TeacherID:   member.TeacherID,
				StudentID:   member.StudentID,
			})
		}
	} else if req.StudentID != nil {
		// If student ID is provided, get memberships by student
		dbMembers, dbErr := s.db.GetClassroomMembersByStudentID(ctx, *req.StudentID)
		if dbErr != nil {
			log.Error(dbErr)
			return nil, dbErr
		}

		for _, member := range dbMembers {
			members = append(members, &ListServiceResponse{
				ID:          member.ID,
				ClassroomID: member.ClassroomID,
				TeacherID:   member.TeacherID,
				StudentID:   member.StudentID,
			})
		}
	} else {
		// If no filter provided, get all classroom members with limit
		// This prevents performance issues with large datasets
		dbMembers, dbErr := s.db.GetAllClassroomMembers(ctx, 100) // limit to 100 records
		if dbErr != nil {
			log.Error(dbErr)
			return nil, dbErr
		}

		for _, member := range dbMembers {
			members = append(members, &ListServiceResponse{
				ID:          member.ID,
				ClassroomID: member.ClassroomID,
				TeacherID:   member.TeacherID,
				StudentID:   member.StudentID,
			})
		}
	}

	span.AddEvent(`classroom_member.svc.list.end`)
	return members, nil
}

package classroommember

import (
	"context"

	"github.com/easy-attend-serviceV3/app/utils"
	"github.com/google/uuid"
)

type ListServiceRequest struct {
	ClassroomID *uuid.UUID `json:"classroom_id,omitempty"`
	StudentID   *uuid.UUID `json:"student_id,omitempty"`
	UserID      uuid.UUID  `json:"-"` // Teacher ID from token context
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

	// If classroom ID is provided, get members by classroom (filtered by teacher)
	if req.ClassroomID != nil {
		// First verify if classroom exists
		_, err := s.classroomDB.GetByIDClassroom(ctx, *req.ClassroomID)
		if err != nil {
			log.Errf("Failed to get classroom: %s", err)
			return nil, err
		}

		// Check if teacher is associated with this classroom via classroom_members
		teacherMembers, err := s.db.GetClassroomMembersByTeacherID(ctx, req.UserID)
		if err != nil {
			log.Errf("Failed to check teacher access: %s", err)
			return nil, err
		}

		// Verify teacher has access to this classroom
		hasAccess := false
		for _, member := range teacherMembers {
			if member.ClassroomID == *req.ClassroomID {
				hasAccess = true
				break
			}
		}

		if !hasAccess {
			log.Infof("Teacher %s does not have access to classroom %s", req.UserID, *req.ClassroomID)
			return []*ListServiceResponse{}, nil // Return empty list
		}

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

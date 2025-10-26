package classroommember

import (
	"context"

	"github.com/easy-attend-serviceV3/app/utils"
	"github.com/google/uuid"
)

type InfoServiceRequest struct {
	ID uuid.UUID `json:"id" binding:"required,uuid"`
}

type InfoServiceResponse struct {
	ID            uuid.UUID   `json:"id"`
	ClassroomID   uuid.UUID   `json:"classroom_id"`
	ClassroomName string      `json:"classroom_name"`
	School        SchoolInfo  `json:"school_info"`
	TeacherID     uuid.UUID   `json:"teacher_id"`
	TeacherName   string      `json:"teacher_name"`
	Student       StudentInfo `json:"student_info"`
}

type SchoolInfo struct {
	ID   uuid.UUID `json:"id"`
	Name string    `json:"name"`
}

type StudentInfo struct {
	ID          uuid.UUID `json:"id"`
	FirstName   string    `json:"first_name"`
	LastName    string    `json:"last_name"`
	StudentCode string    `json:"student_code"`
	Phone       string    `json:"phone"`
}

func (s *Service) InfoService(ctx context.Context, req *InfoServiceRequest) (*InfoServiceResponse, error) {
	span, log := utils.LogSpanFromContext(ctx)
	span.AddEvent(`classroom_member.svc.info.start`)

	member, err := s.db.GetClassroomMemberByID(ctx, req.ID)
	if err != nil {
		log.Error(err)
		return nil, err
	}

	// Get classroom data
	classroom, err := s.classroomDB.GetByIDClassroom(ctx, member.ClassroomID)
	if err != nil {
		log.Error(err)
		return nil, err
	}

	// Get school data
	school, err := s.schoolDB.GetByIDSchool(ctx, classroom.SchoolID)
	if err != nil {
		log.Error(err)
		return nil, err
	}

	// Get teacher data
	teacher, err := s.teacherDB.GetByIDTeacher(ctx, member.TeacherID)
	if err != nil {
		log.Error(err)
		return nil, err
	}

	// Get student data
	student, err := s.studentDB.GetStudentByID(ctx, member.StudentID, nil)
	if err != nil {
		log.Error(err)
		return nil, err
	}

	response := &InfoServiceResponse{
		ID:            member.ID,
		ClassroomID:   member.ClassroomID,
		ClassroomName: classroom.Name,
		School: SchoolInfo{
			ID:   school.ID,
			Name: school.Name,
		},
		TeacherID:   member.TeacherID,
		TeacherName: teacher.FirstName + " " + teacher.LastName,
		Student: StudentInfo{
			ID:          student.ID,
			FirstName:   student.FirstName,
			LastName:    student.LastName,
			StudentCode: student.StudentCode,
			Phone:       student.Phone,
		},
	}

	span.AddEvent(`classroom_member.svc.info.end`)
	return response, nil
}

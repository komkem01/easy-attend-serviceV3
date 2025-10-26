package attendance

import (
	"context"

	"github.com/easy-attend-serviceV3/app/utils"
	"github.com/google/uuid"
)

type InfoServiceRequest struct {
	ID uuid.UUID `json:"id" binding:"required,uuid"`
}

type InfoServiceResponse struct {
	ID        uuid.UUID       `json:"id"`
	Date      string          `json:"date"`
	Time      string          `json:"time"`
	Status    string          `json:"status"`
	Classroom ClassroomDetail `json:"classroom"`
	Teacher   TeacherDetail   `json:"teacher"`
	Student   StudentDetail   `json:"student"`
}

type ClassroomDetail struct {
	ID     uuid.UUID    `json:"id"`
	Name   string       `json:"name"`
	School SchoolDetail `json:"school"`
}

type SchoolDetail struct {
	ID      uuid.UUID `json:"id"`
	Name    string    `json:"name"`
	Address string    `json:"address"`
	Phone   string    `json:"phone"`
}

type TeacherDetail struct {
	ID        uuid.UUID `json:"id"`
	FirstName string    `json:"first_name"`
	LastName  string    `json:"last_name"`
	FullName  string    `json:"full_name"`
	Email     string    `json:"email"`
	Phone     string    `json:"phone"`
}

type StudentDetail struct {
	ID          uuid.UUID `json:"id"`
	FirstName   string    `json:"first_name"`
	LastName    string    `json:"last_name"`
	FullName    string    `json:"full_name"`
	StudentCode string    `json:"student_code"`
	Phone       string    `json:"phone"`
}

func (s *Service) InfoService(ctx context.Context, req *InfoServiceRequest) (*InfoServiceResponse, error) {
	span, log := utils.LogSpanFromContext(ctx)
	span.AddEvent(`attendance.svc.info.start`)

	// Get attendance record
	attendance, err := s.db.GetAttendanceByID(ctx, req.ID)
	if err != nil {
		log.Error(err)
		return nil, err
	}

	// Get classroom data
	classroom, err := s.classroomDB.GetByIDClassroom(ctx, attendance.ClassroomID)
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
	teacher, err := s.teacherDB.GetByIDTeacher(ctx, attendance.TeacherID)
	if err != nil {
		log.Error(err)
		return nil, err
	}

	// Get student data
	student, err := s.studentDB.GetStudentByID(ctx, attendance.StudentID, nil)
	if err != nil {
		log.Error(err)
		return nil, err
	}

	response := &InfoServiceResponse{
		ID:     attendance.ID,
		Date:   attendance.Date,
		Time:   attendance.Time,
		Status: attendance.Status,
		Classroom: ClassroomDetail{
			ID:   classroom.ID,
			Name: classroom.Name,
			School: SchoolDetail{
				ID:      school.ID,
				Name:    school.Name,
				Address: school.Address,
				Phone:   school.Phone,
			},
		},
		Teacher: TeacherDetail{
			ID:        teacher.ID,
			FirstName: teacher.FirstName,
			LastName:  teacher.LastName,
			FullName:  teacher.FirstName + " " + teacher.LastName,
			Email:     teacher.Email,
			Phone:     teacher.Phone,
		},
		Student: StudentDetail{
			ID:          student.ID,
			FirstName:   student.FirstName,
			LastName:    student.LastName,
			FullName:    student.FirstName + " " + student.LastName,
			StudentCode: student.StudentCode,
			Phone:       student.Phone,
		},
	}

	span.AddEvent(`attendance.svc.info.end`)
	return response, nil
}

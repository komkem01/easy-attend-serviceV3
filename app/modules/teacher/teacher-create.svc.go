package teacher

import (
	"context"
	"fmt"
	"time"

	entitiesdto "github.com/easy-attend-serviceV3/app/modules/entities/dto"
	"github.com/easy-attend-serviceV3/app/utils"
	"github.com/easy-attend-serviceV3/app/utils/auth"
	"github.com/google/uuid"
)

type CreateServiceRequest struct {
	SchoolName  string     `json:"school_name"`            // ชื่อโรงเรียน (ระบบจะหาหรือสร้างใหม่)
	ClassroomID *uuid.UUID `json:"classroom_id,omitempty"` // ไม่บังคับกรอก
	PrefixID    uuid.UUID  `json:"prefix_id"`
	GenderID    uuid.UUID  `json:"gender_id"`
	FirstName   string     `json:"first_name"`
	LastName    string     `json:"last_name"`
	Email       string     `json:"email"`
	Password    string     `json:"password"`
	Phone       string     `json:"phone"`
}

type CreateServiceResponse struct {
	ID           uuid.UUID  `json:"id"`
	SchoolID     uuid.UUID  `json:"school_id"`
	SchoolName   string     `json:"school_name"`            // เพิ่มชื่อโรงเรียน
	ClassroomID  *uuid.UUID `json:"classroom_id,omitempty"` // อาจเป็น null
	PrefixID     uuid.UUID  `json:"prefix_id"`
	GenderID     uuid.UUID  `json:"gender_id"`
	FirstName    string     `json:"first_name"`
	LastName     string     `json:"last_name"`
	Email        string     `json:"email"`
	Phone        string     `json:"phone"`
	AccessToken  string     `json:"access_token,omitempty"`
	RefreshToken string     `json:"refresh_token,omitempty"`
	TokenType    string     `json:"token_type,omitempty"`
	ExpiresAt    time.Time  `json:"expires_at,omitempty"`
}

func (s *Service) CreateService(ctx context.Context, req *CreateServiceRequest) (*CreateServiceResponse, error) {
	span, log := utils.LogSpanFromContext(ctx)
	span.AddEvent(`teacher.svc.create.start`)

	// Hash password
	passwordHasher := auth.NewPasswordHasher()
	hashedPassword, err := passwordHasher.HashPassword(req.Password)
	if err != nil {
		log.Error(err)
		return nil, err
	}

	// Find or create school by name
	school, err := s.dbSchool.FindOrCreateSchoolByName(ctx, req.SchoolName)
	if err != nil {
		log.Error(err)
		return nil, err
	}

	// Validate prefix_id exists
	_, err = s.dbPrefix.GetByIDPrefix(ctx, req.PrefixID)
	if err != nil {
		log.Error(err)
		return nil, fmt.Errorf("prefix not found: %s", req.PrefixID.String())
	}

	// Validate gender_id exists
	_, err = s.dbGender.GetByIDGender(ctx, req.GenderID)
	if err != nil {
		log.Error(err)
		return nil, fmt.Errorf("gender not found: %s", req.GenderID.String())
	}

	// Check if email already exists
	existingTeacher, err := s.db.GetTeacherByEmail(ctx, req.Email)
	if err == nil && existingTeacher != nil {
		return nil, fmt.Errorf("email already exists: %s", req.Email)
	}

	// Handle optional classroom_id with validation
	var classroomIDPtr *uuid.UUID
	if req.ClassroomID != nil {
		// Validate that classroom exists before using it
		_, err := s.dbClassroom.GetByIDClassroom(ctx, *req.ClassroomID)
		if err != nil {
			log.Error(err)
			return nil, fmt.Errorf("classroom not found: %s", req.ClassroomID.String())
		}
		classroomIDPtr = req.ClassroomID // ส่ง pointer ตรงๆ
	}
	// ถ้าไม่มีค่า classroomIDPtr จะเป็น nil

	teacher, err := s.db.CreateTeacher(ctx, &entitiesdto.TeacherCreateRequest{
		SchoolID:    school.ID,      // Use found/created school ID
		ClassroomID: classroomIDPtr, // ส่ง pointer ที่อาจเป็น nil
		PrefixID:    req.PrefixID,
		GenderID:    req.GenderID,
		FirstName:   req.FirstName,
		LastName:    req.LastName,
		Email:       req.Email,
		Password:    hashedPassword, // Use hashed password
		Phone:       req.Phone,
	})
	if err != nil {
		log.Error(err)
		return nil, err
	}

	// Generate tokens for the new teacher
	tokenMgr := auth.NewTokenManager("your-secret-key-here") // TODO: Move to config
	tokens, err := tokenMgr.GenerateTokenPair(
		teacher.ID,
		teacher.Email,
		teacher.FirstName,
		teacher.LastName,
		"teacher",
	)
	if err != nil {
		log.Error(err)
		// Don't fail the registration if token generation fails
		// Log the error and continue
	}

	// Handle optional classroom_id for response
	var responseClassroomID *uuid.UUID
	if teacher.ClassroomID != nil {
		responseClassroomID = teacher.ClassroomID // teacher.ClassroomID เป็น pointer อยู่แล้ว
	}

	response := &CreateServiceResponse{
		ID:          teacher.ID,
		SchoolID:    teacher.SchoolID,
		SchoolName:  school.Name,         // เพิ่มชื่อโรงเรียน
		ClassroomID: responseClassroomID, // Handle optional classroom ID
		PrefixID:    teacher.PrefixID,
		GenderID:    teacher.GenderID,
		FirstName:   teacher.FirstName,
		LastName:    teacher.LastName,
		Email:       teacher.Email,
		Phone:       teacher.Phone,
	}

	// Add token information if generated successfully
	if tokens != nil {
		response.AccessToken = tokens.AccessToken
		response.RefreshToken = tokens.RefreshToken
		response.TokenType = tokens.TokenType
		response.ExpiresAt = tokens.ExpiresAt
	}

	span.AddEvent(`teacher.svc.create.end`)
	return response, nil
}

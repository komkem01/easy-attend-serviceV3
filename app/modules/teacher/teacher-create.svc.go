package teacher

import (
	"context"
	"time"

	entitiesdto "github.com/easy-attend-serviceV3/app/modules/entities/dto"
	"github.com/easy-attend-serviceV3/app/utils"
	"github.com/easy-attend-serviceV3/app/utils/auth"
	"github.com/google/uuid"
)

type CreateServiceRequest struct {
	SchoolID    uuid.UUID `json:"school_id"`
	ClassroomID uuid.UUID `json:"classroom_id"`
	PrefixID    uuid.UUID `json:"prefix_id"`
	GenderID    uuid.UUID `json:"gender_id"`
	FirstName   string    `json:"first_name"`
	LastName    string    `json:"last_name"`
	Email       string    `json:"email"`
	Password    string    `json:"password"`
	Phone       string    `json:"phone"`
}

type CreateServiceResponse struct {
	ID           uuid.UUID `json:"id"`
	SchoolID     uuid.UUID `json:"school_id"`
	ClassroomID  uuid.UUID `json:"classroom_id"`
	PrefixID     uuid.UUID `json:"prefix_id"`
	GenderID     uuid.UUID `json:"gender_id"`
	FirstName    string    `json:"first_name"`
	LastName     string    `json:"last_name"`
	Email        string    `json:"email"`
	Phone        string    `json:"phone"`
	AccessToken  string    `json:"access_token,omitempty"`
	RefreshToken string    `json:"refresh_token,omitempty"`
	TokenType    string    `json:"token_type,omitempty"`
	ExpiresAt    time.Time `json:"expires_at,omitempty"`
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

	teacher, err := s.db.CreateTeacher(ctx, &entitiesdto.TeacherCreateRequest{
		SchoolID:    req.SchoolID,
		ClassroomID: req.ClassroomID,
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

	response := &CreateServiceResponse{
		ID:          teacher.ID,
		SchoolID:    teacher.SchoolID,
		ClassroomID: teacher.ClassroomID,
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

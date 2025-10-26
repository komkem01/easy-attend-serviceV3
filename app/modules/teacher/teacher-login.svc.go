package teacher

import (
	"context"
	"errors"
	"time"

	"github.com/easy-attend-serviceV3/app/utils"
	"github.com/easy-attend-serviceV3/app/utils/auth"
	"github.com/google/uuid"
)

type LoginServiceRequest struct {
	Email    string `json:"email" binding:"required,email"`
	Password string `json:"password" binding:"required"`
}

type LoginServiceResponse struct {
	ID           uuid.UUID  `json:"id"`
	SchoolID     uuid.UUID  `json:"school_id"`
	ClassroomID  *uuid.UUID `json:"classroom_id,omitempty"` // ใช้ pointer เพื่อรองรับ NULL
	PrefixID     uuid.UUID  `json:"prefix_id"`
	GenderID     uuid.UUID  `json:"gender_id"`
	FirstName    string     `json:"first_name"`
	LastName     string     `json:"last_name"`
	Email        string     `json:"email"`
	Phone        string     `json:"phone"`
	AccessToken  string     `json:"access_token"`
	RefreshToken string     `json:"refresh_token"`
	TokenType    string     `json:"token_type"`
	ExpiresAt    time.Time  `json:"expires_at"`
}

func (s *Service) LoginService(ctx context.Context, req *LoginServiceRequest) (*LoginServiceResponse, error) {
	span, log := utils.LogSpanFromContext(ctx)
	span.AddEvent(`teacher.svc.login.start`)

	// Get teacher by email
	teacher, err := s.db.GetTeacherByEmail(ctx, req.Email)
	if err != nil {
		log.Error(err)
		return nil, errors.New("invalid email or password")
	}

	// Verify password
	passwordHasher := auth.NewPasswordHasher()
	valid, err := passwordHasher.VerifyPassword(req.Password, teacher.Password)
	if err != nil {
		log.Error(err)
		return nil, errors.New("invalid email or password")
	}

	if !valid {
		return nil, errors.New("invalid email or password")
	}

	// Generate tokens
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
		return nil, errors.New("failed to generate authentication tokens")
	}

	response := &LoginServiceResponse{
		ID:           teacher.ID,
		SchoolID:     teacher.SchoolID,
		ClassroomID:  teacher.ClassroomID, // teacher.ClassroomID เป็น pointer อยู่แล้ว
		PrefixID:     teacher.PrefixID,
		GenderID:     teacher.GenderID,
		FirstName:    teacher.FirstName,
		LastName:     teacher.LastName,
		Email:        teacher.Email,
		Phone:        teacher.Phone,
		AccessToken:  tokens.AccessToken,
		RefreshToken: tokens.RefreshToken,
		TokenType:    tokens.TokenType,
		ExpiresAt:    tokens.ExpiresAt,
	}

	span.AddEvent(`teacher.svc.login.end`)
	return response, nil
}

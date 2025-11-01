package teacher

import (
	"context"
	"errors"
	"time"

	"github.com/easy-attend-serviceV3/app/utils"
	"github.com/easy-attend-serviceV3/app/utils/auth"
)

type RefreshServiceRequest struct {
	RefreshToken string `json:"refresh_token" binding:"required"`
}

type RefreshServiceResponse struct {
	AccessToken  string    `json:"access_token"`
	RefreshToken string    `json:"refresh_token"`
	TokenType    string    `json:"token_type"`
	ExpiresAt    time.Time `json:"expires_at"`
}

func (s *Service) RefreshService(ctx context.Context, req *RefreshServiceRequest) (*RefreshServiceResponse, error) {
	span, log := utils.LogSpanFromContext(ctx)
	span.AddEvent(`teacher.svc.refresh.start`)

	// Initialize token manager with config
	tokenMgr := auth.NewTokenManagerWithConfig(
		s.config.JWT.SecretKey,
		time.Duration(s.config.JWT.AccessTokenExpiry)*time.Hour,
		time.Duration(s.config.JWT.RefreshTokenExpiry)*time.Hour,
	)

	// Refresh the token
	newTokens, err := tokenMgr.RefreshToken(req.RefreshToken)
	if err != nil {
		log.Error(err)
		return nil, errors.New("invalid or expired refresh token")
	}

	response := &RefreshServiceResponse{
		AccessToken:  newTokens.AccessToken,
		RefreshToken: newTokens.RefreshToken,
		TokenType:    newTokens.TokenType,
		ExpiresAt:    newTokens.ExpiresAt,
	}

	span.AddEvent(`teacher.svc.refresh.end`)
	return response, nil
}
package teacher

import (
	"github.com/easy-attend-serviceV3/app/utils"
	"github.com/easy-attend-serviceV3/app/utils/base"
	"github.com/easy-attend-serviceV3/config/i18n"
	"github.com/gin-gonic/gin"
)

type RefreshControllerRequest struct {
	RefreshToken string `json:"refresh_token" binding:"required"`
}

type RefreshControllerResponse struct {
	AccessToken  string `json:"access_token"`
	RefreshToken string `json:"refresh_token"`
	TokenType    string `json:"token_type"`
	ExpiresAt    int64  `json:"expires_at"`
}

func (c *Controller) RefreshController(ctx *gin.Context) {
	span, _ := utils.LogSpanFromContext(ctx.Request.Context())
	span.AddEvent(`teacher.refresh.ctl.start`)

	var request RefreshControllerRequest
	if err := ctx.ShouldBindJSON(&request); err != nil {
		base.BadRequest(ctx, i18n.BadRequest, nil)
		return
	}

	// Call service layer
	refreshResp, err := c.svc.RefreshService(ctx.Request.Context(), &RefreshServiceRequest{
		RefreshToken: request.RefreshToken,
	})
	if err != nil {
		base.HandleError(ctx, err)
		return
	}

	// Build response
	resp := RefreshControllerResponse{
		AccessToken:  refreshResp.AccessToken,
		RefreshToken: refreshResp.RefreshToken,
		TokenType:    refreshResp.TokenType,
		ExpiresAt:    refreshResp.ExpiresAt.Unix(),
	}

	span.AddEvent(`teacher.refresh.ctl.end`)
	base.Success(ctx, resp)
}
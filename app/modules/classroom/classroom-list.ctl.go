package classroom

import (
	"github.com/easy-attend-serviceV3/app/utils"
	"github.com/easy-attend-serviceV3/app/utils/auth"
	"github.com/easy-attend-serviceV3/app/utils/base"
	"github.com/easy-attend-serviceV3/config/i18n"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

type ListControllerRequest struct {
	base.RequestPaginate
}

type ListControllerResponse struct {
	ID        uuid.UUID `json:"id"`
	Name      string    `json:"name"`
	SchoolID  uuid.UUID `json:"school_id"`
	CreatedAt int64     `json:"created_at"`
	UpdatedAt int64     `json:"updated_at"`
}

func (c *Controller) ListController(ctx *gin.Context) {
	span, _ := utils.LogSpanFromGin(ctx)

	var req ListControllerRequest
	if err := ctx.ShouldBind(&req); err != nil {
		base.BadRequest(ctx, i18n.BadRequest, nil)
		return
	}
	span.AddEvent(`prefix.ctl.list.request`)

	// Get user ID from token context
	userID, err := auth.GetUserID(ctx)
	if err != nil {
		base.Unauthorized(ctx, "User not authenticated", nil)
		return
	}

	data, _, err := c.svc.ListService(ctx, &ListServiceRequest{
		RequestPaginate: req.RequestPaginate,
		UserID:          userID,
	})
	if err != nil {
		base.HandleError(ctx, err)
		return
	}
	span.AddEvent(`prefix.ctl.list.callsvc`)

	var resp []*ListControllerResponse
	if err := utils.CopyNTimeToUnix(&resp, data); err != nil {
		base.InternalServerError(ctx, err.Error(), nil)
		return
	}

	base.Paginate(ctx, resp, nil)
}

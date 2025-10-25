package classroom

import (
	"github.com/easy-attend-serviceV3/app/utils"
	"github.com/easy-attend-serviceV3/app/utils/base"
	"github.com/easy-attend-serviceV3/config/i18n"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

type InfoControllerRequest struct {
	ID string `uri:"id"`
}

type InfoControllerResponse struct {
	ID        uuid.UUID `json:"id"`
	Name      string    `json:"name"`
	SchoolID  uuid.UUID `json:"school_id"`
	CreatedAt int64     `json:"created_at"`
	UpdatedAt int64     `json:"updated_at"`
}

func (c *Controller) InfoController(ctx *gin.Context) {
	span, _ := utils.LogSpanFromGin(ctx)

	var req InfoControllerRequest
	if err := ctx.ShouldBindUri(&req); err != nil {
		base.BadRequest(ctx, i18n.BadRequest, nil)
		return
	}
	span.AddEvent(`prefix.ctl.info.request`)

	id, err := uuid.Parse(req.ID)
	if err != nil {
		base.BadRequest(ctx, i18n.BadRequest, nil)
		return
	}
	data, err := c.svc.InfoService(ctx, id)
	if err != nil {
		base.HandleError(ctx, err)
		return
	}
	var resp InfoControllerResponse
	span.AddEvent(`prefix.ctl.info.callsvc`)
	if err := utils.CopyNTimeToUnix(&resp, data); err != nil {
		base.InternalServerError(ctx, err.Error(), nil)
		return
	}

	base.Success(ctx, resp)
}

package prefix

import (
	"github.com/easy-attend-serviceV3/app/utils"
	"github.com/easy-attend-serviceV3/app/utils/base"
	"github.com/easy-attend-serviceV3/config/i18n"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

type ListControllerRequest struct {
	base.RequestPaginate
}

type ListControllerResponse struct {
	ID   uuid.UUID `json:"id"`
	Name string    `json:"name"`
}

func (c *Controller) ListController(ctx *gin.Context) {
	span, _ := utils.LogSpanFromGin(ctx)

	var req ListControllerRequest
	if err := ctx.ShouldBind(&req); err != nil {
		base.BadRequest(ctx, i18n.BadRequest, nil)
		return
	}
	span.AddEvent(`prefix.ctl.list.request`)

	data, _, err := c.svc.ListService(ctx, &ListServiceRequest{
		RequestPaginate: req.RequestPaginate,
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

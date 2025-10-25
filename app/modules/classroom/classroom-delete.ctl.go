package classroom

import (
	"github.com/easy-attend-serviceV3/app/utils"
	"github.com/easy-attend-serviceV3/app/utils/base"
	"github.com/easy-attend-serviceV3/config/i18n"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

type DeleteControllerRequest struct {
	ID string `uri:"id"`
}

func (c *Controller) DeleteController(ctx *gin.Context) {
	span, _ := utils.LogSpanFromContext(ctx.Request.Context())
	span.AddEvent(`school.delete.ctl.start`)

	var req DeleteControllerRequest
	if err := ctx.ShouldBindUri(&req); err != nil {
		base.BadRequest(ctx, i18n.BadRequest, nil)
		return
	}
	id, err := uuid.Parse(req.ID)
	if err != nil {
		base.BadRequest(ctx, i18n.BadRequest, nil)
		return
	}
	span.AddEvent(`school.delete.ctl.request`)

	if err := c.svc.DeleteService(ctx.Request.Context(), &DeleteServiceRequest{
		ID: id,
	}); err != nil {
		base.HandleError(ctx, err)
		return
	}

	span.AddEvent(`school.delete.ctl.end`)
	base.Success(ctx, nil)
}

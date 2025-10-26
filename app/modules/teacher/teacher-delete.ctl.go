package teacher

import (
	"github.com/easy-attend-serviceV3/app/utils"
	"github.com/easy-attend-serviceV3/app/utils/base"
	"github.com/easy-attend-serviceV3/config/i18n"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

func (c *Controller) DeleteController(ctx *gin.Context) {
	span, _ := utils.LogSpanFromContext(ctx.Request.Context())
	span.AddEvent(`teacher.delete.ctl.start`)

	// Get ID from URL parameter
	idParam := ctx.Param("id")
	id, err := uuid.Parse(idParam)
	if err != nil {
		base.BadRequest(ctx, i18n.BadRequest, nil)
		return
	}
	span.AddEvent(`teacher.delete.ctl.request`)

	if err := c.svc.DeleteService(ctx.Request.Context(), &DeleteServiceRequest{
		ID: id,
	}); err != nil {
		base.HandleError(ctx, err)
		return
	}

	span.AddEvent(`teacher.delete.ctl.end`)
	base.Success(ctx, nil)
}

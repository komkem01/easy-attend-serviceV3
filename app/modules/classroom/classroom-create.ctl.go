package classroom

import (
	"github.com/easy-attend-serviceV3/app/utils"
	"github.com/easy-attend-serviceV3/app/utils/base"
	"github.com/easy-attend-serviceV3/config/i18n"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

type CreateControllerRequest struct {
	SchoolID string `json:"school_id"`
	Name     string `json:"name" binding:"required"`
}

func (c *Controller) CreateController(ctx *gin.Context) {
	span, _ := utils.LogSpanFromContext(ctx.Request.Context())
	span.AddEvent(`classroom.create.ctl.start`)

	var request CreateControllerRequest
	if err := ctx.ShouldBindJSON(&request); err != nil {
		base.BadRequest(ctx, i18n.BadRequest, nil)
		return
	}
	span.AddEvent(`classroom.create.ctl.request`)

	if request.Name == "" {
		base.BadRequest(ctx, i18n.BadRequest, nil)
		return
	}

	schoolID, err := uuid.Parse(request.SchoolID)
	if err != nil {
		base.BadRequest(ctx, i18n.BadRequest, nil)
		return
	}

	if err := c.svc.CreateService(ctx.Request.Context(), &CreateServiceRequest{
		SchoolID: schoolID,
		Name:     request.Name,
	}); err != nil {
		base.HandleError(ctx, err)
		return
	}
	span.AddEvent(`classroom.create.ctl.end`)
	base.Success(ctx, nil)
}

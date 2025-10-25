package classroom

import (
	"github.com/easy-attend-serviceV3/app/utils"
	"github.com/easy-attend-serviceV3/app/utils/base"
	"github.com/easy-attend-serviceV3/config/i18n"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

type UpdateControllerRequest struct {
	Name     string `json:"name" binding:"required"`
	SchoolID string `json:"school_id"`
}

func (c *Controller) UpdateController(ctx *gin.Context) {
	span, _ := utils.LogSpanFromContext(ctx.Request.Context())
	span.AddEvent(`school.update.ctl.start`)

	// Get ID from URL parameter
	idParam := ctx.Param("id")
	id, err := uuid.Parse(idParam)
	if err != nil {
		base.BadRequest(ctx, i18n.BadRequest, nil)
		return
	}

	var request UpdateControllerRequest
	if err := ctx.ShouldBindJSON(&request); err != nil {
		base.BadRequest(ctx, i18n.BadRequest, nil)
		return
	}
	span.AddEvent(`school.update.ctl.request`)

	// Validate required fields
	if request.Name == "" {
		base.BadRequest(ctx, i18n.BadRequest, nil)
		return
	}

	schoolID, err := uuid.Parse(request.SchoolID)
	if err != nil {
		base.BadRequest(ctx, i18n.BadRequest, nil)
		return
	}

	if err := c.svc.UpdateService(ctx.Request.Context(), &UpdateServiceRequest{
		ID:       id,
		Name:     request.Name,
		SchoolID: schoolID,
	}); err != nil {
		base.HandleError(ctx, err)
		return
	}

	span.AddEvent(`school.update.ctl.end`)
	base.Success(ctx, nil)
}

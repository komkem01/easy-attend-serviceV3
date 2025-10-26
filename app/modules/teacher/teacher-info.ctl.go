package teacher

import (
	"github.com/easy-attend-serviceV3/app/utils"
	"github.com/easy-attend-serviceV3/app/utils/base"
	"github.com/easy-attend-serviceV3/config/i18n"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

type InfoControllerResponse struct {
	ID          string `json:"id"`
	SchoolID    string `json:"school_id"`
	ClassroomID string `json:"classroom_id"`
	PrefixID    string `json:"prefix_id"`
	GenderID    string `json:"gender_id"`
	FirstName   string `json:"first_name"`
	LastName    string `json:"last_name"`
	Email       string `json:"email"`
	Phone       string `json:"phone"`
	CreatedAt   int64  `json:"created_at"`
	UpdatedAt   int64  `json:"updated_at"`
}

func (c *Controller) InfoController(ctx *gin.Context) {
	span, _ := utils.LogSpanFromContext(ctx.Request.Context())
	span.AddEvent(`teacher.info.ctl.start`)

	// Get ID from URL parameter
	idParam := ctx.Param("id")
	id, err := uuid.Parse(idParam)
	if err != nil {
		base.BadRequest(ctx, i18n.BadRequest, nil)
		return
	}
	span.AddEvent(`teacher.info.ctl.request`)

	data, err := c.svc.InfoService(ctx.Request.Context(), &InfoServiceRequest{
		ID: id,
	})
	if err != nil {
		base.HandleError(ctx, err)
		return
	}
	span.AddEvent(`teacher.info.ctl.callsvc`)

	var resp InfoControllerResponse
	if err := utils.CopyNTimeToUnix(&resp, data); err != nil {
		base.InternalServerError(ctx, err.Error(), nil)
		return
	}
	span.AddEvent(`teacher.info.ctl.end`)

	base.Success(ctx, resp)
}

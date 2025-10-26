package teacher

import (
	"github.com/easy-attend-serviceV3/app/utils"
	"github.com/easy-attend-serviceV3/app/utils/base"
	"github.com/easy-attend-serviceV3/config/i18n"
	"github.com/gin-gonic/gin"
)

type ListControllerRequest struct {
	base.RequestPaginate
}

type ListControllerResponse struct {
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

func (c *Controller) ListController(ctx *gin.Context) {
	span, _ := utils.LogSpanFromContext(ctx.Request.Context())
	span.AddEvent(`teacher.list.ctl.start`)

	var request ListControllerRequest
	if err := ctx.ShouldBind(&request); err != nil {
		base.BadRequest(ctx, i18n.BadRequest, nil)
		return
	}
	span.AddEvent(`teacher.list.ctl.request`)

	data, page, err := c.svc.ListService(ctx.Request.Context(), &ListServiceRequest{
		RequestPaginate: request.RequestPaginate,
	})
	if err != nil {
		base.HandleError(ctx, err)
		return
	}
	span.AddEvent(`teacher.list.ctl.callsvc`)

	var resp []*ListControllerResponse
	if err := utils.CopyNTimeToUnix(&resp, data); err != nil {
		base.InternalServerError(ctx, err.Error(), nil)
		return
	}
	span.AddEvent(`teacher.list.ctl.end`)

	base.Paginate(ctx, resp, page)
}

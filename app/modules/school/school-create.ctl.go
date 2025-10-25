package school

import (
	"github.com/easy-attend-serviceV3/app/utils"
	"github.com/easy-attend-serviceV3/app/utils/base"
	"github.com/easy-attend-serviceV3/config/i18n"
	"github.com/gin-gonic/gin"
)

type CreateControllerRequest struct {
	Name    string `json:"name" binding:"required"`
	Address string `json:"address"`
	Phone   string `json:"phone"`
}

func (c *Controller) CreateController(ctx *gin.Context) {
	span, _ := utils.LogSpanFromContext(ctx.Request.Context())
	span.AddEvent(`school.create.ctl.start`)

	var request CreateControllerRequest
	if err := ctx.ShouldBindJSON(&request); err != nil {
		base.BadRequest(ctx, i18n.BadRequest, nil)
		return
	}
	span.AddEvent(`school.create.ctl.request`)

	// Validate required fields
	if request.Name == "" {
		base.BadRequest(ctx, i18n.BadRequest, nil)
		return
	}

	if err := c.svc.CreateService(ctx.Request.Context(), &CreateServiceRequest{
		Name:    request.Name,
		Address: request.Address,
		Phone:   request.Phone,
	}); err != nil {
		base.HandleError(ctx, err)
		return
	}
	span.AddEvent(`school.create.ctl.end`)
	base.Success(ctx, nil)
}

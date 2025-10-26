package teacher

import (
	"github.com/easy-attend-serviceV3/app/utils"
	"github.com/easy-attend-serviceV3/app/utils/base"
	"github.com/easy-attend-serviceV3/config/i18n"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

type CreateControllerRequest struct {
	SchoolID    string `json:"school_id" binding:"required"`
	ClassroomID string `json:"classroom_id"`
	PrefixID    string `json:"prefix_id" binding:"required"`
	GenderID    string `json:"gender_id" binding:"required"`
	FirstName   string `json:"first_name" binding:"required"`
	LastName    string `json:"last_name" binding:"required"`
	Email       string `json:"email" binding:"required,email"`
	Password    string `json:"password" binding:"required,min=6"`
	Phone       string `json:"phone"`
}

type CreateControllerResponse struct {
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

func (c *Controller) CreateController(ctx *gin.Context) {
	span, _ := utils.LogSpanFromContext(ctx.Request.Context())
	span.AddEvent(`teacher.create.ctl.start`)

	var request CreateControllerRequest
	if err := ctx.ShouldBindJSON(&request); err != nil {
		base.BadRequest(ctx, i18n.BadRequest, nil)
		return
	}
	span.AddEvent(`teacher.create.ctl.request`)

	// Parse and validate UUIDs
	schoolID, err := uuid.Parse(request.SchoolID)
	if err != nil {
		base.BadRequest(ctx, i18n.BadRequest, nil)
		return
	}

	var classroomID uuid.UUID
	if request.ClassroomID != "" {
		classroomID, err = uuid.Parse(request.ClassroomID)
		if err != nil {
			base.BadRequest(ctx, i18n.BadRequest, nil)
			return
		}
	}

	prefixID, err := uuid.Parse(request.PrefixID)
	if err != nil {
		base.BadRequest(ctx, i18n.BadRequest, nil)
		return
	}

	genderID, err := uuid.Parse(request.GenderID)
	if err != nil {
		base.BadRequest(ctx, i18n.BadRequest, nil)
		return
	}

	teacher, err := c.svc.CreateService(ctx.Request.Context(), &CreateServiceRequest{
		SchoolID:    schoolID,
		ClassroomID: classroomID,
		PrefixID:    prefixID,
		GenderID:    genderID,
		FirstName:   request.FirstName,
		LastName:    request.LastName,
		Email:       request.Email,
		Password:    request.Password,
		Phone:       request.Phone,
	})
	if err != nil {
		base.HandleError(ctx, err)
		return
	}
	span.AddEvent(`teacher.create.ctl.callsvc`)

	var resp CreateControllerResponse
	if err := utils.CopyNTimeToUnix(&resp, teacher); err != nil {
		base.InternalServerError(ctx, err.Error(), nil)
		return
	}
	span.AddEvent(`teacher.create.ctl.end`)

	base.Success(ctx, resp)
}

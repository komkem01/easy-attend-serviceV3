package student

import (
	"github.com/easy-attend-serviceV3/app/utils"
	"github.com/easy-attend-serviceV3/app/utils/base"
	"github.com/easy-attend-serviceV3/config/i18n"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

type UpdateControllerRequest struct {
	SchoolID    string `json:"school_id" binding:"required"`
	ClassroomID string `json:"classroom_id" binding:"required"`
	PrefixID    string `json:"prefix_id" binding:"required"`
	GenderID    string `json:"gender_id" binding:"required"`
	StudentCode string `json:"student_code" binding:"required"`
	FirstName   string `json:"first_name" binding:"required"`
	LastName    string `json:"last_name" binding:"required"`
	Phone       string `json:"phone"`
}

func (c *Controller) UpdateController(ctx *gin.Context) {
	span, _ := utils.LogSpanFromContext(ctx.Request.Context())
	span.AddEvent(`student.update.ctl.start`)

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
	span.AddEvent(`student.update.ctl.request`)

	// Parse UUIDs
	schoolID, err := uuid.Parse(request.SchoolID)
	if err != nil {
		base.BadRequest(ctx, i18n.BadRequest, nil)
		return
	}
	classroomID, err := uuid.Parse(request.ClassroomID)
	if err != nil {
		base.BadRequest(ctx, i18n.BadRequest, nil)
		return
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

	if err := c.svc.UpdateService(ctx.Request.Context(), &UpdateServiceRequest{
		ID:          id,
		SchoolID:    schoolID,
		ClassroomID: classroomID,
		PrefixID:    prefixID,
		GenderID:    genderID,
		StudentCode: request.StudentCode,
		FirstName:   request.FirstName,
		LastName:    request.LastName,
		Phone:       request.Phone,
	}); err != nil {
		base.HandleError(ctx, err)
		return
	}

	span.AddEvent(`student.update.ctl.end`)
	base.Success(ctx, nil)
}

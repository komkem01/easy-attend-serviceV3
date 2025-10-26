package student

import (
	"github.com/easy-attend-serviceV3/app/utils"
	"github.com/easy-attend-serviceV3/app/utils/base"
	"github.com/easy-attend-serviceV3/config/i18n"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

type CreateControllerRequest struct {
	SchoolID    string `json:"school_id" binding:"required"`
	ClassroomID string `json:"classroom_id"` // ไม่บังคับกรอก
	PrefixID    string `json:"prefix_id" binding:"required"`
	GenderID    string `json:"gender_id" binding:"required"`
	StudentCode string `json:"student_code" binding:"required"`
	FirstName   string `json:"first_name" binding:"required"`
	LastName    string `json:"last_name" binding:"required"`
	Phone       string `json:"phone"`
}

type CreateControllerResponse struct {
	ID          uuid.UUID  `json:"id"`
	SchoolID    uuid.UUID  `json:"school_id"`
	ClassroomID *uuid.UUID `json:"classroom_id,omitempty"` // อาจเป็น null
	PrefixID    uuid.UUID  `json:"prefix_id"`
	GenderID    uuid.UUID  `json:"gender_id"`
	StudentCode string     `json:"student_code"`
	FirstName   string     `json:"first_name"`
	LastName    string     `json:"last_name"`
	Phone       string     `json:"phone"`
}

func (c *Controller) CreateController(ctx *gin.Context) {
	span, _ := utils.LogSpanFromContext(ctx.Request.Context())
	span.AddEvent(`student.create.ctl.start`)

	var request CreateControllerRequest
	if err := ctx.ShouldBindJSON(&request); err != nil {
		base.BadRequest(ctx, i18n.BadRequest, nil)
		return
	}
	span.AddEvent(`student.create.ctl.request`)

	// Validate and parse UUIDs
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

	// Validate required fields
	if request.FirstName == "" || request.LastName == "" || request.StudentCode == "" {
		base.BadRequest(ctx, i18n.BadRequest, nil)
		return
	}

	student, err := c.svc.CreateService(ctx.Request.Context(), &CreateServiceRequest{
		SchoolID:    schoolID,
		ClassroomID: classroomID,
		PrefixID:    prefixID,
		GenderID:    genderID,
		StudentCode: request.StudentCode,
		FirstName:   request.FirstName,
		LastName:    request.LastName,
		Phone:       request.Phone,
	})
	if err != nil {
		base.HandleError(ctx, err)
		return
	}
	span.AddEvent(`student.create.ctl.end`)

	var classroomIDPtr *uuid.UUID
	if student.ClassroomID != uuid.Nil {
		classroomIDPtr = &student.ClassroomID
	}

	response := &CreateControllerResponse{
		ID:          student.ID,
		SchoolID:    student.SchoolID,
		ClassroomID: classroomIDPtr,
		PrefixID:    student.PrefixID,
		GenderID:    student.GenderID,
		StudentCode: student.StudentCode,
		FirstName:   student.FirstName,
		LastName:    student.LastName,
		Phone:       student.Phone,
	}
	base.Success(ctx, response)
}

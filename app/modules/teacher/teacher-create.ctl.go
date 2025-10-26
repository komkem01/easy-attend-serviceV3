package teacher

import (
	"github.com/easy-attend-serviceV3/app/utils"
	"github.com/easy-attend-serviceV3/app/utils/base"
	"github.com/easy-attend-serviceV3/config/i18n"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"go.opentelemetry.io/otel/attribute"
)

type CreateControllerRequest struct {
	SchoolName  string `json:"school_name" binding:"required"` // เปลี่ยนจาก school_id เป็น school_name
	ClassroomID string `json:"classroom_id"`                   // ไม่บังคับกรอก
	PrefixID    string `json:"prefix_id" binding:"required"`
	GenderID    string `json:"gender_id" binding:"required"`
	FirstName   string `json:"first_name" binding:"required"`
	LastName    string `json:"last_name" binding:"required"`
	Email       string `json:"email" binding:"required,email"`
	Password    string `json:"password" binding:"required,min=6"`
	Phone       string `json:"phone"`
}

type CreateControllerResponse struct {
	ID           string  `json:"id"`
	SchoolID     string  `json:"school_id"`
	SchoolName   string  `json:"school_name"`            // เพิ่มชื่อโรงเรียน
	ClassroomID  *string `json:"classroom_id,omitempty"` // อาจเป็น null
	PrefixID     string  `json:"prefix_id"`
	GenderID     string  `json:"gender_id"`
	FirstName    string  `json:"first_name"`
	LastName     string  `json:"last_name"`
	Email        string  `json:"email"`
	Phone        string  `json:"phone"`
	AccessToken  string  `json:"access_token,omitempty"`
	RefreshToken string  `json:"refresh_token,omitempty"`
	TokenType    string  `json:"token_type,omitempty"`
	ExpiresAt    int64   `json:"expires_at,omitempty"`
	CreatedAt    int64   `json:"created_at"`
	UpdatedAt    int64   `json:"updated_at"`
}

func (c *Controller) CreateController(ctx *gin.Context) {
	span, _ := utils.LogSpanFromContext(ctx.Request.Context())
	span.AddEvent(`teacher.create.ctl.start`)

	// Debug: Log that registration endpoint was called
	span.AddEvent("teacher.registration.endpoint.called")

	var request CreateControllerRequest
	if err := ctx.ShouldBindJSON(&request); err != nil {
		span.AddEvent("teacher.registration.bind.failed")
		base.BadRequest(ctx, i18n.BadRequest, nil)
		return
	}
	span.AddEvent(`teacher.create.ctl.request`)

	// Debug: Log the parsed request (without password)
	span.SetAttributes(
		attribute.String("email", request.Email),
		attribute.String("school_name", request.SchoolName),
		attribute.String("first_name", request.FirstName),
	)

	// Parse and validate UUIDs
	var classroomID *uuid.UUID
	if request.ClassroomID != "" {
		parsed, err := uuid.Parse(request.ClassroomID)
		if err != nil {
			base.BadRequest(ctx, i18n.BadRequest, nil)
			return
		}
		classroomID = &parsed
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
		SchoolName:  request.SchoolName, // ส่งชื่อโรงเรียนแทน ID
		ClassroomID: classroomID,        // ส่งเป็น pointer
		PrefixID:    prefixID,
		GenderID:    genderID,
		FirstName:   request.FirstName,
		LastName:    request.LastName,
		Email:       request.Email,
		Password:    request.Password,
		Phone:       request.Phone,
	})
	if err != nil {
		base.HandleCustomError(ctx, err)
		return
	}
	span.AddEvent(`teacher.create.ctl.callsvc`)

	// Handle optional classroom_id for response
	var responseClassroomID *string
	if teacher.ClassroomID != nil {
		classroomIDStr := teacher.ClassroomID.String()
		responseClassroomID = &classroomIDStr
	}

	resp := CreateControllerResponse{
		ID:           teacher.ID.String(),
		SchoolID:     teacher.SchoolID.String(),
		SchoolName:   teacher.SchoolName,
		ClassroomID:  responseClassroomID,
		PrefixID:     teacher.PrefixID.String(),
		GenderID:     teacher.GenderID.String(),
		FirstName:    teacher.FirstName,
		LastName:     teacher.LastName,
		Email:        teacher.Email,
		Phone:        teacher.Phone,
		AccessToken:  teacher.AccessToken,
		RefreshToken: teacher.RefreshToken,
		TokenType:    teacher.TokenType,
		ExpiresAt:    teacher.ExpiresAt.Unix(),
	}
	span.AddEvent(`teacher.create.ctl.end`)

	base.Success(ctx, resp)
}

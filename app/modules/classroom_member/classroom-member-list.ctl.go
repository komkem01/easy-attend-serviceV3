package classroommember

import (
	"net/http"

	"github.com/easy-attend-serviceV3/app/utils"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

func (c *Controller) ListController(ctx *gin.Context) {
	span, log := utils.LogSpanFromContext(ctx)
	span.AddEvent(`classroom_member.ctl.list.start`)

	// Parse query parameters
	var req ListServiceRequest

	if classroomIDStr := ctx.Query("classroom_id"); classroomIDStr != "" {
		classroomID, err := uuid.Parse(classroomIDStr)
		if err != nil {
			log.Error(err)
			ctx.JSON(http.StatusBadRequest, gin.H{
				"code":    "400",
				"message": "Invalid classroom_id format",
				"data":    nil,
			})
			return
		}
		req.ClassroomID = &classroomID
	}

	if studentIDStr := ctx.Query("student_id"); studentIDStr != "" {
		studentID, err := uuid.Parse(studentIDStr)
		if err != nil {
			log.Error(err)
			ctx.JSON(http.StatusBadRequest, gin.H{
				"code":    "400",
				"message": "Invalid student_id format",
				"data":    nil,
			})
			return
		}
		req.StudentID = &studentID
	}

	result, err := c.svc.ListService(ctx, &req)
	if err != nil {
		log.Error(err)
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"code":    "500",
			"message": err.Error(),
			"data":    nil,
		})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"code":    "200",
		"message": "Success",
		"data":    result,
	})

	span.AddEvent(`classroom_member.ctl.list.end`)
}

package attendance

import (
	"net/http"

	"github.com/easy-attend-serviceV3/app/utils"
	"github.com/gin-gonic/gin"
)

func (c *Controller) CreateController(ctx *gin.Context) {
	span, log := utils.LogSpanFromContext(ctx)
	span.AddEvent(`attendance.ctl.create.start`)

	var req CreateServiceRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		log.Error(err)
		ctx.JSON(http.StatusBadRequest, gin.H{
			"code":    "400",
			"message": "Invalid request body",
			"data":    nil,
		})
		return
	}

	result, err := c.svc.CreateService(ctx, &req)
	if err != nil {
		log.Error(err)
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"code":    "500",
			"message": err.Error(),
			"data":    nil,
		})
		return
	}

	ctx.JSON(http.StatusCreated, gin.H{
		"code":    "201",
		"message": "Attendance record created successfully",
		"data":    result,
	})

	span.AddEvent(`attendance.ctl.create.end`)
}

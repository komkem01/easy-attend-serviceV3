package attendance

import (
	"net/http"

	"github.com/easy-attend-serviceV3/app/utils"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

func (c *Controller) UpdateController(ctx *gin.Context) {
	span, log := utils.LogSpanFromContext(ctx)
	span.AddEvent(`attendance.ctl.update.start`)

	idStr := ctx.Param("id")
	id, err := uuid.Parse(idStr)
	if err != nil {
		log.Error(err)
		ctx.JSON(http.StatusBadRequest, gin.H{
			"code":    "400",
			"message": "Invalid ID format",
			"data":    nil,
		})
		return
	}

	var req UpdateServiceRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		log.Error(err)
		ctx.JSON(http.StatusBadRequest, gin.H{
			"code":    "400",
			"message": "Invalid request body",
			"data":    nil,
		})
		return
	}

	req.ID = id

	result, err := c.svc.UpdateService(ctx, &req)
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
		"message": "Attendance record updated successfully",
		"data":    result,
	})

	span.AddEvent(`attendance.ctl.update.end`)
}

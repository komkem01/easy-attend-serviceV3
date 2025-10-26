package classroommember

import (
	"net/http"

	"github.com/easy-attend-serviceV3/app/utils"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

func (c *Controller) DeleteController(ctx *gin.Context) {
	span, log := utils.LogSpanFromContext(ctx)
	span.AddEvent(`classroom_member.ctl.delete.start`)

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

	req := &DeleteServiceRequest{
		ID: id,
	}

	result, err := c.svc.DeleteService(ctx, req)
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
		"message": result.Message,
		"data":    result,
	})

	span.AddEvent(`classroom_member.ctl.delete.end`)
}

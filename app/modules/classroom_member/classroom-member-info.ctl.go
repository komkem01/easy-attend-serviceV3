package classroommember

import (
	"net/http"

	"github.com/easy-attend-serviceV3/app/utils"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

func (c *Controller) InfoController(ctx *gin.Context) {
	span, log := utils.LogSpanFromContext(ctx)
	span.AddEvent(`classroom_member.ctl.info.start`)

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

	req := &InfoServiceRequest{
		ID: id,
	}

	result, err := c.svc.InfoService(ctx, req)
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

	span.AddEvent(`classroom_member.ctl.info.end`)
}

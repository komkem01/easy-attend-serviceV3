package teacher

import (
	"net/http"

	"github.com/easy-attend-serviceV3/app/utils"
	"github.com/gin-gonic/gin"
)

func (c *Controller) Login(ctx *gin.Context) {
	span, log := utils.LogSpanFromContext(ctx)
	span.AddEvent(`teacher.ctl.login.start`)

	var req LoginServiceRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		log.Error(err)
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error":   "Invalid request body",
			"message": err.Error(),
		})
		return
	}

	result, err := c.svc.LoginService(ctx, &req)
	if err != nil {
		log.Error(err)
		status := http.StatusInternalServerError
		if err.Error() == "invalid email or password" {
			status = http.StatusUnauthorized
		}
		ctx.JSON(status, gin.H{
			"error":   "Login failed",
			"message": err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"message": "Login successful",
		"data":    result,
	})

	span.AddEvent(`teacher.ctl.login.end`)
}

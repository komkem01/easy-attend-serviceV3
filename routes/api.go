package routes

import (
	"fmt"
	"net/http"

	"github.com/easy-attend-serviceV3/app/modules"

	"github.com/gin-gonic/gin"
)

func WarpH(router *gin.RouterGroup, prefix string, handler http.Handler) {
	router.Any(fmt.Sprintf("%s/*w", prefix), gin.WrapH(http.StripPrefix(fmt.Sprintf("%s%s", router.BasePath(), prefix), handler)))
}

func api(r *gin.RouterGroup, mod *modules.Modules) {
	r.GET("/example/:id", mod.Example.Ctl.Get)
	r.GET("/example-http", mod.Example.Ctl.GetHttpReq)
	r.POST("/example", mod.Example.Ctl.Create)

	// Gender routes
	r.GET("/gender", mod.Gender.Ctl.ListController)
	r.GET("/gender/:id", mod.Gender.Ctl.InfoController)

	// Prefix routes
	r.GET("/prefix", mod.Prefix.Ctl.ListController)
	r.GET("/prefix/:id", mod.Prefix.Ctl.InfoController)

	// School routes
	r.GET("/school", mod.School.Ctl.ListController)
	r.GET("/school/:id", mod.School.Ctl.InfoController)
	r.POST("/school", mod.School.Ctl.CreateController)
	r.PATCH("/school/:id", mod.School.Ctl.UpdateController)
	r.DELETE("/school/:id", mod.School.Ctl.DeleteController)

	// Classroom routes
	r.GET("/classroom", mod.Classroom.Ctl.ListController)
	r.GET("/classroom/:id", mod.Classroom.Ctl.InfoController)
	r.POST("/classroom", mod.Classroom.Ctl.CreateController)
	r.PATCH("/classroom/:id", mod.Classroom.Ctl.UpdateController)
	r.DELETE("/classroom/:id", mod.Classroom.Ctl.DeleteController)

	// Student routes
	r.GET("/student", mod.Student.Ctl.ListController)
	r.GET("/student/:id", mod.Student.Ctl.InfoController)
	r.POST("/student", mod.Student.Ctl.CreateController)
	r.PATCH("/student/:id", mod.Student.Ctl.UpdateController)
	r.DELETE("/student/:id", mod.Student.Ctl.DeleteController)

	// Teacher routes
	r.GET("/teacher", mod.Teacher.Ctl.ListController)
	r.GET("/teacher/:id", mod.Teacher.Ctl.InfoController)
	r.POST("/teacher", mod.Teacher.Ctl.CreateController)
	r.PATCH("/teacher/:id", mod.Teacher.Ctl.UpdateController)
	r.DELETE("/teacher/:id", mod.Teacher.Ctl.DeleteController)

}

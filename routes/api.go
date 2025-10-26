package routes

import (
	"fmt"
	"net/http"

	"github.com/easy-attend-serviceV3/app/modules"
	"github.com/easy-attend-serviceV3/app/utils/auth"

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

	// Classroom Member routes
	r.GET("/classroom-member", mod.ClassroomMember.Ctl.ListController)
	r.GET("/classroom-member/:id", mod.ClassroomMember.Ctl.InfoController)
	r.POST("/classroom-member", mod.ClassroomMember.Ctl.CreateController)
	r.PATCH("/classroom-member/:id", mod.ClassroomMember.Ctl.UpdateController)
	r.DELETE("/classroom-member/:id", mod.ClassroomMember.Ctl.DeleteController)

	// Student routes
	r.GET("/student", mod.Student.Ctl.ListController)
	r.GET("/student/:id", mod.Student.Ctl.InfoController)
	r.POST("/student", mod.Student.Ctl.CreateController)
	r.PATCH("/student/:id", mod.Student.Ctl.UpdateController)
	r.DELETE("/student/:id", mod.Student.Ctl.DeleteController)

	// Teacher routes - Public
	r.POST("/teacher", mod.Teacher.Ctl.CreateController) // Registration
	r.POST("/teacher/login", mod.Teacher.Ctl.Login)      // Login

	// Teacher routes - Protected
	teacherProtected := r.Group("/teacher")
	teacherProtected.Use(auth.RequireAuth())
	{
		teacherProtected.GET("", mod.Teacher.Ctl.ListController)
		teacherProtected.GET("/:id", mod.Teacher.Ctl.InfoController)
		teacherProtected.PATCH("/:id", mod.Teacher.Ctl.UpdateController)
		teacherProtected.DELETE("/:id", mod.Teacher.Ctl.DeleteController)
	}

	// Attendance routes
	r.GET("/attendance", mod.Attendance.Ctl.ListController)
	r.GET("/attendance/:id", mod.Attendance.Ctl.InfoController)
	r.POST("/attendance", mod.Attendance.Ctl.CreateController)
	r.PATCH("/attendance/:id", mod.Attendance.Ctl.UpdateController)
	r.DELETE("/attendance/:id", mod.Attendance.Ctl.DeleteController)

}

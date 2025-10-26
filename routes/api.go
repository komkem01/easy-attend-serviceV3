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
	// Public routes (no authentication required)
	r.POST("/teacher", mod.Teacher.Ctl.CreateController) // Registration
	r.POST("/teacher/login", mod.Teacher.Ctl.Login)      // Login

	// Protected routes (authentication required)
	protected := r.Group("")
	protected.Use(auth.RequireAuth())
	{
		// Example routes
		protected.GET("/example/:id", mod.Example.Ctl.Get)
		protected.GET("/example-http", mod.Example.Ctl.GetHttpReq)
		protected.POST("/example", mod.Example.Ctl.Create)

		// Gender routes
		protected.GET("/gender", mod.Gender.Ctl.ListController)
		protected.GET("/gender/:id", mod.Gender.Ctl.InfoController)

		// Prefix routes
		protected.GET("/prefix", mod.Prefix.Ctl.ListController)
		protected.GET("/prefix/:id", mod.Prefix.Ctl.InfoController)

		// School routes
		protected.GET("/school", mod.School.Ctl.ListController)
		protected.GET("/school/:id", mod.School.Ctl.InfoController)
		protected.POST("/school", mod.School.Ctl.CreateController)
		protected.PATCH("/school/:id", mod.School.Ctl.UpdateController)
		protected.DELETE("/school/:id", mod.School.Ctl.DeleteController)

		// Classroom routes
		protected.GET("/classroom", mod.Classroom.Ctl.ListController)
		protected.GET("/classroom/:id", mod.Classroom.Ctl.InfoController)
		protected.POST("/classroom", mod.Classroom.Ctl.CreateController)
		protected.PATCH("/classroom/:id", mod.Classroom.Ctl.UpdateController)
		protected.DELETE("/classroom/:id", mod.Classroom.Ctl.DeleteController)

		// Classroom Member routes
		protected.GET("/classroom-member", mod.ClassroomMember.Ctl.ListController)
		protected.GET("/classroom-member/:id", mod.ClassroomMember.Ctl.InfoController)
		protected.POST("/classroom-member", mod.ClassroomMember.Ctl.CreateController)
		protected.PATCH("/classroom-member/:id", mod.ClassroomMember.Ctl.UpdateController)
		protected.DELETE("/classroom-member/:id", mod.ClassroomMember.Ctl.DeleteController)

		// Student routes
		protected.GET("/student", mod.Student.Ctl.ListController)
		protected.GET("/student/:id", mod.Student.Ctl.InfoController)
		protected.POST("/student", mod.Student.Ctl.CreateController)
		protected.PATCH("/student/:id", mod.Student.Ctl.UpdateController)
		protected.DELETE("/student/:id", mod.Student.Ctl.DeleteController)

		// Teacher routes
		protected.GET("/teacher", mod.Teacher.Ctl.ListController)
		protected.GET("/teacher/:id", mod.Teacher.Ctl.InfoController)
		protected.PATCH("/teacher/:id", mod.Teacher.Ctl.UpdateController)
		protected.DELETE("/teacher/:id", mod.Teacher.Ctl.DeleteController)

		// Attendance routes
		protected.GET("/attendance", mod.Attendance.Ctl.ListController)
		protected.GET("/attendance/:id", mod.Attendance.Ctl.InfoController)
		protected.POST("/attendance", mod.Attendance.Ctl.CreateController)
		protected.PATCH("/attendance/:id", mod.Attendance.Ctl.UpdateController)
		protected.DELETE("/attendance/:id", mod.Attendance.Ctl.DeleteController)
	}

}

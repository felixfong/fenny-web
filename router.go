package main

import (
	"fenny-web/framework/gin"
	"fenny-web/framework/middleware"
)

func registerRouter(core *gin.Engine) {
	core.GET("/user/login", middleware.Test1(), UserLoginController)

	subjectApi := core.Group("/subject")
	{
		subjectApi.Use(middleware.Test1())
		// 动态路由
		subjectApi.GET("/:id", SubjectListController)
		subjectApi.PUT("/:id", SubjectListController)
	}
}

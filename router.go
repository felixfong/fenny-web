package main

import (
	"fenng-web/framework"
	"fenng-web/framework/middleware"
)

func registerRouter(core *framework.Core) {
	core.Use(
			middleware.Test1(),
			middleware.Test2(), FooControllerHandler)

	subjectApi := core.Group("/subject")
	{
		subjectApi.Use(middleware.Test1(), SubjectListController)
	}
}

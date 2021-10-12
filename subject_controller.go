package main

import (
	"fenny-web/framework/gin"
	"fenny-web/provider/demo"
)

func SubjectListController(ctx *gin.Context) {
	demoService := ctx.MustMake(demo.Key).(demo.Service)

	foo := demoService.GetFoo()

	ctx.ISetOkStatus().IJson(foo)
}

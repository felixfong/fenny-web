package main

import (
	"fenny-web/framework/gin"
	"time"
)

func UserLoginController(ctx *gin.Context) {
	foo, _ := ctx.DefaultQueryString("foo", "def")
	time.Sleep(10*time.Second)
	ctx.ISetOkStatus().IJson("OK, UserLoginController: " + foo)
}

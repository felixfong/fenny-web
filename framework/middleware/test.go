package middleware

import (
	"fenny-web/framework/gin"
	"fmt"
)

func Test1() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		fmt.Println("middleware pre test1")
		ctx.Next()
		fmt.Println("middleware post test1")
	}
}

func Test2() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		fmt.Println("middleware pre test2")
		ctx.Next()
		fmt.Println("middleware post test2")
	}
}

package middleware

import (
	"context"
	"fenny-web/framework/gin"
	"fmt"
	"log"
	"time"
)

func Timeout(d time.Duration) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		finish := make(chan struct{}, 1)
		panicChan := make(chan interface{}, 1)

		durationCtx, cancel := context.WithTimeout(context.Background(), time.Second)
		defer cancel()

		go func(){
			defer func(){
				if p := recover(); p != nil {
					panicChan <- p
				}
			}()

			ctx.Next()

			finish <- struct{}{}
		}()

		select {
		case p := <- panicChan:
			ctx.JSON(500, "exception")
			log.Println(p)
		case <- finish:
			fmt.Println("finish")
		case <- durationCtx.Done():
			ctx.JSON(500, "time out")
		}

	}
}

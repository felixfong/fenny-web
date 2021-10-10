package middleware

import (
	"context"
	"fenng-web/framework"
	"fmt"
	"log"
	"time"
)

func Timeout(d time.Duration) framework.ControllerHandler {
	return func(ctx *framework.Context) error {
		finish := make(chan struct{}, 1)
		panicChan := make(chan interface{}, 1)

		durationCtx, cancel := context.WithTimeout(ctx.BaseContext(), time.Second)
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
			ctx.Json(500, "exception")
			log.Println(p)
		case <- finish:
			fmt.Println("finish")
		case <- durationCtx.Done():
			ctx.Json(500, "time out")
		}

		return nil
	}
}

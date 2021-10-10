package framework

import (
	"context"
	"fmt"
	"log"
	"time"
)

func Timeout(d time.Duration) ControllerHandler {
	return func(ctx *Context) error {
		finish := make(chan struct{}, 1)
		panicChan := make(chan interface{}, 1)

		durationCtx, cancel := context.WithTimeout(ctx.BaseContext(), d)
		defer cancel()

		ctx.request.WithContext(durationCtx)

		go func(){
			defer func(){
				if p:= recover(); p != nil {
					panicChan <- p
				}
			}()

			fun(ctx)

			finish <- struct{}{}
		}()

		select {
		case p := <-panicChan:
			log.Println(p)
			ctx.response.WriteHeader(500)
		case <-finish:
			fmt.Println("finish")
		case <-durationCtx.Done():
			fmt.Println("time out")
			ctx.response.Write([]byte("time out"))
		}
		return nil
	}
}

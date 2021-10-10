package main

import (
	"context"
	"fenng-web/framework"
	"fmt"
	"time"
)

func FooControllerHandler(ctx *framework.Context) error {
	durationCtx, cancel := context.WithTimeout(ctx.BaseContext(), time.Duration(10*time.Second))
	defer cancel()

	finish := make(chan struct{}, 1)
	paincChan := make(chan interface{}, 1)


	go func(){
		defer func(){
			if p:= recover(); p != nil {
				paincChan <- p
			}
		}()

		time.Sleep(1*time.Second)
		ctx.Json(200, "ok")

		finish <- struct{}{}
	}()

	select {
		case <-paincChan:
			ctx.Json(500, "panic")

		case <-finish:
			fmt.Println("finish")

		case <-durationCtx.Done():
			ctx.Json(500, "time out")
	}

	return nil
}

//func Foo1(request *http.Request, response http.ResponseWriter) {
//	obj := map[string]interface{}{
//		"data": nil,
//	}
//
//	response.Header().Set("Content-Type", "application/json")
//
//	foo := request.PostFormValue("foo")
//
//	if foo == "" {
//		foo = "10"
//	}
//	fooInt, err := strconv.Atoi(foo)
//	if err != nil {
//		response.WriteHeader(500)
//		return
//	}
//	obj["data"] = fooInt
//	byt, err := json.Marshal(obj)
//	if err != nil {
//		response.WriteHeader(500)
//		return
//	}
//	response.WriteHeader(200)
//	response.Write(byt)
//	return
//}
//
//func Foo2(ctx *framework.Context) error {
//	obj := map[string]interface{}{
//		"data": nil,
//	}
//
//	fooInt := ctx.FormInt("foo", 10)
//
//	obj["data"] = fooInt
//
//	return ctx.Json(http.StatusOK, obj)
//}

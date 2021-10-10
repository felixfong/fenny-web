package framework

import (
	"context"
	"encoding/json"
	"net/http"
)

type Context struct {
	request *http.Request
	response http.ResponseWriter

	// 当前请求handler链条
	handlers []ControllerHandler
	index int // 当前请求调用到调用链的哪个节点

}

func NewContext(request *http.Request, response http.ResponseWriter) *Context{
	return &Context{
		request:  request,
		response: response,
		index: -1,
	}
}

func (ctx *Context) FormInt(key string, val int) int{
	return 0
}

func(ctx *Context) Json(code int, obj interface{}) error {
	ctx.response.Header().Set("Content-type", "application/json")
	ctx.response.WriteHeader(code)

	byt, err := json.Marshal(obj)
	if err != nil {
		ctx.response.WriteHeader(500)
		return err
	}
	ctx.response.Write(byt)
	return nil
}

func (ctx *Context) BaseContext() context.Context {
	return ctx.request.Context()
}

func (ctx *Context) Next() error {
	ctx.index++
	if ctx.index < len(ctx.handlers) {
		if err := ctx.handlers[ctx.index](ctx); err != nil {
			return err
		}
	}
	return nil
}

func (ctx *Context) SetHandlers(handlers []ControllerHandler) {
	ctx.handlers = handlers
}
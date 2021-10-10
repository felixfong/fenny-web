package framework

import (
	"bytes"
	"encoding/json"
	"errors"
	"github.com/spf13/cast"
	"io/ioutil"
)

type IRequest interface {
	// 请求地址URL中带的参数
	// 形如：foo.com?a=1&b=bar&c[]=bar
	QueryInt(key string, def int) (int, bool)
	QueryInt64(key string, def int64) (int64, bool)
	QueryFloat64(key string, def float64) (float64, bool)
	QueryFloat32(key string, def float32) (float32, bool)
	QueryString(key string, def string) (string, bool)
	QueryBool(key string, def bool) (bool, bool)
	QueryStringSlice(key string, def []string) ([]string, bool)
	Query(key string) interface{}


	// 路由匹配中带的参数
	// 形如：/book/:id
	ParamInt(key string, def int) (int, bool)
	ParamSting( key string, def string) (string, bool)
	Param(key string) interface{}

	// form表单中带的参数
	FormInt(key string, def int) (int, bool)
	Form(key string) interface{}

	// json body
	BindJson(obj interface{}) error

	// 基础信息
	Uri() string
	Method() string
	Host() string
	ClientIp() string

	// header
	Headers() map[string]string
	Header(key string) (string, bool)
}

// 获取请求地址url中带的所有参数
func (ctx *Context) QueryAll() map[string][]string {
	if ctx.request != nil {
		return map[string][]string(ctx.request.URL.Query())
	}
	return map[string][]string{}
}

func (ctx *Context) QueryInt(key string, def int) (int, bool) {
	params := ctx.QueryAll()
	if vals, ok := params[key]; ok {
		if len(vals) > 0 {
			return cast.ToInt(vals[0]), true
		}
	}
	return def, false
}

func (ctx *Context) Query(key string) interface{} {
	params := ctx.QueryAll()
	if vals, ok := params[key]; ok {
		return vals[0]
	}
	return nil
}

// 获取路由参数
func (ctx *Context) BindJson(obj interface{}) error {
	if ctx.request != nil {
		body, err := ioutil.ReadAll(ctx.request.Body)
		if err != nil {
			return err
		}
		ctx.request.Body = ioutil.NopCloser(bytes.NewBuffer(body))
		err = json.Unmarshal(body, obj)
		if err != nil {
			return err
		}
	} else {
		return errors.New("ctx.request empty")
	}
	return nil
}

// 基础信息
func (ctx *Context) Uri() string {
	return ctx.request.RequestURI
}

func (ctx *Context) Method() string {
	return ctx.request.Method
}

func (ctx *Context) Host() string {
	return ctx.request.URL.Host
}

func (ctx *Context) ClientIp() string {
	r := ctx.request
	ipAddress := r.Header.Get("X-Real-Ip")
	if ipAddress == "" {
		ipAddress = r.Header.Get("X-Forwarded-For")
	}
	if ipAddress == "" {
		ipAddress = r.RemoteAddr
	}
	return ipAddress
}

// header
func (ctx *Context) Headers() map[string][]string {
	return map[string][]string(ctx.request.Header)
}

func (ctx *Context) Header(key string) (string, bool) {
	vals := ctx.request.Header.Get(key)
	if len(vals) <= 0 {
		return "", false
	}
	return vals, true
}



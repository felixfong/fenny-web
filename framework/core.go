package framework

import (
	"log"
	"net/http"
	"strings"
)

// 框架核心结构
type Core struct {
	router map[string]*Tree
	middlewares []ControllerHandler
}

// 实例化框架核心结构
func NewCore() *Core {
	router := map[string]*Tree{}
	router["GET"] = NewTree()
	router["POST"] = NewTree()
	router["PUT"] = NewTree()
	router["DELETE"] = NewTree()
	return &Core{
		router: router,
	}
}

func (c *Core) Get(url string, handlers ...ControllerHandler) {
	c.router["GET"].AddRouter(url, handlers)
}

func (c *Core) Post(url string, handlers ...ControllerHandler) {
	c.router["Post"].AddRouter(url, handlers)
}

func (c *Core) Put(url string, handlers ...ControllerHandler) {
	c.router["Put"].AddRouter(url, handlers)
}

func (c *Core) Delete(url string, handlers ...ControllerHandler) {
	c.router["Delete"].AddRouter(url, handlers)
}

func (c *Core) FindRouteByRequest(request *http.Request) []ControllerHandler {
	uri := request.URL.Path
	method := request.Method
	upperMethod := strings.ToUpper(method)

	if methodHandlers, ok := c.router[upperMethod]; ok {
		return methodHandlers.FindHandler(uri)
	}

	return nil
}

// 实现Handler接口
func (c *Core) ServeHTTP(response http.ResponseWriter, request *http.Request){
	log.Println("core.ServeHTTP")
	ctx := NewContext(request, response)

	router := c.FindRouteByRequest(request)
	if router == nil {
		ctx.Json(404, "not found route")
	}

	log.Println("core.router")

	ctx.SetHandlers(router)

	ctx.Next()

}

func (c *Core) Group(prefix string) *Group {
	return NewGroup(c, prefix)
}

func (c *Core) Use(middlewares ...ControllerHandler) {
	c.middlewares = middlewares
}

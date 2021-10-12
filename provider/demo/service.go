package demo

import (
	"fenny-web/framework"
	"fmt"
)

type DemoService struct {
	Service
	// 参数
	c framework.Container
}

func NewDemoService(params ...interface{}) (interface{}, error) {
	c := params[0].(framework.Container)

	fmt.Println("new demo service")

	return &DemoService{
		c:       c,
	}, nil
}

func (s *DemoService) GetFoo() Foo {
	return Foo{Name: "i am foo"}
}

package main

import (
	"fenng-web/framework"
	"fenng-web/framework/middleware"
	"net/http"
)

func main(){
	core := framework.NewCore()
	core.Use(middleware.Recovery())
	registerRouter(core)
	server := &http.Server{
		Addr:              ":8080",
		Handler:           core,

	}
	server.ListenAndServe()
}

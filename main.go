package main

import (
	"context"
	"fenny-web/framework/gin"
	"fenny-web/framework/middleware"
	"fenny-web/provider/demo"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"
)

func main(){
	core := gin.New()
	core.Bind(&demo.DemoServiceProvider{})
	core.Use(middleware.Recovery())
	registerRouter(core)
	server := &http.Server{
		Addr:              ":8080",
		Handler:           core,

	}
	go func() {
		server.ListenAndServe()
	}()

	quit := make(chan os.Signal)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM, syscall.SIGQUIT)
	<- quit

	timeoutCtx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	if err := server.Shutdown(timeoutCtx); err != nil {
		log.Fatal("server shutdown:", err)
	}

}

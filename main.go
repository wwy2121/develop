package main

import (
	"context"
	"flag"
	"fmt"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
	"log"
	"mc3_monitor_service/controller"
	"mc3_monitor_service/core"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"
)

//入口
func main() {
	env := flag.String("env", "dev", "set env /dev/qa/product")
	cfgFile := fmt.Sprintf("config/%s/config.toml", *env)
	flag.Parse()
	core.LoadConfigFile(cfgFile)
	core.InitConfile()
	Run()
}

//启动进程
func Run() {
	//路由
	router := gin.Default()
	router.Use(cors.Default())
	controller.InitRoute(router)
	srv := &http.Server{
		Addr:    viper.GetString("server.port"),
		Handler: router,
	}
	go func() {
		if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatalf("listen: %s\n", err)
		}
	}()
	quit := make(chan os.Signal)
	signal.Notify(quit, syscall.SIGTERM)
	signal.Notify(quit, syscall.SIGINT)

	<-quit
	log.Println("Shutdown Server ...")
	//设置5秒超时
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	if err := srv.Shutdown(ctx); err != nil {
		log.Fatal("Server Shutdown:", err)
	}
	log.Println("Server exiting")
}

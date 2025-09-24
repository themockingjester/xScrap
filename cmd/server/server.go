package main

import (
	"fmt"
	"os"
	"os/signal"
	"sync"
	"syscall"

	"github.com/gin-gonic/gin"
	config "github.com/xscrap/configs"
	"github.com/xscrap/internal/handlers"
	"github.com/xscrap/internal/utils/server_setup"
)

var AppConfig config.AppConfig
var AppDIContainer config.AppDIContainer

func main() {
	appWaitGroup := sync.WaitGroup{}
	AppConfig = config.AppConfig{
		WaitGroup: &appWaitGroup,
	}

	AppDIContainer = config.AppDIContainer{}
	AppDIContainer.ShuttingDown = false
	config.LoadConfig(&AppConfig)
	server_setup.SetupServer(&AppDIContainer, &AppConfig)

	go listenForShutDown()
	startRESTServer(&AppDIContainer, &AppConfig)
}

func startRESTServer(appDIContainer *config.AppDIContainer, appConfig *config.AppConfig) {
	server := gin.Default()
	handlers.RegisterRoutes(server, &AppConfig, &AppDIContainer)
	var addr string = fmt.Sprintf(":%d", AppConfig.Server.Port)
	server.Run(addr)

}

func listenForShutDown() {
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit
	AppDIContainer.ShuttingDown = true
	server_setup.ShutDownServer(&AppDIContainer, &AppConfig)
	fmt.Println("Shutdown complete.")

	os.Exit(0)
}

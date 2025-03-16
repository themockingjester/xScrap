package main

import (
	"fmt"
	"os"
	"os/signal"
	"sync"
	"syscall"

	"github.com/gin-gonic/gin"
	"github.com/xscrap/chromeDp"
	"github.com/xscrap/routes"
	"github.com/xscrap/structs"
)

var AppConfig structs.AppConfig

func main() {

	appWaitGroup := sync.WaitGroup{}
	chromeDp.InitChrome()
	AppConfig = structs.AppConfig{
		WaitGroup: &appWaitGroup,
	}

	server := gin.Default()
	routes.RegisterRoutes(server, &AppConfig)
	go listenForShutDown()
	server.Run(":7004")
}
func listenForShutDown() {
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit
	shutdown()
	os.Exit(0)
}

func shutdown() {
	fmt.Println("Would run cleanup tasks...")

	AppConfig.WaitGroup.Wait()
	chromeDp.CloseChrome()

	fmt.Println("Shutdown complete.")

}

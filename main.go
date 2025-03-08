package main

import (
	"github.com/gin-gonic/gin"
	"github.com/xscrap/routes"
)

func main() {
	server := gin.Default()
	routes.RegisterRoutes(server)
	server.Run(":7004")
}

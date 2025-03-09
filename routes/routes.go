package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/xscrap/structs"
)

func RegisterRoutes(server *gin.Engine, appConfig *structs.AppConfig) {

	eventRouterGroup := server.Group("/web-scrapper")
	RegisterWebScrapRoutes(eventRouterGroup, appConfig)
}

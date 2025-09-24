package handlers

import (
	"github.com/gin-gonic/gin"
	config "github.com/xscrap/configs"
	"github.com/xscrap/internal/middlewares"
)

func RegisterRoutes(server *gin.Engine, appConfig *config.AppConfig, appDIContainer *config.AppDIContainer) {

	scrapperRoutes := server.Group("/web-scrapper")
	scrapperRoutes.Use(middlewares.ApplyWaitGroup(appConfig, appDIContainer))

	ScrapperRoutes(scrapperRoutes, appConfig, appDIContainer)
}

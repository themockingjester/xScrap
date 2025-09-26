package handlers

import (
	config "github.com/xscrap/configs"
	controller "github.com/xscrap/internal/controllers"
	"github.com/xscrap/internal/middlewares"

	"github.com/gin-gonic/gin"
)

func ScrapperRoutes(router *gin.RouterGroup, appConfig *config.AppConfig, appDIContainer *config.AppDIContainer) {
	router.GET("/scrap", middlewares.LimitMiddleware(appConfig, appDIContainer), controller.ScrapData(appConfig, appDIContainer))

}

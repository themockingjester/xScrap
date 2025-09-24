package handlers

import (
	config "github.com/xscrap/configs"
	controller "github.com/xscrap/internal/controllers"

	"github.com/gin-gonic/gin"
)

func ScrapperRoutes(router *gin.RouterGroup, appConfig *config.AppConfig, appDIContainer *config.AppDIContainer) {
	router.GET("/scrap", controller.ScrapData(appConfig, appDIContainer))

}

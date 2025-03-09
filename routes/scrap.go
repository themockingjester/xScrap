package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/xscrap/controllers"
	"github.com/xscrap/structs"
)

func RegisterWebScrapRoutes(router *gin.RouterGroup, appConfig *structs.AppConfig) {
	router.POST("/item", controllers.ScrapData(appConfig))

}

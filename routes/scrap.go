package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/xscrap/controllers"
)

func RegisterWebScrapRoutes(router *gin.RouterGroup) {
	router.POST("/", controllers.ScrapData)

}

package routes

import (
	"github.com/gin-gonic/gin"
)

func RegisterRoutes(server *gin.Engine) {

	eventRouterGroup := server.Group("/scrap")
	RegisterWebScrapRoutes(eventRouterGroup)
}

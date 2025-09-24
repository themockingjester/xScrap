package middlewares

import (
	config "github.com/xscrap/configs"

	"github.com/gin-gonic/gin"
)

func ApplyWaitGroup(appConfig *config.AppConfig, appDIContainer *config.AppDIContainer) gin.HandlerFunc {
	return func(c *gin.Context) {

		appConfig.WaitGroup.Add(1) // Increase WaitGroup counter at the beginning

		// Ensure Done is called once request is completed (whether success or failure)
		defer appConfig.WaitGroup.Done()

		c.Next() // Process request
	}
}

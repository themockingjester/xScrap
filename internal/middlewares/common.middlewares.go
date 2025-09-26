package middlewares

import (
	"github.com/gin-gonic/gin"
	config "github.com/xscrap/configs"
)

func ApplyWaitGroup(appConfig *config.AppConfig, appDIContainer *config.AppDIContainer) gin.HandlerFunc {
	return func(c *gin.Context) {

		appConfig.WaitGroup.Add(1) // Increase WaitGroup counter at the beginning

		// Ensure Done is called once request is completed (whether success or failure)
		defer appConfig.WaitGroup.Done()

		c.Next() // Process request
	}
}

func LimitMiddleware(appConfig *config.AppConfig, appDIContainer *config.AppDIContainer) gin.HandlerFunc {

	l := appDIContainer.DynamicLimiter
	return func(c *gin.Context) {
		l.Acquire()
		defer l.Release()
		c.Next()
	}
}

package controller

import (
	"github.com/gin-gonic/gin"
	config "github.com/xscrap/configs"
	"github.com/xscrap/internal/constants"
	service "github.com/xscrap/internal/services"
)

func ScrapData(appConfig *config.AppConfig, appDIContainer *config.AppDIContainer) gin.HandlerFunc {
	return func(context *gin.Context) {
		var body constants.ScrapDataRequest
		if err := context.ShouldBindQuery(&body); err != nil {
			context.JSON(400, gin.H{"error": err.Error()})
			return
		}

		result, err := service.ScrapDataBySystem(appConfig, appDIContainer, &body)
		if err != nil {
			context.JSON(500, constants.ResponseObject{
				Success: false,
				Message: err.Error(),
				Data:    make(map[string]any),
				Code:    500,
			})
			return
		}

		// Only safe to use result here
		context.JSON(result.Code, constants.ResponseObject{
			Success: result.Success,
			Message: result.Message,
			Data:    result.Data,
			Code:    result.Code,
		})
	}
}

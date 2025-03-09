package controllers

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/xscrap/chromeDp"
	"github.com/xscrap/structs"
)

func ScrapData(appConfig *structs.AppConfig) gin.HandlerFunc {
	return func(context *gin.Context) {
		appConfig.WaitGroup.Add(1)
		defer appConfig.WaitGroup.Done()

		var scrapDataObject structs.ScrapDataObject
		err := context.ShouldBindJSON(&scrapDataObject)
		if err != nil {
			fmt.Println(err.Error(), 76665)

			context.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}

		scrapperInput := structs.ScrapDataInputObject{
			Url:     scrapDataObject.Url,
			ByXPath: scrapDataObject.ByXPath,
		}
		result, err := chromeDp.ScrapDataUsingChromeDp(&scrapperInput)
		if err != nil {

			fmt.Println(err.Error(), 543)

			context.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		context.JSON(200, gin.H{
			"message": "Successfully scraped data",
			"data":    result.Data,
		})
	}

}

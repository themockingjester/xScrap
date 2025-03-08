package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/xscrap/chromeDp"
	"github.com/xscrap/structs"
)

func ScrapData(context *gin.Context) {
	var scrapDataObject structs.ScrapDataObject
	err := context.ShouldBindJSON(&scrapDataObject)
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	scrapperInput := structs.ScrapDataInputObject{
		Url:     scrapDataObject.Url,
		ByXPath: scrapDataObject.ByXPath,
	}
	result, err := chromeDp.ScrapDataUsingChromeDp(&scrapperInput)
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	context.JSON(200, gin.H{
		"message": "Successfully scraped data",
		"data":    result.Data,
	})

}

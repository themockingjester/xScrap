package service

import (
	"fmt"
	"net/http"

	config "github.com/xscrap/configs"
	"github.com/xscrap/internal/constants"
)

func ScrapDataBySystem(appConfig *config.AppConfig, appDIContainer *config.AppDIContainer, body *constants.ScrapDataRequest) (*constants.ResultObject, error) {
	output := ""
	var err error
	switch body.ScrappingMode {
	case "cssSelector":
		if body.CssSelectorForScraping == nil {
			return nil, fmt.Errorf("cssSelector is required for cssSelector mode")
		}
		output, err = appDIContainer.WebScrapper.Scrapper.ScrapDataUsingCSSSelector(body.Url, *body.CssSelectorForScraping)

	case "xPath":
		if body.XPathForScraping == nil {
			return nil, fmt.Errorf("xPath is required for xPath mode")
		}
		output, err = appDIContainer.WebScrapper.Scrapper.ScrapDataUsingXPath(body.Url, *body.XPathForScraping)

	case "id":
		if body.IdForScraping == nil {
			return nil, fmt.Errorf("id is required for id mode")
		}
		output, err = appDIContainer.WebScrapper.Scrapper.ScrapDataUsingId(body.Url, *body.IdForScraping)

	default:
		return nil, fmt.Errorf("unsupported scrapping mode: %s", body.ScrappingMode)
	}
	if err != nil {
		return nil, err
	}
	return &constants.ResultObject{
		Success: true,
		Message: "Successfully scraped data!",
		Data:    map[string]interface{}{"data": output},
		Code:    http.StatusOK,
	}, nil

}

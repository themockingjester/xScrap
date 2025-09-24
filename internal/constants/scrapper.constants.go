package constants

type ScrapDataRequest struct {
	Url                    string  `form:"url" binding:"required,url"`
	ScrappingMode          string  `form:"scrappingMode" binding:"required,oneof=xPath cssSelector id"`
	XPathForScraping       *string `form:"xPathForScraping" binding:"required_if=ScrappingMode xPath"`
	IdForScraping          *string `form:"idForScraping" binding:"required_if=ScrappingMode id"`
	CssSelectorForScraping *string `form:"cssSelectorForScraping" binding:"required_if=ScrappingMode cssSelector"`
}

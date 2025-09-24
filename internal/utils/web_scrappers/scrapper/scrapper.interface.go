package web_scrapper

type ScrapperInterface interface {
	ScrapDataUsingXPath(url string, xPath string) (string, error)
	ScrapDataUsingCSSSelector(url string, cssSelector string) (string, error)
	ScrapDataUsingId(url string, id string) (string, error)
}

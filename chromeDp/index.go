package chromeDp

import (
	"context"

	"github.com/chromedp/chromedp"
	"github.com/xscrap/structs"
)

func ScrapDataUsingChromeDp(data *structs.ScrapDataInputObject) (*structs.ScrapDataDefaultResult, error) {
	var url string = data.Url
	byClass := data.ByClass
	byXPath := data.ByXPath
	byId := data.ById

	if len(byXPath) > 0 {
		// Implement scraping using XPath

		input := structs.ScrapDataUsingXPath{
			Url:     url,
			ByXPath: byXPath,
		}
		result, err := scrapUsingXPath(&input)
		if err != nil {
			return nil, err
		}
		return result, nil
	} else if len(byClass) > 0 {

	} else if len(byId) > 0 {

	} else {

	}
	return nil, nil // Placeholder return for now, replace with actual scraping logic when implemented
}

func scrapUsingXPath(scrapInput *structs.ScrapDataUsingXPath) (*structs.ScrapDataDefaultResult, error) {

	xPath := scrapInput.ByXPath
	url := scrapInput.Url
	// Create a new Chrome context
	allocCtx, cancel := chromedp.NewExecAllocator(context.Background(), append(
		chromedp.DefaultExecAllocatorOptions[:],
		chromedp.Flag("headless", true),              // Run without UI
		chromedp.Flag("no-sandbox", true),            // Bypass sandbox restrictions
		chromedp.Flag("disable-gpu", true),           // No GPU rendering
		chromedp.Flag("disable-dev-shm-usage", true), // Fix shared memory issues
	)...)
	defer cancel()

	ctx, cancel := chromedp.NewContext(allocCtx)
	defer cancel()
	// Variable to store extracted text
	var extractedText string

	// Run chromedp tasks
	err := chromedp.Run(ctx,
		chromedp.Navigate(url),
		chromedp.Text(xPath, &extractedText), // XPath to extract data
	)

	if err != nil {
		return nil, err
	}

	output := structs.ScrapDataDefaultResult{
		Url:  url,
		Mode: "xPath",
		Data: extractedText,
	}
	return &output, nil
}

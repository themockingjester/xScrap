package chromeDp

import (
	"context"
	"fmt"
	"log"

	"github.com/chromedp/chromedp"
	"github.com/xscrap/structs"
)

var chromeCtx *context.Context
var cancelFunc context.CancelFunc // Store cancel function globally
var allocCtx context.Context
var defaultTabCtx context.Context

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

	if chromeCtx == nil {
		return nil, fmt.Errorf("Chrome instance is not initialized")
	}

	if defaultTabCtx == nil {
		createDefaultTab()

	}
	// Create a new tab (not a new browser instance)
	tabCtx, _ := chromedp.NewContext(*chromeCtx) // New tab inside existing browser
	// defer cancel()                                    // Close tab after execution
	xPath := scrapInput.ByXPath
	url := scrapInput.Url

	// Variable to store extracted text
	var extractedText string

	// Run chromedp tasks
	err := chromedp.Run(tabCtx,
		chromedp.Navigate(url),
		chromedp.Text(xPath, &extractedText, chromedp.NodeVisible), // XPath to extract data
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

// Initializes a single Chrome instance
func InitChrome() *context.Context {

	// Prevent reinitialization if Chrome is already running
	if chromeCtx != nil {
		fmt.Println("Chrome is already initialized.")
		return chromeCtx
	}
	opts := append(chromedp.DefaultExecAllocatorOptions[:],
		// chromedp.Flag("headless", true),
		// chromedp.Flag("disable-gpu", true),
		// chromedp.Flag("no-sandbox", true),
		chromedp.Flag("headless", false),    // Run with UI
		chromedp.Flag("disable-gpu", false), // Allow GPU acceleration
		chromedp.Flag("no-sandbox", true),   // Needed for some environments
	)

	allocCtx, _ = chromedp.NewExecAllocator(context.Background(), opts...)
	chromeCtx2, cancelFunc2 := chromedp.NewContext(allocCtx) // Shared Chrome instance
	cancelFunc = cancelFunc2
	chromeCtx = &chromeCtx2
	return chromeCtx
}
func CloseChrome() {
	if cancelFunc != nil {
		cancelFunc() // Close the entire Chrome instance
	}
}

// create default tab

func createDefaultTab() {

	// Create a new tab (isolated session)
	// Create a default tab
	defaultTabCtx, _ = chromedp.NewContext(*chromeCtx)

	fmt.Println(&chromeCtx, 8766)

	// Open a default page (optional)
	go func() {
		err := chromedp.Run(defaultTabCtx,
			chromedp.Navigate("https://google.com"), // Change URL if needed
		)
		if err != nil {
			log.Println("Failed to open default tab:", err)
		} else {
			log.Println("Default tab opened successfully!")
		}
	}()

}

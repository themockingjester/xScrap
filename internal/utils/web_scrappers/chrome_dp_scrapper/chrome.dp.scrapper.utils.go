package chrome_dp_scrapper

import (
	"context"
	"fmt"
	"log"
	"sync"
	"time"

	"github.com/chromedp/chromedp"
)

type ChromeDPScrapperStruct struct {
	mu              sync.Mutex
	closed          bool
	initialized     bool
	rootCtx         context.Context
	rootCancel      context.CancelFunc
	parentCtxCancel context.CancelFunc
}

// NewChromeDPScrapper creates a new Chrome DP scraper instance
func NewChromeDPScrapper() *ChromeDPScrapperStruct {
	return &ChromeDPScrapperStruct{
		closed:          false,
		initialized:     false,
		rootCtx:         nil,
		rootCancel:      nil,
		parentCtxCancel: nil,
	}
}

// Initialize initializes the browser window (should be called once)
// Initialize initializes the browser window (should be called once)
func (c *ChromeDPScrapperStruct) Initialize() error {
	c.mu.Lock()
	defer c.mu.Unlock()

	if c.closed {
		return fmt.Errorf("scraper is already closed")
	}

	if c.initialized {
		return nil // Already initialized
	}

	// Allocate Chrome with options (visible browser)
	opts := append(chromedp.DefaultExecAllocatorOptions[:],
		chromedp.Flag("headless", false),       // keep visible
		chromedp.Flag("start-maximized", true), // maximize window
	)

	allocCtx, cancelCtx := chromedp.NewExecAllocator(context.Background(), opts...)
	c.parentCtxCancel = cancelCtx
	c.rootCtx, c.rootCancel = chromedp.NewContext(allocCtx)

	// Keep a dummy tab open indefinitely
	if err := chromedp.Run(c.rootCtx, chromedp.Navigate("about:blank")); err != nil {
		return fmt.Errorf("failed to open persistent tab: %w", err)
	}

	c.initialized = true
	log.Println("Chrome DP scraper initialized successfully (persistent tab open)")
	return nil
}

// ensureInitialized ensures the scraper is initialized before use
func (c *ChromeDPScrapperStruct) ensureInitialized() error {
	c.mu.Lock()
	defer c.mu.Unlock()

	if c.closed {
		return fmt.Errorf("scraper is closed")
	}

	if !c.initialized {
		// Try to initialize now
		if err := c.Initialize(); err != nil {
			return fmt.Errorf("failed to initialize scraper: %w", err)
		}
	}

	return nil
}

// Close closes the browser window and cleans up resources
func (c *ChromeDPScrapperStruct) Close() error {
	c.mu.Lock()
	defer c.mu.Unlock()

	if c.closed {
		return nil
	}
	defer c.rootCancel()
	defer c.parentCtxCancel()
	c.closed = true
	c.initialized = false
	log.Println("Chrome DP scraper closed")
	return nil
}

// ScrapDataUsingCSSSelector scrapes data using CSS selector
func (c *ChromeDPScrapperStruct) ScrapDataUsingCSSSelector(url string, cssSelector string) (string, error) {

	if err := c.ensureInitialized(); err != nil {
		return "", err
	}
	if c.closed {
		return "", fmt.Errorf("scraper is closed")
	}

	// Create a new tab from the root browser context
	tabCtx, tabCancel := chromedp.NewContext(c.rootCtx)
	// IMPORTANT: don’t close the tab immediately, let Chrome handle it
	// If you cancel here, the tab closes as soon as function exits
	//defer tabCancel()

	// Add timeout to this tab only
	ctx, cancel := context.WithTimeout(tabCtx, 15*time.Second)
	defer cancel()

	var result string
	tasks := chromedp.Tasks{
		chromedp.Navigate(url),
		chromedp.WaitVisible(cssSelector, chromedp.ByQuery),
		chromedp.Text(cssSelector, &result, chromedp.ByQuery),
	}

	if err := chromedp.Run(ctx, tasks); err != nil {
		tabCancel() // close tab only if failed
		return "", fmt.Errorf("failed to scrape using XPath: %w", err)
	}

	// Explicitly close this tab after work
	tabCancel()

	return result, nil
}

// ScrapDataUsingXPath scrapes data using XPath
func (c *ChromeDPScrapperStruct) ScrapDataUsingXPath(url string, xPath string) (string, error) {
	if err := c.ensureInitialized(); err != nil {
		return "", err
	}
	if c.closed {
		return "", fmt.Errorf("scraper is closed")
	}

	// Create a new tab from the root browser context
	tabCtx, tabCancel := chromedp.NewContext(c.rootCtx)
	// IMPORTANT: don’t close the tab immediately, let Chrome handle it
	// If you cancel here, the tab closes as soon as function exits
	//defer tabCancel()

	// Add timeout to this tab only
	ctx, cancel := context.WithTimeout(tabCtx, 15*time.Second)
	defer cancel()

	var result string
	tasks := chromedp.Tasks{
		chromedp.Navigate(url),
		chromedp.WaitVisible(xPath, chromedp.BySearch),
		chromedp.Text(xPath, &result, chromedp.BySearch),
	}

	if err := chromedp.Run(ctx, tasks); err != nil {
		tabCancel() // close tab only if failed
		return "", fmt.Errorf("failed to scrape using XPath: %w", err)
	}

	// Explicitly close this tab after work
	tabCancel()

	return result, nil
}

// ScrapDataUsingId scrapes data using element ID
func (c *ChromeDPScrapperStruct) ScrapDataUsingId(url string, id string) (string, error) {

	if err := c.ensureInitialized(); err != nil {
		return "", err
	}
	if c.closed {
		return "", fmt.Errorf("scraper is closed")
	}

	// Create a new tab from the root browser context
	tabCtx, tabCancel := chromedp.NewContext(c.rootCtx)
	// IMPORTANT: don’t close the tab immediately, let Chrome handle it
	// If you cancel here, the tab closes as soon as function exits
	//defer tabCancel()

	// Add timeout to this tab only
	ctx, cancel := context.WithTimeout(tabCtx, 15*time.Second)
	defer cancel()

	var result string

	selector := "#" + id
	tasks := chromedp.Tasks{
		chromedp.Navigate(url),
		chromedp.WaitVisible(selector, chromedp.ByQuery),
		chromedp.Text(selector, &result, chromedp.ByQuery),
	}

	if err := chromedp.Run(ctx, tasks); err != nil {
		tabCancel() // close tab only if failed
		return "", fmt.Errorf("failed to scrape using XPath: %w", err)
	}

	// Explicitly close this tab after work
	tabCancel()

	return result, nil

}

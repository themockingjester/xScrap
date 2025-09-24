package server_setup

import (
	config "github.com/xscrap/configs"
	"github.com/xscrap/internal/utils/web_scrappers/chrome_dp_scrapper"

	"log"
)

func SetupServer(appDIContainer *config.AppDIContainer, appConfig *config.AppConfig) {

	// Setting up web scrapper
	chromeScraper := chrome_dp_scrapper.NewChromeDPScrapper()

	// Try to initialize the browser window, but don't fail if it doesn't work
	// The scraper will initialize itself when first used
	if err := chromeScraper.Initialize(); err != nil {
		log.Printf("Warning: Failed to initialize Chrome DP scraper during startup: %v", err)
		log.Println("Chrome DP scraper will be initialized on first use")
	} else {
		log.Println("✅ Chrome DP scraper initialized during startup")
	}
	appDIContainer.WebScrapper.Scrapper = chromeScraper

	log.Println("✅ Server setup complete!")
}

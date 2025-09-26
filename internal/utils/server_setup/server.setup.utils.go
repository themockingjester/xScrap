package server_setup

import (
	"fmt"
	"time"

	config "github.com/xscrap/configs"
	"github.com/xscrap/internal/utils/machine"
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
	ManageDynamicConcurrencyLimiter(appDIContainer, appConfig)
	log.Println("✅ Server setup complete!")
}

func ManageDynamicConcurrencyLimiter(appDIContainer *config.AppDIContainer, appConfig *config.AppConfig) {

	limiter := machine.NewDynamicLimiter(appConfig.Server.InitialMaxAllowedScrappingWindows) // start with 3
	appDIContainer.DynamicLimiter = limiter
	go func() {
		for {
			ram := machine.GetAvailableRAMMB()
			fmt.Println(ram, 6555)
			if ram < uint64(appConfig.Server.DynamicConcurrencyRamThreshold) { // if less than Allocated Threshold
				limiter.UpdateLimit(-2, true) // restrict concurrency
			} else {
				limiter.UpdateLimit(10, true) // allow more when RAM free
			}
			time.Sleep(time.Duration(appConfig.Server.TimeIntervalForAvailableRAMCheck) * time.Second)
		}
	}()
}

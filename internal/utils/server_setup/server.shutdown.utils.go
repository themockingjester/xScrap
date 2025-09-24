package server_setup

import (
	"log"

	config "github.com/xscrap/configs"
	"github.com/xscrap/internal/utils/web_scrappers/chrome_dp_scrapper"
)

func ShutDownServer(appDIContainer *config.AppDIContainer, appConfig *config.AppConfig) (bool, error) {

	log.Println("🔄 Shutting down server...")

	// Close the Chrome DP scraper if it exists
	if appDIContainer.WebScrapper.Scrapper != nil {
		if chromeScraper, ok := appDIContainer.WebScrapper.Scrapper.(*chrome_dp_scrapper.ChromeDPScrapperStruct); ok {
			if err := chromeScraper.Close(); err != nil {
				log.Printf("Error closing Chrome DP scraper: %v", err)
			} else {
				log.Println("✅ Chrome DP scraper closed successfully")
			}
		}
	}

	log.Println("Would run cleanup tasks...")
	appConfig.WaitGroup.Wait()
	log.Println("✅ Server shutdown complete!")
	return true, nil
}

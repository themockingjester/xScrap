package config

import (
	"context"
	"sync"

	"github.com/xscrap/internal/utils/machine"
	web_scrapper "github.com/xscrap/internal/utils/web_scrappers/scrapper"
)

// Define the config struct
type AppConfig struct {
	WaitGroup *sync.WaitGroup

	Server struct {
		Port                              int    `yaml:"port"`
		Environment                       string `yaml:"environment"`
		DynamicConcurrencyRamThreshold    int    `yaml:"dynamicConcurrencyRamThreshold"`
		InitialMaxAllowedScrappingWindows int    `yaml:"initialMaxAllowedScrappingWindows"`
		TimeIntervalForAvailableRAMCheck  int    `yaml:"timeIntervalForAvailableRAMCheck"`
	} `yaml:"server"`
}

type WebScrapper struct {
	Strategy   string
	Scrapper   web_scrapper.ScrapperInterface
	Ctx        context.Context
	CancleFunc context.CancelFunc
}

type AppDIContainer struct {
	ShuttingDown   bool
	WebScrapper    WebScrapper
	DynamicLimiter *machine.DynamicLimiter
}

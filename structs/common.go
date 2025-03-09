package structs

import "sync"

type ScrapDataObject struct {
	Url     string
	ByXPath string
	ByClass string
	ById    string
}

type AppConfig struct {
	WaitGroup *sync.WaitGroup
}

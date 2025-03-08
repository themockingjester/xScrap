package structs

type ScrapDataInputObject struct {
	Url     string
	ById    string
	ByXPath string
	ByClass string
}

type ScrapDataUsingXPath struct {
	Url     string
	ByXPath string
}

type ScrapDataDefaultResult struct {
	Url  string
	Mode string
	Data any
}

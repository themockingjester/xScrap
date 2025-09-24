package constants

type ResultObject struct {
	Success bool
	Message string
	Data    map[string]any
	Code    int
	Error   string
}

type ResponseObject struct {
	Success bool
	Message string
	Data    map[string]any
	Code    int
	Error   string
}

package constants

type InternalErrorDetail struct {
	Name string
	Code string
}
type errorsInternalToApp struct {
	NO_ERROR                map[string]InternalErrorDetail
	USER_WEB_SOCKET_RELATED map[string]InternalErrorDetail
	DATABASE_RELATED        map[string]InternalErrorDetail
	JSON_RELATED            map[string]InternalErrorDetail
	EXECUTION_RELATED       map[string]InternalErrorDetail
	CACHE_RELATED           map[string]InternalErrorDetail
}

var ERRORS_INTERNAL_TO_APP = errorsInternalToApp{
	NO_ERROR: map[string]InternalErrorDetail{
		"SUCCESS": {
			Name: "SUCCESS",
			Code: "0x200",
		},
	},
	USER_WEB_SOCKET_RELATED: map[string]InternalErrorDetail{
		"SOCKET_NOT_FOUND": {
			Name: "SOCKET_NOT_FOUND",
			Code: "1X404",
		},
		"RETRY_QUEUE_FULL": {
			Name: "RETRY_QUEUE_FULL",
			Code: "1XA500",
		},
	},
	DATABASE_RELATED: map[string]InternalErrorDetail{},
	JSON_RELATED: map[string]InternalErrorDetail{
		"JSON_MARSHLLING_ERROR": {
			Name: "JSON_MARSHLLING_ERROR",
			Code: "3xA500",
		},
	},
	EXECUTION_RELATED: map[string]InternalErrorDetail{
		"CACHED_KEY_PARSING_FAILURE": {
			Name: "CACHED_KEY_PARSING_FAILURE",
			Code: "4xA500",
		},
		"MESSAGE_CORRECT_REDIRECTION_FAILURE": {
			Name: "MESSAGE_CORRECT_REDIRECTION_FAILURE",
			Code: "4xB500",
		},
		"WHILE_PREPARING_MESSAGE_FOR_REDIRECTION": {
			Name: "WHILE_PREPARING_MESSAGE_FOR_REDIRECTION",
			Code: "4xC500",
		},
	},
	CACHE_RELATED: map[string]InternalErrorDetail{
		"UNABLE_TO_GET_DATA_USING_KEY": {
			Name: "UNABLE_TO_GET_DATA_USING_KEY",
			Code: "5xA500",
		},
	},
}

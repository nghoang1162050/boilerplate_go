package constants

import "errors"

const (
	// output status codes
	STATUS_CODE_SERVICE_SUCCESS                            string = "20001"
	STATUS_CODE_DELETE_SUCCESS                             string = "20201"
	STATUS_CODE_ERROR_BINDING_BODY                         string = "40002"
	STATUS_CODE_INVALID_TRANSACTION                        string = "40003"
	STATUS_CODE_PRIMARY_KEY_REQUIRED                       string = "40004"
	STATUS_CODE_MODEL_VALUE_REQUIRED                       string = "40005"
	STATUS_CODE_UNSUPPORTED_DRIVER                         string = "40006"
	STATUS_CODE_REGISTERED                                 string = "40007"
	STATUS_CODE_INVALID_FIELD                              string = "40008"
	STATUS_CODE_INVALID_DATA                               string = "40009"
	STATUS_CODE_INVALID_DB                                 string = "40010"
	STATUS_CODE_INVALID_VALUE                              string = "40011"
	STATUS_CODE_NOT_IMPLEMENTED                            string = "40012"
	STATUS_CODE_MISSING_WHERE_CLAUSE                       string = "40013"
	STATUS_CODE_UNSUPPORTED_RELATION                       string = "40014"
	STATUS_CODE_EMPTY_SLICE                                string = "40015"
	STATUS_CODE_DRY_RUN_UNSUPPORTED                        string = "40016"
	STATUS_CODE_INVALID_VALUE_LENGTH                       string = "40017"
	STATUS_CODE_PRELOAD_NOT_ALLOWED                        string = "40018"
	STATUS_CODE_VALIDATION_ERROR                           string = "40019"
	STATUS_CODE_IDS_NOT_READABLE                           string = "40020"
	STATUS_CODE_NOT_AUTHORIZED_WITHOUT_HEADER              string = "40101"
	STATUS_CODE_NOT_AUTHORIZED                             string = "40102"
	STATUS_CODE_ROUTE_NOT_FOUND                            string = "40401"
	STATUS_CODE_RECORD_NOT_FOUND                           string = "40402"
	STATUS_CODE_DEPENDENCY_NOT_FOUND                       string = "40403"
	STATUS_CODE_ID_NOT_FOUND                               string = "40404"
	STATUS_CODE_NOT_ACCEPTABLE_WITHOUT_ACCEPT_HEADER       string = "40601"
	STATUS_CODE_NOT_ACCEPTABLE_WITHOUT_CONTENT_TYPE_HEADER string = "40602"
	STATUS_CODE_INTERNAL_SERVER_ERROR                      string = "50001"
	STATUS_CODE_FAILED_TO_DECODE_VALUE                     string = "50011"

	// output messages
	MSG_SERVER_SHUTTING_DOWN        string = "server is shutting down"
	MSG_NOT_ACCEPTABLE              string = "not acceptable"
	MSG_MISSING_ACCEPT_HEADER       string = "unknown accept format"
	MSG_MISSING_CONTENT_TYPE_HEADER string = "unknown content format"
	MSG_SUCCESS                     string = "success"
	MSG_ERROR                       string = "error"
	MSG_VALIDATION_ERROR            string = "validation error"
	MSG_ROUTE_NOT_FOUND             string = "route not found"
	MSG_RECORD_NOT_FOUND            string = "record not found"
	MSG_DEPENDENCY_NOT_FOUND        string = "dependency not found"
	MSG_SESSION_NOT_FOUND           string = "session not found"
	MSG_ACCESS_IDS_NOT_FOUND        string = "access ids not found or not readable"
	MSG_NOT_AUTHORIZED              string = "not authorized"
	MSG_ID_NOT_READABLE             string = "ID not found or not readable"
	MSG_UNABLE_TO_BIND_BODY         string = "error binding body"
	MSG_FORBIDDEN                   string = "forbidden"
	MSG_UNKNOWN_DB_PLATFORM         string = "unknown database platform"
	MSG_INTERNAL_SERVER             string = "internal server error"

	// TODO: remove later
	JWT_SECRET_KEY string = "78df74878a73fc1bbad6f30b4458eee12298b6faa8cc3063ac680febdea40fc2"
)

var (
	// log levels
	LOG_LEVELS []string = []string{"debug", "info", "warn", "error", "fatal", "panic"}

	// custom errors
	ERROR_NOT_AUTHORIZED       = errors.New(MSG_NOT_AUTHORIZED)
	ERROR_SESSION_NOT_FOUND    = errors.New(MSG_SESSION_NOT_FOUND)
	ERROR_ID_NOT_FOUND         = errors.New(MSG_ACCESS_IDS_NOT_FOUND)
	ERROR_ACCESS_IDS_NOT_FOUND = errors.New(MSG_SESSION_NOT_FOUND)
	ERROR_BINDING_BODY         = errors.New(MSG_UNABLE_TO_BIND_BODY)
	ERROR_UNKNOWN_DB_PLATFORM  = errors.New(MSG_UNKNOWN_DB_PLATFORM)
	ERROR_INTERNAL_SERVER      = errors.New(MSG_INTERNAL_SERVER)
)
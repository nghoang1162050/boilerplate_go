package handlers

type ApiResponse struct {
	Code    string      `json:"code"`
	Message string      `json:"message"`
	Errors  []string    `json:"errors"`
	Payload interface{} `json:"payload"`
}

func BuildResponse(code string, message string, errors []string, payload interface{}) *ApiResponse {
	r := &ApiResponse{}
	r.Code = code
	r.Message = message
	r.Errors = errors
	r.Payload = payload
	return r
}

// func GetStatusCode(err error) int {
// 	if err == nil {
// 		return http.StatusOK
// 	}

// 	switch err {
// 	case domain.ErrInternalServerError:
// 		return http.StatusInternalServerError
// 	case domain.ErrNotFound:
// 		return http.StatusNotFound
// 	case domain.ErrConflict:
// 		return http.StatusConflict
// 	default:
// 		return http.StatusInternalServerError
// 	}
// }
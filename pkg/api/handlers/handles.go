package handlers

func Success(payload interface{}) *ApiResponse {
	return BuildResponse("200", "Success", []string{}, payload)
}

func Accepted() *ApiResponse {
	return BuildResponse(
		"200",
		"Success",
		[]string{},
		nil)
}

func Error(status_code string, err error) *ApiResponse {
	return BuildResponse(
		status_code,
		"Error",
		[]string{err.Error()},
		nil)
}
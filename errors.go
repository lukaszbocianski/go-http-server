package server

type ApiError struct {
	Code string `json:"Code"`
	S    string `json:"Message"`
}

func (o ApiError) Error() string {
	return o.S
}

func ErrorString(errorMessage string) ApiError {
	return ApiError{S: errorMessage}
}

func ErrorDb(errorCode string, errorMessage string) ApiError {
	return ApiError{S: errorMessage, Code: errorCode}
}

package server

type ApiError struct {
	Code string `json:"Code"`
	S    string `json:"Message"`
}

func (o ApiError) Error() string {
	return o.S
}

func NewErrorS(errorMessage string) ApiError {
	return ApiError{S: errorMessage}
}

func NewError(errorCode string, errorMessage string) ApiError {
	return ApiError{S: errorMessage, Code: errorCode}
}

//common errors

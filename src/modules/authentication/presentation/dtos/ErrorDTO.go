package dtos

type ErrorDto struct {
	StatusCode   int    `json:"statusCode"`
	ErrorMessage string `json:"message"`
}

func NewErrorDto(statusCode int, message string) *ErrorDto {
	return &ErrorDto{
		StatusCode:   statusCode,
		ErrorMessage: message,
	}
}

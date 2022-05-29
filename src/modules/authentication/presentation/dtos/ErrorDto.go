package dtos

type ErrorDto struct {
	ErrorMessage string `json:"message"`
}

func NewErrorDto(message string) *ErrorDto {
	return &ErrorDto{
		ErrorMessage: message,
	}
}

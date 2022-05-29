package errors

type ValidationError struct {
	message string
}

func (v *ValidationError) Message() string {
	return v.message
}

func NewValidationError(message string) DomainError {
	return &ValidationError{
		message: message,
	}
}

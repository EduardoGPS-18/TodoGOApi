package errors

type SessionExpiredError struct {
	message string
}

func (s *SessionExpiredError) Message() string {
	return s.message
}

func NewSessionExpiredError(message string) DomainError {
	return &SessionExpiredError{
		message: message,
	}
}

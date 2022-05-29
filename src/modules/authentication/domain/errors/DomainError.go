package errors

type DomainError interface {
	Message() string
}

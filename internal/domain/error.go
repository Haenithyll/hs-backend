package domain

type ErrorCode string

const (
	ErrNotFound            ErrorCode = "NOT_FOUND"
	ErrBadRequest          ErrorCode = "BAD_REQUEST"
	ErrForbidden           ErrorCode = "FORBIDDEN"
	ErrInternalServerError ErrorCode = "INTERNAL_SERVER_ERROR"
)

type DomainError struct {
	Code    ErrorCode
	Message string
}

func NewDomainError(code ErrorCode, msg string) *DomainError {
	return &DomainError{
		Code:    code,
		Message: msg,
	}
}

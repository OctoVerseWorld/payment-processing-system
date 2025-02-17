package internal

import (
	"fmt"
)

// ErrorCode ...
type ErrorCode int

const (
	ErrorCodeUnknown ErrorCode = iota
	ErrorCodeDecoding
	ErrorCodeNotFound
	ErrorCodeInvalidArgument
	UserIDNotProvided
)

// Error ...
type Error struct {
	innerErr  error
	errorCode ErrorCode
	message   string
}

// NewErrorf returns a new error.
func NewErrorf(code ErrorCode, message string, a ...interface{}) *Error {
	return &Error{
		innerErr:  nil,
		errorCode: code,
		message:   fmt.Sprintf(message, a...),
	}
}

// WrapErrorf returns a wrapped error.
func WrapErrorf(err error, code ErrorCode, message string, a ...interface{}) error {
	return &Error{
		innerErr:  err,
		errorCode: code,
		message:   fmt.Sprintf(message, a...),
	}
}

// Error implements the error interface
func (e *Error) Error() string {
	return fmt.Sprintf("%s", e.message)
}

// GetOrig returns the original error
func (e *Error) GetOrig() string {
	return fmt.Sprintf("%v", e.innerErr)
}

// Unwrap returns the wrapped error, if any.
func (e *Error) Unwrap() error {
	return e.innerErr
}

// GetCode returns the error code
func (e *Error) GetCode() ErrorCode {
	return e.errorCode
}

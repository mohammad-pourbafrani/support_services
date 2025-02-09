package types

import (
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

// code 1001 insert db in authentication service
// code 1002 response insert db in authentication service
// code 1003 response select db in authentication service
// code 1004
// code 1005 select user data issue in user repository
// code 1006 update user.verify data issue in user repository
// code 1007 insert token data issue in token repository
// code 1008 generate access token issue in token service
// code 1009 generate refresh token issue in token service

// Error represents an application-specific error with a message and code.
type Error struct {
	Message string // Error message
	Code    uint16 // Error code
}

// NewError creates a new custom error with the given code and message.
func NewError(code uint16, message string) *Error {
	return &Error{Code: code, Message: message}
}

// NewInternalError creates a new internal server error with the given message.
func NewInternalError(message string) *Error {
	return &Error{Code: 500, Message: message}
}

// NewNotFoundError creates a new not found error with the given message.
func NewNotFoundError(message string) *Error {
	return &Error{Code: 404, Message: message}
}

// NewPermissionDeniedError creates a new permission denied error with the given message.
func NewPermissionDeniedError(message string) *Error {
	return &Error{Code: 700, Message: message}
}

// NewBadRequestError creates a new bad request error with the given message.
func NewBadRequestError(message string) *Error {
	return &Error{Code: 300, Message: message}
}

// ErrorToGRPCStatus converts the custom error to a gRPC status error based on the error code.
func (c *Error) ErrorToGRPCStatus() error {
	switch c.Code {
	case 500:
		return status.Error(codes.Internal, c.Message)
	case 404:
		return status.Error(codes.NotFound, c.Message)
	case 700:
		return status.Error(codes.PermissionDenied, c.Message)
	case 300:
		return status.Error(codes.Aborted, c.Message)
	default:
		return status.Error(codes.Aborted, c.Message)
	}
}

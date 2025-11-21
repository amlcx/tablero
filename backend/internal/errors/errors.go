package errors

import (
	"errors"
	"fmt"
	"net/http"

	"connectrpc.com/connect"
	repo "github.com/amlcx/tablero/backend/internal/repo"
)

type ErrorCode string

const (
	CodeNotFound         ErrorCode = "NOT_FOUND"
	CodeInvalidInput     ErrorCode = "INVALID_INPUT"
	CodeInternal         ErrorCode = "INTERNAL"
	CodePermission       ErrorCode = "PERMISSION_DENIED"
	CodeUnauthenticated  ErrorCode = "AUTHORIZATION_FAILED"
	CodeContextCancelled ErrorCode = "CONTEXT_CANCELLED"
	CodeConflict         ErrorCode = "CONFLICT"
)

type AppError struct {
	Code    ErrorCode
	Message string
	Wrapped error
	Field   string
}

func (e *AppError) Error() string {
	if e.Wrapped != nil {
		return fmt.Sprintf("%s: %s", e.Message, e.Wrapped.Error())
	}
	return e.Message
}

func (e *AppError) Unwrap() error {
	return e.Wrapped
}

func (e *AppError) HTTPStatus() int {
	switch e.Code {
	case CodeNotFound:
		return http.StatusNotFound
	case CodeInvalidInput:
		return http.StatusBadRequest
	case CodePermission:
		return http.StatusForbidden
	case CodeConflict:
		return http.StatusConflict
	case CodeUnauthenticated:
		return http.StatusUnauthorized
	default:
		return http.StatusInternalServerError
	}
}

func (e *AppError) ConnectRPCStatus() connect.Code {
	switch e.Code {
	case CodeNotFound:
		return connect.CodeNotFound
	case CodeInvalidInput:
		return connect.CodeInvalidArgument
	case CodePermission:
		return connect.CodePermissionDenied
	case CodeConflict:
		return connect.CodeAlreadyExists
	case CodeUnauthenticated:
		return connect.CodeUnauthenticated
	default:
		return connect.CodeInternal
	}
}

func NotFound(message string, err error) *AppError {
	return &AppError{
		Code:    CodeNotFound,
		Message: message,
		Wrapped: err,
	}
}

func InvalidInput(message string, err error) *AppError {
	return &AppError{
		Code:    CodeInvalidInput,
		Message: message,
		Wrapped: err,
	}
}

func Internal(message string, err error) *AppError {
	return &AppError{
		Code:    CodeInternal,
		Message: message,
		Wrapped: err,
	}
}

func Permission(message string, err error) *AppError {
	return &AppError{
		Code:    CodePermission,
		Message: message,
		Wrapped: err,
	}
}

func Authorization(message string, err error) *AppError {
	return &AppError{
		Code:    CodeUnauthenticated,
		Message: message,
		Wrapped: err,
	}
}

func ContextCancelled(message string, err error) *AppError {
	return &AppError{
		Code:    CodeContextCancelled,
		Message: message,
		Wrapped: err,
	}
}

func Conflict(message string, err error) *AppError {
	return &AppError{
		Code:    CodeConflict,
		Message: message,
		Wrapped: err,
	}
}

func FromError(err error) *AppError {
	var appErr *AppError

	if errors.As(err, &appErr) {
		return appErr
	}

	var dbErr *repo.DBErr

	if !errors.As(err, &dbErr) {
		// at this point, it's neither an app error nor a database error
		return Internal("unexpected_error", err)
	}

	switch dbErr.Code {
	case repo.NotNull:
		return InvalidInput(dbErr.Message, dbErr)
	case repo.UniqueConflict:
		return Conflict(dbErr.Message, dbErr)
	case repo.TooLong:
		return InvalidInput(dbErr.Message, dbErr)
	case repo.ForeignKey:
		return InvalidInput(dbErr.Message, dbErr)
	default:
		return Internal("unexpected_error", err)
	}
}

func WrapErrors(parent, child error) error {
	return fmt.Errorf("%w: %w", parent, child)
}

func New(msg string) error {
	return errors.New(msg)
}

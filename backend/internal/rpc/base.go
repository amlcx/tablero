package rpc

import (
	"connectrpc.com/connect"
	"github.com/amlcx/tablero/backend/internal/errors"
	"github.com/amlcx/tablero/backend/sentinel"
	"github.com/charmbracelet/log"
)

type BaseHandler interface {
	log() *log.Logger
	normalize(err error) error
}

type baseHandler struct {
	logger *log.Logger
}

var _ BaseHandler = (*baseHandler)(nil)

func NewBaseHandler(
	logger *log.Logger,
) BaseHandler {
	sentinel.Assert(logger != nil, "failed to initialize base handler: nil logger")

	return &baseHandler{
		logger: logger,
	}
}

func (h *baseHandler) log() *log.Logger {
	return h.logger
}

func (h *baseHandler) normalize(err error) error {
	if err == nil {
		return nil
	}

	appErr := errors.FromError(err)

	return connect.NewError(appErr.ConnectRPCStatus(), err)
}

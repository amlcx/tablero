package rpc

import (
	"context"
	"fmt"

	apiv1 "github.com/amlcx/tablero/backend/gen/api/v1"
	"github.com/amlcx/tablero/backend/internal/auth"
	"github.com/amlcx/tablero/backend/sentinel"
)

type GreetHandler interface {
	Greet(ctx context.Context, req *apiv1.GreetRequest) (*apiv1.GreetResponse, error)
}

type greetHandler struct {
	BaseHandler
}

var _ GreetHandler = (*greetHandler)(nil)

func NewGreetHandler(bh BaseHandler) GreetHandler {
	sentinel.Assert(bh != nil, "failed to initialize greet handler: nil base repo")

	return &greetHandler{
		BaseHandler: bh,
	}
}

func (h *greetHandler) Greet(ctx context.Context, req *apiv1.GreetRequest) (*apiv1.GreetResponse, error) {
	user, err := auth.UserFromContext(ctx)
	if err != nil {
		return nil, h.normalize(err)
	}

	reqID := req.GetGreet().GetId()
	reqName := req.GetGreet().GetName()

	resp := &apiv1.GreetResponse{
		Msg: fmt.Sprintf("Hello, %s (%s), this is golang backend. I see that your ID is %s and your role is %s", reqName, reqID, user.ID.String(), user.Role),
	}

	return resp, nil
}

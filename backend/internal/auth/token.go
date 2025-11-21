package auth

import (
	"context"
	"net/http"

	"connectrpc.com/authn"
	"connectrpc.com/connect"
	"github.com/amlcx/tablero/backend/internal/auth/keys"
	"github.com/amlcx/tablero/backend/internal/dto"
	"github.com/amlcx/tablero/backend/internal/errors"
	"github.com/amlcx/tablero/backend/sentinel"
	"github.com/charmbracelet/log"
	"github.com/google/uuid"
	"github.com/lestrrat-go/jwx/v3/jwt"
)

type JWTMiddleware interface {
	Guard(ctx context.Context, r *http.Request) (any, error)
}

type jwtMiddleware struct {
	keySvc keys.KeyServicer
	logger *log.Logger
}

var _ JWTMiddleware = (*jwtMiddleware)(nil)

func NewJWTMiddleware(
	logger *log.Logger,
	jwksURL string,
) JWTMiddleware {
	sentinel.Assert(logger != nil, "failed to initialize jwt middleware: nil logger")

	keySvc := keys.NewKeyServicer(jwksURL)

	return &jwtMiddleware{
		keySvc: keySvc,
		logger: logger,
	}
}

func (m *jwtMiddleware) Guard(ctx context.Context, r *http.Request) (any, error) {
	keySet, err := m.keySvc.GetKeySet(ctx)
	if err != nil {
		return nil, connect.NewError(
			connect.CodeInternal,
			errors.Internal(
				"jwt middleware failed to get key set",
				err,
			),
		)
	}

	token, err := jwt.ParseRequest(r, jwt.WithKeySet(keySet))
	if err != nil {
		m.logger.Error("jwt middleware error during jwt parsing", "err", err)
		return nil, authn.Errorf("%s", err.Error())
	}

	var ok bool
	idStr, ok := token.Subject()
	if !ok {
		return nil, authn.Errorf("no subject claim in jwt")
	}

	if idStr == "" {
		return nil, authn.Errorf("invalid id in jwt claims: %s", idStr)
	}

	var role string
	if err = token.Get("role", &role); err != nil || role == "" {
		return nil, authn.Errorf("no role claim in jwt")
	}

	id, err := uuid.Parse(idStr)
	if err != nil {
		return nil, authn.Errorf("failed to parse uuid")
	}

	user := dto.UserFromToken{
		ID:   id,
		Role: role,
	}

	return user, nil
}

func UserFromContext(ctx context.Context) (dto.UserFromToken, error) {
	if ctx == nil {
		return dto.UserFromToken{}, errors.InvalidInput("extracting user from context failed: nil context", nil)
	}

	val := authn.GetInfo(ctx)

	user, ok := val.(dto.UserFromToken)

	if !ok {
		return dto.UserFromToken{}, errors.Internal("extracting user from context failed: failed to cast type UserFromToken", nil)
	}

	return user, nil
}

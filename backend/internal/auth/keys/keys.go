package keys

import (
	"context"
	"sync"

	"github.com/amlcx/tablero/backend/sentinel"
	"github.com/lestrrat-go/httprc/v3"
	"github.com/lestrrat-go/jwx/v3/jwk"
)

type KeyServicer interface {
	GetKeySet(ctx context.Context) (jwk.Set, error)
}

type keyServicer struct {
	cache *jwk.Cache
	url   string
	once  sync.Once
}

var _ KeyServicer = (*keyServicer)(nil)

func NewKeyServicer(url string) KeyServicer {
	sentinel.Assert(url != "", "failed to initialize key servicer: empty URL")

	return &keyServicer{
		url: url,
	}
}

func (s *keyServicer) init(ctx context.Context) {
	s.once.Do(func() {
		var err error
		s.cache, err = jwk.NewCache(ctx, httprc.NewClient())
		sentinel.AssertError(err, "failed to initialize keys servicer: failed to create cache")

		err = s.cache.Register(ctx, s.url)
		sentinel.AssertError(err, "failed to initialize keys servicer: failed to register URL")
	})
}

func (s *keyServicer) GetKeySet(ctx context.Context) (jwk.Set, error) {
	s.init(ctx)

	return s.cache.Lookup(ctx, s.url)
}

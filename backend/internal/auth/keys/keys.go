package keys

import (
	"context"
	"sync"
	"time"

	"github.com/amlcx/tablero/backend/sentinel"
	"github.com/charmbracelet/log"
	"github.com/lestrrat-go/httprc/v3"
	"github.com/lestrrat-go/jwx/v3/jwk"
)

type KeyServicer interface {
	GetKeySet(ctx context.Context) (jwk.Set, error)
}

type keyServicer struct {
	cache  *jwk.Cache
	logger *log.Logger
	url    string
	once   sync.Once
}

var _ KeyServicer = (*keyServicer)(nil)

func NewKeyServicer(url string, logger *log.Logger) KeyServicer {
	sentinel.Assert(url != "", "failed to initialize key servicer: empty URL")
	sentinel.Assert(logger != nil, "failed to initialize key servicer: nil logger")

	return &keyServicer{
		logger: logger,
		url:    url,
	}
}

func (s *keyServicer) init(ctx context.Context) {
	s.once.Do(func() {
		var err error
		time.Sleep(5 * time.Second)
		s.cache, err = jwk.NewCache(ctx, httprc.NewClient())
		sentinel.AssertError(err, "failed to initialize keys servicer: failed to create cache")

		err = s.cache.Register(ctx, s.url)
		sentinel.AssertError(err, "failed to initialize keys servicer: failed to register URL")
	})
}

func (s *keyServicer) GetKeySet(ctx context.Context) (jwk.Set, error) {
	s.init(ctx)

	timeout, cancel := context.WithTimeout(ctx, 5*time.Second)
	defer func() {
		s.logger.Debug("key servicer timeout")
		cancel()
	}()

	return s.cache.Lookup(timeout, s.url)
}

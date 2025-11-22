package wiring

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"time"

	"connectrpc.com/authn"
	"github.com/amlcx/tablero/backend/gen/api/v1/apiv1connect"
	"github.com/amlcx/tablero/backend/internal/auth"
	"github.com/amlcx/tablero/backend/internal/rpc"
	"github.com/amlcx/tablero/backend/sentinel"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"golang.org/x/net/http2"
	"golang.org/x/net/http2/h2c"
)

type App struct {
	cfg  *Config
	deps *Dependencies

	router chi.Router
	server *http.Server
}

func NewApp() *App {
	app := &App{}
	var err error

	app.cfg, err = LoadConfig()
	sentinel.AssertError(err, "failed to initialize app")

	app.deps = InitDependencies()

	app.initRouter()
	app.mountRoutes()
	app.initServer()

	return app
}

func (app *App) initRouter() {
	r := chi.NewRouter()
	app.router = r

	app.router.Use(middleware.Logger, middleware.Recoverer)
}

func (app *App) mountRoutes() {
	bh := rpc.NewBaseHandler(app.deps.Logger)
	ch := rpc.NewCategoryHandler(bh)
	gh := rpc.NewGreetHandler(bh)

	categoryPath, categoryHandler := apiv1connect.NewCategoryServiceHandler(ch)
	greetPath, greetHandler := apiv1connect.NewGreetServiceHandler(gh)

	app.router.Mount(categoryPath, categoryHandler)
	app.router.Mount(greetPath, greetHandler)
}

func (app *App) initServer() {
	jwtMiddleware := auth.NewJWTMiddleware(app.deps.Logger, app.cfg.JWKS.URL)
	mid := authn.NewMiddleware(jwtMiddleware.Guard)

	h := mid.Wrap(app.router)

	http2Handler := h2c.NewHandler(h, &http2.Server{})

	addr := fmt.Sprintf("%s:%d", app.cfg.Server.Hostname, app.cfg.Server.Port)

	app.server = &http.Server{
		Addr:    addr,
		Handler: http2Handler,
	}
}

func (app *App) Start() error {
	app.deps.Logger.Info("starting application")

	app.deps.Logger.Info("listening for incoming requests", "addr", app.server.Addr)
	if err := app.server.ListenAndServe(); err != nil {
		if !errors.Is(err, http.ErrServerClosed) {
			app.deps.Logger.Fatal("fatal error during app execution", "err", err)
			return err
		}
	}

	return nil
}

func (app *App) Shutdown() {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	if err := app.server.Shutdown(ctx); err != nil {
		app.deps.Logger.Error("app shutdown error", "err", err)
	}
}

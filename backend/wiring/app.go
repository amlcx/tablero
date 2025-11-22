package wiring

import (
	"context"
	"encoding/json"
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

	app.router.Get("/ping", func(w http.ResponseWriter, r *http.Request) {
		type Response struct {
			Message string `json:"message"`
		}

		resp := Response{
			Message: "pong",
		}

		result, _ := json.Marshal(resp)

		fmt.Fprint(w, result)
	})
}

func (app *App) initServer() {
	app.deps.Logger.Debug("initializing server")
	app.deps.Logger.Debug("creating new jwt middleware", "jwks url", app.cfg.JWKS.URL)
	jwtMiddleware := auth.NewJWTMiddleware(app.deps.Logger, app.cfg.JWKS.URL)
	mid := authn.NewMiddleware(jwtMiddleware.Guard)

	h := mid.Wrap(app.router)

	p := new(http.Protocols)
	p.SetHTTP1(true)
	p.SetUnencryptedHTTP2(true)

	addr := fmt.Sprintf("%s:%d", app.cfg.Server.Hostname, app.cfg.Server.Port)

	app.server = &http.Server{
		Addr:      addr,
		Handler:   h,
		Protocols: p,
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

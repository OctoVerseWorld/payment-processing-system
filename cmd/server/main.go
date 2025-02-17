package main

import (
	"PaymentProcessingSystem/internal"
	"PaymentProcessingSystem/internal/configs"
	"PaymentProcessingSystem/internal/infra/db"
	rest_handlers "PaymentProcessingSystem/internal/infra/rest/handlers"
	rest_middlewares "PaymentProcessingSystem/internal/infra/rest/middlewares"
	"PaymentProcessingSystem/internal/infra/zap_logger"

	"context"
	"errors"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/render"
	"github.com/jackc/pgx/v5/pgxpool"
	"go.uber.org/zap"
)

func main() {
	log := zap_logger.InitLogger()

	errC, err := run()
	if err != nil {
		log.Fatal("Couldn't run", zap.String("error", err.Error()))
	}
	if err := <-errC; err != nil {
		log.Fatal("Error while running", zap.String("error", err.Error()))
	}

	if err != nil {
		return
	}

	err = log.Sync()
}

func run() (<-chan error, error) {

	err := configs.InitConfig(".env")
	if err != nil {
		return nil, internal.WrapErrorf(err, internal.ErrorCodeUnknown, "configs.NewConfig")
	}
	cfg := configs.GetConfig()

	//-

	zap.L().Info("Database: Connecting")
	pool, err := db.NewDB(cfg.DatabaseConfig)
	if err != nil {
		return nil, internal.WrapErrorf(err, internal.ErrorCodeUnknown, "internal.NewPostgreSQL")
	}
	zap.L().Info("Database: Connected")

	//-

	zap.L().Info("Server: Starting", zap.String("address", cfg.ServerConfig.GetAddress()))
	httpServer, err := newServer(mainConfig{
		Cfg: cfg.ServerConfig,
		DB:  pool,
		Middlewares: []func(http.Handler) http.Handler{
			rest_middlewares.OTELTraceMiddleware,
			rest_middlewares.HTTPLoggingMiddleware,
			rest_middlewares.CurrentUserMiddleware,
		},
	})
	if err != nil {
		return nil, internal.WrapErrorf(err, internal.ErrorCodeUnknown, "newServer")
	}
	zap.L().Info("Server: Started")

	errC := make(chan error, 1)

	ctx, stop := signal.NotifyContext(
		context.Background(),
		os.Interrupt,
		syscall.SIGTERM,
		syscall.SIGQUIT,
	)

	go func() {
		<-ctx.Done()

		ctxTimeout, cancel := context.WithTimeout(context.Background(), 5*time.Second)

		defer func() {
			pool.Close()
			err := httpServer.Close()
			if err != nil {
				return
			}

			stop()
			cancel()
			close(errC)
		}()

		httpServer.SetKeepAlivesEnabled(false)

		if err := httpServer.Shutdown(ctxTimeout); err != nil { //nolint: context check
			errC <- err
		}

	}()

	go func() {
		// "ListenAndServe always returns a non-nil error. After Shutdown or Close, the returned error is
		// ErrServerClosed."
		if err := httpServer.ListenAndServe(); err != nil && !errors.Is(err, http.ErrServerClosed) {
			errC <- err
		}
	}()

	return errC, nil
}

type mainConfig struct {
	Cfg         *configs.ServerConfig
	DB          *pgxpool.Pool
	Middlewares []func(next http.Handler) http.Handler
}

func newServer(conf mainConfig) (*http.Server, error) {

	router := chi.NewRouter()
	router.Use(render.SetContentType(render.ContentTypeJSON))

	for _, mw := range conf.Middlewares {
		router.Use(mw)
	}

	//-

	baseHandler := rest_handlers.NewBaseHandler(conf.DB)
	baseHandler.Register(router)

	//-

	return &http.Server{
		Addr:              conf.Cfg.GetAddress(),
		Handler:           router,
		ReadTimeout:       1 * time.Second,
		ReadHeaderTimeout: 1 * time.Second,
		WriteTimeout:      1 * time.Second,
		IdleTimeout:       1 * time.Second,
	}, nil
}

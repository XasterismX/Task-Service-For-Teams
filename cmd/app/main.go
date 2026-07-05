package main

import (
	"Task-Service-For-Teams/config"
	"Task-Service-For-Teams/internal/controller/auth"
	middleware2 "Task-Service-For-Teams/internal/controller/middleware"
	"context"
	"fmt"
	"log/slog"
	"net/http"
	"os"
	"os/signal"
	"syscall"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/spf13/viper"
	"golang.org/x/sync/errgroup"
)

func main() {
	config.MustLoad()

	mainCtx, stop := signal.NotifyContext(context.Background(), os.Interrupt, syscall.SIGTERM)
	defer stop()

	limiter := middleware2.NewIPRateLimiter()

	logger := slog.New(slog.NewTextHandler(os.Stderr, &slog.HandlerOptions{
		Level: slog.LevelDebug,
	}))
	slog.SetDefault(logger)

	router := chi.NewRouter()
	router.Use(middleware.RequestID)
	router.Use(middleware.Logger)
	router.Use(middleware2.RateLimitMiddleware(limiter))
	router.Route("/api/v1", func(r chi.Router) {
		r.Post("/register", auth.Register)
		r.Post("/login", auth.Login)
		r.Route("/teams", func(r chi.Router) {
			r.Route("/{id}", func(r chi.Router) {
				r.Use(middleware2.AuthAdminMiddleware)
				r.Get("/", func(w http.ResponseWriter, r *http.Request) {
					w.Header().Set("Content-Type", "application/json")
					w.WriteHeader(http.StatusOK)
					w.Write([]byte(`{"data": "ok"}`))
				})
			})
		})
		r.Route("/tasks", func(r chi.Router) {

		})

	})

	httpServer := &http.Server{
		Addr:    fmt.Sprintf(":%d", viper.GetInt("app.port")),
		Handler: router,
	}
	g, gCtx := errgroup.WithContext(mainCtx)
	g.Go(func() error {
		logger.Info(fmt.Sprintf("starting http server on port %d", viper.GetInt("app.port")))
		return httpServer.ListenAndServe()

	})
	g.Go(func() error {
		<-gCtx.Done()
		return httpServer.Shutdown(context.Background())
	})
	if err := g.Wait(); err != nil {
		logger.Warn(fmt.Sprintf("exit reason: %s \n", err))
	}

}

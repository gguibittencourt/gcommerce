package server

import (
	"context"
	"fmt"
	"net/http"
	"time"

	"github.com/go-chi/chi/v5"
	"go.uber.org/fx"
)

func NewHTTPRouter() *chi.Mux {
	return chi.NewRouter()
}

func StartHTTPServer(lifecycle fx.Lifecycle, router *chi.Mux) {
	server := &http.Server{
		Addr:         ":8080",
		Handler:      router,
		ReadTimeout:  5000 * time.Second,
		WriteTimeout: 5000 * time.Second,
	}
	lifecycle.Append(fx.Hook{
		OnStart: func(context.Context) error {
			fmt.Println("Starting server...")
			go func() {
				err := server.ListenAndServe()
				if err != nil {
					fmt.Printf("Error starting server: %s\n", err.Error())
				}
			}()
			return nil
		},
		OnStop: func(ctx context.Context) error {
			return server.Shutdown(ctx)
		},
	})
}

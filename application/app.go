package application

import (
	"context"
	"fmt"
	"github.com/redis/go-redis/v9"
	"net/http"
	"time"
)

type App struct {
	router http.Handler
	rdb    *redis.Client
}

func New() *App {
	app := &App{
		router: loadRoutes(),
		//rdb:    redis.NewClient(&redis.Options{Addr: "localhost:6379"}),
		rdb: redis.NewClient(&redis.Options{}),
	}

	return app
}

func (a *App) Start(ctx context.Context) error {
	server := &http.Server{
		Addr:    ":3000",
		Handler: a.router,
	}

	err := a.rdb.Ping(ctx).Err()

	defer func() {
		if err := a.rdb.Close(); err != nil {
			fmt.Printf("Failed to close redis connection: %v", err)
		}
	}()

	fmt.Println("Starting server")

	ch := make(chan error, 1)
	go func() {
		err = server.ListenAndServe()
		if err != nil {
			//ch <- fmt.Errorf("failed to start server: %w", err)
		}
		close(ch)
	}()

	select {
	case err = <-ch:
		return err
	case <-ctx.Done():
		fmt.Println("Shutting down server")
		timeout, cancel := context.WithTimeout(context.Background(), time.Second*20)
		defer cancel()
		err = server.Shutdown(timeout)
		if err != nil {
			return fmt.Errorf("failed to shutdown server: %w", err)
		}
	}

	return nil
}

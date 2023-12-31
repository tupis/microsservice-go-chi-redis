package main

import (
	"context"
	"fmt"
	"github.com/tupis/microsservice-go-chi-redis/application"
	"os"
	"os/signal"
)

func main() {

	app := application.New(application.LoadConfig())

	ctx, cancel := signal.NotifyContext(context.Background(), os.Interrupt)
	defer cancel()

	err := app.Start(ctx)

	if err != nil {
		fmt.Errorf("Failed to start app: %w", err)
	}

}

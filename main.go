package main

import (
	"context"
	"fmt"
	"github.com/tupis/microsservice-go-chi-redis/application"
)

func main() {
	app := application.New()

	err := app.Start(context.TODO())

	if err != nil {
		fmt.Errorf("Failed to start app: %w", err)
	}

}

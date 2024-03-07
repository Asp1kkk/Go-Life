package main

import (
	"context"
	"errors"
	"log"
	"os"

	"github.com/Asp1kkk/Go-Life/internal/application"
)

func main() {
	ctx := context.Background()

	os.Exit(mainWithContext(ctx))
}

func mainWithContext(ctx context.Context) int {
	cfg := application.Config{
		Height: 10,
		Width:  10,
	}
	app := application.New(cfg)
	if err := app.Run(ctx); err != nil {
		switch {
		case errors.Is(err, context.Canceled):
			log.Println("Processing cancelled.")
		default:
			log.Println("Application run error", err)
		}
		return 1
	}
	return 0
}

package application

import (
	"context"
	"fmt"
	"time"

	"github.com/Asp1kkk/Go-Life/pkg/life"
)

type Config struct {
	Width  int
	Height int
}

type Application struct {
	Cfg Config
}

func New(config Config) *Application {
	return &Application{
		Cfg: config,
	}
}

func (a *Application) Run(ctx context.Context) error {
	currentWorld := life.NewWorld(a.Cfg.Height, a.Cfg.Width)
	nextWorld := life.NewWorld(a.Cfg.Height, a.Cfg.Width)
	currentWorld.RandInit(40)

	for {
		fmt.Println(currentWorld)
		life.NextState(currentWorld, nextWorld)
		currentWorld = nextWorld

		select {
		case <-ctx.Done():
			return ctx.Err()
		default:
			time.Sleep(2 * time.Second)
		}

		fmt.Print("\033[H\033[2J")
	}
}

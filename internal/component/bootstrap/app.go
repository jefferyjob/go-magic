package bootstrap

import "context"

type App struct {
	ctx    context.Context
	AppEnv string
}

func NewApp(ctx context.Context) *App {
	return &App{
		ctx: ctx,
	}
}

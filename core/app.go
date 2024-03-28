package core

import (
	"context"
	"github.com/jefferyjob/go-magic/core/kit/conf"
)

type App struct {
	ctx    context.Context
	AppEnv string
	Config *conf.Config
}

func NewApp(ctx context.Context, conf *conf.Config) *App {
	return &App{
		ctx:    ctx,
		AppEnv: conf.App.Env,
		Config: conf,
	}
}

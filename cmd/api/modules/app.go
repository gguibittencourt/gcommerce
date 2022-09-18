package modules

import (
	"go.uber.org/fx"
)

func NewApp() *fx.App {
	options := []fx.Option{
		internalModule,
	}
	return fx.New(
		fx.Options(options...),
	)
}

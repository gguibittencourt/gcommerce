package modules

import (
	"go.uber.org/fx"
)

func NewApp() *fx.App {
	options := []fx.Option{
		internalModule,
		freightModule,
		couponModule,
		orderModule,
	}
	return fx.New(
		fx.Options(options...),
	)
}

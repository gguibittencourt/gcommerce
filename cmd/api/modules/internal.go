package modules

import (
	"go.uber.org/fx"

	"github.com/gguibittencourt/gcommerce/internal/server"
)

var internalModule = fx.Options(
	fx.Invoke(
		server.StartHTTPServer,
	),
)

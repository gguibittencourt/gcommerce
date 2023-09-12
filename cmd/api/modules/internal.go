package modules

import (
	"go.uber.org/fx"

	"github.com/gguibittencourt/gcommerce/internal/database"
	"github.com/gguibittencourt/gcommerce/internal/httpclient"
	"github.com/gguibittencourt/gcommerce/internal/server"
)

var serverDependencies = fx.Provide(
	server.NewHTTPRouter,
	database.NewConnection,
	httpclient.NewRequester,
	httpclient.NewClient,
)

var internalModule = fx.Options(
	serverDependencies,
	fx.Invoke(
		server.StartHTTPServer,
	),
)

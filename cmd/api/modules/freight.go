package modules

import (
	"go.uber.org/fx"

	"github.com/gguibittencourt/gcommerce/app/freight/repositories"
	"github.com/gguibittencourt/gcommerce/internal/httpclient"
)

var freightRepositoriesFactory = fx.Provide(
	func(r httpclient.Client) repositories.Client { return r },
	repositories.NewRepository,
)

var freightModule = fx.Options(
	freightRepositoriesFactory,
)

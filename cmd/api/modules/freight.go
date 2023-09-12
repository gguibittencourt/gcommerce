package modules

import (
	"go.uber.org/fx"

	"github.com/gguibittencourt/gcommerce/app/freight/repositories"
	"github.com/gguibittencourt/gcommerce/app/freight/services"
	"github.com/gguibittencourt/gcommerce/internal/httpclient"
)

var freightRepositoriesFactory = fx.Provide(
	func(r httpclient.Client) repositories.Client { return r },
	repositories.NewRepository,
)

var freightServicesFactory = fx.Provide(
	func(r repositories.Repository) services.Repository { return r },
	services.NewCalculateService,
)

var freightModule = fx.Options(
	freightRepositoriesFactory,
	freightServicesFactory,
)

package modules

import (
	"go.uber.org/fx"

	"github.com/gguibittencourt/gcommerce/app/coupon/repositories"
)

var couponRepositoriesFactory = fx.Provide(
	repositories.NewReader,
)

var couponModule = fx.Options(
	couponRepositoriesFactory,
)

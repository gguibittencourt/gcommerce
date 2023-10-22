package modules

import (
	"net/http"

	"github.com/go-chi/chi/v5"
	"go.uber.org/fx"

	couponRepositories "github.com/gguibittencourt/gcommerce/app/coupon/repositories"
	freightRepositories "github.com/gguibittencourt/gcommerce/app/freight/repositories"
	"github.com/gguibittencourt/gcommerce/app/order/handlers/rabbitmq"
	"github.com/gguibittencourt/gcommerce/app/order/handlers/rest/create"
	"github.com/gguibittencourt/gcommerce/app/order/handlers/rest/find"
	"github.com/gguibittencourt/gcommerce/app/order/repositories"
	"github.com/gguibittencourt/gcommerce/app/order/usecases"
	r "github.com/gguibittencourt/gcommerce/internal/rabbitmq"
)

var repositoriesFactory = fx.Provide(
	repositories.NewReader,
	repositories.NewWriter,
)

var useCasesFactory = fx.Provide(
	func(r couponRepositories.Reader) usecases.CouponReader { return r },
	func(r freightRepositories.Repository) usecases.FreightReader { return r },
	func(r repositories.Writer) usecases.Writer { return r },
	func(r r.Publisher) usecases.Publisher { return r },
	usecases.NewCreateUsecase,
	func(r repositories.Reader) usecases.FindRepository { return r },
	usecases.NewFindUsecase,
)

var handlersFactory = fx.Provide(
	func(s usecases.CreateUsecase) create.UseCase { return s },
	create.NewHandler,
	func(s usecases.FindUsecase) find.UseCase { return s },
	find.NewHandler,
	func(s r.Consumer) rabbitmq.Service { return s },
)

var orderModule = fx.Options(
	repositoriesFactory,
	useCasesFactory,
	handlersFactory,
	fx.Invoke(
		RegisterCreateHandler,
		RegisterGetHandler,
		rabbitmq.NewConsumer,
	),
)

func RegisterCreateHandler(mux *chi.Mux, handler create.Handler) {
	mux.Route("/create", func(route chi.Router) {
		route.Method(http.MethodPost, "/", handler)
	})
}

func RegisterGetHandler(mux *chi.Mux, handler find.Handler) {
	mux.Route("/{order_id}", func(route chi.Router) {
		route.Method(http.MethodGet, "/", handler)
	})
}

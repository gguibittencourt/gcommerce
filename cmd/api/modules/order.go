package modules

import (
	"net/http"

	"github.com/go-chi/chi/v5"
	"go.uber.org/fx"

	"github.com/gguibittencourt/gcommerce/app/freight/services"
	"github.com/gguibittencourt/gcommerce/app/order/handlers/rest/create"
	"github.com/gguibittencourt/gcommerce/app/order/handlers/rest/find"
	"github.com/gguibittencourt/gcommerce/app/order/repositories"
	"github.com/gguibittencourt/gcommerce/app/order/usecases"
)

var repositoriesFactory = fx.Provide(
	repositories.NewReader,
	repositories.NewWriter,
)

var useCasesFactory = fx.Provide(
	func(r services.Service) usecases.FreightService { return r },
	func(r repositories.Writer) usecases.Repository { return r },
	usecases.NewCreateUsecase,
	func(r repositories.Reader) usecases.FindRepository { return r },
	usecases.NewFindUsecase,
)

var handlersFactory = fx.Provide(
	func(s usecases.CreateUsecase) create.Service { return s },
	create.NewHandler,
	func(s usecases.FindUsecase) find.Service { return s },
	find.NewHandler,
)

var orderModule = fx.Options(
	repositoriesFactory,
	useCasesFactory,
	handlersFactory,
	fx.Invoke(
		RegisterCreateHandler,
		RegisterGetHandler,
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

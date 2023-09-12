package create

import (
	"context"
	"encoding/json"
	"net/http"

	"github.com/gguibittencourt/gcommerce/app/order"
)

type (
	UseCase interface {
		Create(ctx context.Context, order order.Order) error
	}
	Handler struct {
		useCase UseCase
	}
)

func NewHandler(useCase UseCase) Handler {
	return Handler{
		useCase: useCase,
	}
}

func (h Handler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	var p Payload
	if err := json.NewDecoder(r.Body).Decode(&p); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	if err := h.useCase.Create(ctx, p.toOrder()); err != nil {
		w.WriteHeader(http.StatusUnprocessableEntity)
		return
	}
	w.WriteHeader(http.StatusCreated)
}

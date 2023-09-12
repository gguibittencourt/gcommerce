package find

import (
	"context"
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/go-chi/chi/v5"

	"github.com/gguibittencourt/gcommerce/app/order"
)

type (
	Service interface {
		FindByID(ctx context.Context, id uint64) (order.Order, error)
	}
	Handler struct {
		service Service
	}
)

func NewHandler(service Service) Handler {
	return Handler{
		service: service,
	}
}

func (h Handler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	param := chi.URLParam(r, "order_id")
	orderID, err := strconv.ParseUint(param, 10, 64)
	if orderID == 0 || err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	o, err := h.service.FindByID(ctx, orderID)
	if err != nil {
		w.WriteHeader(http.StatusUnprocessableEntity)
		return
	}
	marshal, _ := json.Marshal(o)
	w.WriteHeader(http.StatusCreated)
	_, _ = w.Write(marshal)
}

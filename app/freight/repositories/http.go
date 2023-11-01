package repositories

import (
	"context"
	"fmt"
	"time"

	"github.com/gguibittencourt/gcommerce/app/freight"
	"github.com/gguibittencourt/gcommerce/app/order"
	"github.com/gguibittencourt/gcommerce/internal/httpclient"
)

type (
	Client interface {
		Post(ctx context.Context, url string, body any) (httpclient.Response, error)
		Get(ctx context.Context, url string) (httpclient.Response, error)
	}
	Repository struct {
		client Client
	}
)

func NewRepository(client Client) Repository {
	return Repository{
		client: client,
	}
}

func (r Repository) Calculate(ctx context.Context, order order.Order) (freight.Freight, error) {
	//result, err := r.client.Post(ctx, "http://localhost:8080/freight/calculate", order)
	result, err := r.client.Get(ctx, "https://dog.ceo/api/breeds/image/random")
	if err != nil {
		fmt.Println(result)
		return freight.Freight{}, err
	}
	//if err != nil {
	//	return freight.Freight{}, err
	//}
	//return result.(freight.Freight), nil
	return freight.Freight{
		Code:          "123",
		Price:         100,
		DurationInMin: 150,
		ETA:           time.Now().Add(150 * time.Minute),
	}, nil
}

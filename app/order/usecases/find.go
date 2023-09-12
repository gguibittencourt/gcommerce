package usecases

import (
	"context"

	"github.com/gguibittencourt/gcommerce/app/order"
)

type (
	FindRepository interface {
		FindByID(ctx context.Context, id uint64) (order.Order, error)
	}
	FindUsecase struct {
		repository FindRepository
	}
)

func NewFindUsecase(repository FindRepository) FindUsecase {
	return FindUsecase{
		repository: repository,
	}
}

func (s FindUsecase) FindByID(ctx context.Context, id uint64) (order.Order, error) {
	return s.repository.FindByID(ctx, id)
}

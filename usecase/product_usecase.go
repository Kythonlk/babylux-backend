package usecase

import (
	"context"
	"time"

	"go-backend/domain"
)

type productUsecase struct {
	productRepository domain.ProductRepository
	contextTimeout    time.Duration
}

func NewProductUsecase(productRepository domain.ProductRepository, timeout time.Duration) domain.ProductUsecase {
	return &productUsecase{
		productRepository: productRepository,
		contextTimeout:    timeout,
	}
}

func (tu *productUsecase) Create(c context.Context, task *domain.Product) error {
	ctx, cancel := context.WithTimeout(c, tu.contextTimeout)
	defer cancel()
	return tu.productRepository.Create(ctx, task)
}

func (tu *productUsecase) FetchByUserID(c context.Context, userID string) ([]domain.Product, error) {
	ctx, cancel := context.WithTimeout(c, tu.contextTimeout)
	defer cancel()
	return tu.productRepository.FetchByUserID(ctx, userID)
}

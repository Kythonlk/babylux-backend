package usecase

import (
	"context"
	"time"

	"go-backend/domain"
)

type ProductUsecase struct {
	taskRepository domain.TaskRepository
	contextTimeout time.Duration
}

func NewProductUsecase(taskRepository domain.TaskRepository, timeout time.Duration) domain.TaskUsecase {
	return &ProductUsecase{
		taskRepository: taskRepository,
		contextTimeout: timeout,
	}
}

func (tu *ProductUsecase) Create(c context.Context, task *domain.Products) error {
	ctx, cancel := context.WithTimeout(c, tu.contextTimeout)
	defer cancel()
	return tu.taskRepository.Create(ctx, product)
}

func (tu *ProductUsecase) FetchByUserID(c context.Context, userID string) ([]domain.Products, error) {
	ctx, cancel := context.WithTimeout(c, tu.contextTimeout)
	defer cancel()
	return tu.taskRepository.FetchByUserID(ctx, userID)
}

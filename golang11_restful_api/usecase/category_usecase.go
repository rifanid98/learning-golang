package usecase

import "context"

type CategoryUsecase interface {
	Create(ctx context.Context, input *CategoryInput) *CategoryOutput
	Update(ctx context.Context, input *CategoryInput) *CategoryOutput
	Delete(ctx context.Context, id int)
	FindById(ctx context.Context, id int) *CategoryOutput
	FindAll(ctx context.Context) []CategoryOutput
}

type CategoryInput struct {
	Name string `json:"name"`
}

type CategoryOutput struct {
	Id   int    `json:"id"`
	Name string `json:"name"`
}

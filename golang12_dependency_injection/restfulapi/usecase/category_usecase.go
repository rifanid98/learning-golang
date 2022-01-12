package usecase

import (
	"context"
	"golang12_dependency_injection/restfulapi/domain/entity"
)

type CategoryUsecase interface {
	Create(ctx context.Context, input *CategoryInput) *CategoryOutput
	Update(ctx context.Context, input *CategoryInput) *CategoryOutput
	Delete(ctx context.Context, id int)
	FindById(ctx context.Context, id int) *CategoryOutput
	FindAll(ctx context.Context) []CategoryOutput
}

type CategoryInput struct {
	Id   int    `json:"id" validate:"number"`
	Name string `json:"name" validate:"required,min=1,max=200"`
}

type CategoryOutput struct {
	Id   int    `json:"id"`
	Name string `json:"name"`
}

func NewCategoryInput(category *entity.Category) *CategoryInput {
	return &CategoryInput{
		Name: category.Name,
	}
}

func NewCategoryOutput(category *entity.Category) *CategoryOutput {
	return &CategoryOutput{
		Id:   category.Id,
		Name: category.Name,
	}
}

package usecase

import (
	"context"
	"database/sql"
	"github.com/go-playground/validator"
	"golang11_restful_api/common"
	"golang11_restful_api/domain/entity"
	"golang11_restful_api/domain/repository"
)

type CategoryInteractor struct {
	CategoryRepository repository.CategoryRepository
	DB                 *sql.DB
	Validator          *validator.Validate
}

func (uc CategoryInteractor) Create(ctx context.Context, input *CategoryInput) *CategoryOutput {
	err := uc.Validator.Struct(input)
	common.PanicIfError(err)

	tx, err := uc.DB.Begin()
	common.PanicIfError(err)
	defer common.CommitOrRollback(tx)

	category := &entity.Category{
		Name: input.Name,
	}

	category = uc.CategoryRepository.Save(ctx, tx, category)

	return NewCategoryOutput(category)
}

func (uc CategoryInteractor) Update(ctx context.Context, input *CategoryInput) *CategoryOutput {
	err := uc.Validator.Struct(input)
	common.PanicIfError(err)

	tx, err := uc.DB.Begin()
	common.PanicIfError(err)
	defer common.CommitOrRollback(tx)

	category, err := uc.CategoryRepository.FindById(ctx, tx, input.Id)
	common.PanicIfError(err)

	category.Name = input.Name

	category = uc.CategoryRepository.Save(ctx, tx, category)

	return NewCategoryOutput(category)
}

func (uc CategoryInteractor) Delete(ctx context.Context, id int) {
	tx, err := uc.DB.Begin()
	common.PanicIfError(err)
	defer common.CommitOrRollback(tx)

	category, err := uc.CategoryRepository.FindById(ctx, tx, id)
	common.PanicIfError(err)

	uc.CategoryRepository.Delete(ctx, tx, category)
}

func (uc CategoryInteractor) FindById(ctx context.Context, id int) *CategoryOutput {
	tx, err := uc.DB.Begin()
	common.PanicIfError(err)
	defer common.CommitOrRollback(tx)

	category, err := uc.CategoryRepository.FindById(ctx, tx, id)
	common.PanicIfError(err)

	return NewCategoryOutput(category)
}

func (uc CategoryInteractor) FindAll(ctx context.Context) []CategoryOutput {
	tx, err := uc.DB.Begin()
	common.PanicIfError(err)
	defer common.CommitOrRollback(tx)

	categories := uc.CategoryRepository.FindAll(ctx, tx)

	var categoriesOutput []CategoryOutput

	for _, category := range categories {
		categoriesOutput = append(categoriesOutput, *NewCategoryOutput(&category))
	}

	return categoriesOutput
}

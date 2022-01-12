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

func NewCategoryInteractor(
	categoryRepository repository.CategoryRepository,
	DB *sql.DB,
	validator *validator.Validate,
) CategoryUsecase {
	return &CategoryInteractor{
		CategoryRepository: categoryRepository,
		DB:                 DB,
		Validator:          validator,
	}
}

func (uc CategoryInteractor) Create(ctx context.Context, input *CategoryInput) *CategoryOutput {
	err := uc.Validator.Struct(input)
	common.PanicIfError(err)

	tx, err := uc.DB.Begin()
	common.PanicIfError(err)

	category := &entity.Category{
		Name: input.Name,
	}

	category = uc.CategoryRepository.Save(ctx, tx, category)

	common.CommitOrRollback(tx)
	return NewCategoryOutput(category)
}

func (uc CategoryInteractor) Update(ctx context.Context, input *CategoryInput) *CategoryOutput {
	err := uc.Validator.Struct(input)
	common.PanicIfError(err)

	tx, err := uc.DB.Begin()
	common.PanicIfError(err)

	category, err := uc.CategoryRepository.FindById(ctx, tx, input.Id)
	if err != nil {
		panic(common.NewNotFoundError(err.Error()))
	}

	category.Name = input.Name

	category = uc.CategoryRepository.Update(ctx, tx, category)

	common.CommitOrRollback(tx)
	return NewCategoryOutput(category)
}

func (uc CategoryInteractor) Delete(ctx context.Context, id int) {
	tx, err := uc.DB.Begin()
	common.PanicIfError(err)

	category, err := uc.CategoryRepository.FindById(ctx, tx, id)
	if err != nil {
		panic(common.NewNotFoundError(err.Error()))
	}

	uc.CategoryRepository.Delete(ctx, tx, category)
	common.CommitOrRollback(tx)
}

func (uc CategoryInteractor) FindById(ctx context.Context, id int) *CategoryOutput {
	tx, err := uc.DB.Begin()
	common.PanicIfError(err)

	category, err := uc.CategoryRepository.FindById(ctx, tx, id)
	if err != nil {
		panic(common.NewNotFoundError(err.Error()))
	}

	common.CommitOrRollback(tx)
	return NewCategoryOutput(category)
}

func (uc CategoryInteractor) FindAll(ctx context.Context) []CategoryOutput {
	tx, err := uc.DB.Begin()
	common.PanicIfError(err)

	categories := uc.CategoryRepository.FindAll(ctx, tx)

	var categoriesOutput []CategoryOutput

	for _, category := range categories {
		categoriesOutput = append(categoriesOutput, *NewCategoryOutput(&category))
	}

	common.CommitOrRollback(tx)
	return categoriesOutput
}

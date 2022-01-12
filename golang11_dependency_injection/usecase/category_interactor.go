package usecase

import (
	"context"
	"database/sql"
	"github.com/go-playground/validator"
	"golang11_dependency_injection/common/database"
	error2 "golang11_dependency_injection/common/error"
	"golang11_dependency_injection/common/exception"
	"golang11_dependency_injection/domain/entity"
	"golang11_dependency_injection/domain/repository"
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
	error2.PanicIfError(err)

	tx, err := uc.DB.Begin()
	error2.PanicIfError(err)

	category := &entity.Category{
		Name: input.Name,
	}

	category = uc.CategoryRepository.Save(ctx, tx, category)

	database.CommitOrRollback(tx)
	return NewCategoryOutput(category)
}

func (uc CategoryInteractor) Update(ctx context.Context, input *CategoryInput) *CategoryOutput {
	err := uc.Validator.Struct(input)
	error2.PanicIfError(err)

	tx, err := uc.DB.Begin()
	error2.PanicIfError(err)

	category, err := uc.CategoryRepository.FindById(ctx, tx, input.Id)
	if err != nil {
		panic(exception.NewNotFoundError(err.Error()))
	}

	category.Name = input.Name

	category = uc.CategoryRepository.Update(ctx, tx, category)

	database.CommitOrRollback(tx)
	return NewCategoryOutput(category)
}

func (uc CategoryInteractor) Delete(ctx context.Context, id int) {
	tx, err := uc.DB.Begin()
	error2.PanicIfError(err)

	category, err := uc.CategoryRepository.FindById(ctx, tx, id)
	if err != nil {
		panic(exception.NewNotFoundError(err.Error()))
	}

	uc.CategoryRepository.Delete(ctx, tx, category)
	database.CommitOrRollback(tx)
}

func (uc CategoryInteractor) FindById(ctx context.Context, id int) *CategoryOutput {
	tx, err := uc.DB.Begin()
	error2.PanicIfError(err)

	category, err := uc.CategoryRepository.FindById(ctx, tx, id)
	if err != nil {
		panic(exception.NewNotFoundError(err.Error()))
	}

	database.CommitOrRollback(tx)
	return NewCategoryOutput(category)
}

func (uc CategoryInteractor) FindAll(ctx context.Context) []CategoryOutput {
	tx, err := uc.DB.Begin()
	error2.PanicIfError(err)

	categories := uc.CategoryRepository.FindAll(ctx, tx)

	var categoriesOutput []CategoryOutput

	for _, category := range categories {
		categoriesOutput = append(categoriesOutput, *NewCategoryOutput(&category))
	}

	database.CommitOrRollback(tx)
	return categoriesOutput
}

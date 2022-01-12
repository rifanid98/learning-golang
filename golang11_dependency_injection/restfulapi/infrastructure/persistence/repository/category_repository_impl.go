package repository

import (
	"context"
	"database/sql"
	"errors"
	error2 "golang11_dependency_injection/common/error"
	"golang11_dependency_injection/domain/entity"
	"golang11_dependency_injection/domain/repository"
)

type CategoryRepositoryImpl struct{}

func NewCategoryRepositoryImpl() repository.CategoryRepository {
	return &CategoryRepositoryImpl{}
}

func (c CategoryRepositoryImpl) Save(ctx context.Context, tx *sql.Tx, category *entity.Category) *entity.Category {
	query := "INSERT INTO category(name) VALUES (?)"

	result, err := tx.ExecContext(ctx, query, category.Name)
	error2.PanicIfError(err)

	id, err := result.LastInsertId()
	error2.PanicIfError(err)

	category.Id = int(id)
	return category
}

func (c CategoryRepositoryImpl) Update(ctx context.Context, tx *sql.Tx, category *entity.Category) *entity.Category {
	query := "UPDATE category SET name = ? where id = ?"
	_, err := tx.ExecContext(ctx, query, category.Name, category.Id)
	error2.PanicIfError(err)

	return category
}

func (c CategoryRepositoryImpl) Delete(ctx context.Context, tx *sql.Tx, category *entity.Category) {
	query := "DELETE FROM category where id = ?"
	_, err := tx.ExecContext(ctx, query, category.Id)
	error2.PanicIfError(err)
}

func (c CategoryRepositoryImpl) FindById(ctx context.Context, tx *sql.Tx, id int) (*entity.Category, error) {
	query := "SELECT * FROM category WHERE id = ?"
	rows, err := tx.QueryContext(ctx, query, id)
	error2.PanicIfError(err)
	defer func(rows *sql.Rows) {
		err := rows.Close()
		error2.PanicIfError(err)
	}(rows)

	category := entity.Category{}
	if rows.Next() {
		err := rows.Scan(&category.Id, &category.Name)
		error2.PanicIfError(err)
		return &category, nil
	} else {
		return &category, errors.New("Category is not found")
	}
}

func (c CategoryRepositoryImpl) FindAll(ctx context.Context, tx *sql.Tx) []entity.Category {
	query := "SELECT * FROM category"
	rows, err := tx.QueryContext(ctx, query)
	error2.PanicIfError(err)
	defer func(rows *sql.Rows) {
		err := rows.Close()
		error2.PanicIfError(err)
	}(rows)

	var categories []entity.Category

	for rows.Next() {
		category := entity.Category{}
		err := rows.Scan(&category.Id, &category.Name)
		error2.PanicIfError(err)
		categories = append(categories, category)
	}

	return categories
}

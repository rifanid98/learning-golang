package repository

import (
	"context"
	"database/sql"
	"golang12_dependency_injection/restfulapi/domain/entity"
)

type CategoryRepository interface {
	Save(ctx context.Context, tx *sql.Tx, category *entity.Category) *entity.Category
	Update(ctx context.Context, tx *sql.Tx, category *entity.Category) *entity.Category
	Delete(ctx context.Context, tx *sql.Tx, category *entity.Category)
	FindById(ctx context.Context, tx *sql.Tx, id int) (*entity.Category, error)
	FindAll(ctx context.Context, tx *sql.Tx) []entity.Category
}

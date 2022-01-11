package repository

import (
	"context"
	"database/sql"
	"golang11_restful_api/domain/entity"
)

type CategoryRepository interface {
	Save(ctx context.Context, tx sql.Tx, category *entity.Category) (*entity.Category, error)
	Update(ctx context.Context, tx sql.Tx, category *entity.Category) (*entity.Category, error)
	Delete(ctx context.Context, tx sql.Tx, category *entity.Category) error
	FindById(ctx context.Context, tx sql.Tx, id int) (*entity.Category, error)
	FindAll(ctx context.Context, tx sql.Tx) ([]entity.Category, error)
}

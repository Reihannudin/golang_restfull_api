package repository

import (
	"context"
	"database/sql"
	"golang_rest_api_learn/model/entity"
)

type CategoriesRepository interface {
	Save(ctx context.Context, tx *sql.Tx, categories entity.Categories) entity.Categories
	Update(ctx context.Context, tx *sql.Tx, categories entity.Categories) entity.Categories
	Delete(ctx context.Context, tx *sql.Tx, categories entity.Categories)
	FindById(ctx context.Context, tx *sql.Tx, categoriesId int) (entity.Categories, error)
	FindAll(ctx context.Context, tx *sql.Tx) []entity.Categories
}

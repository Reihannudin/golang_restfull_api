package repository

import (
	"context"
	"database/sql"
	"errors"
	"github.com/go-playground/validator"
	"golang_rest_api_learn/helper"
	"golang_rest_api_learn/model/entity"
)

type CategoriesRepositoryImpl struct {
}

func NewCategoriesRespository() *CategoriesRepositoryImpl {
	return &CategoriesRepositoryImpl{}
}

func (repository *CategoriesRepositoryImpl) Save(ctx context.Context, tx *sql.Tx, categories entity.Categories) entity.Categories {
	SQL := "insert into categories(name) values (?)"
	result, err := tx.ExecContext(ctx, SQL, categories.Name)
	helper.PanicIfError(err)

	id, err := result.LastInsertId()
	helper.PanicIfError(err)

	categories.Id = int(id)
	return categories
}

func (repository *CategoriesRepositoryImpl) Update(ctx context.Context, tx *sql.Tx, categories entity.Categories) entity.Categories {
	SQL := "update categories set name = ? where id = ?"
	req, err := tx.ExecContext(ctx, SQL, categories.Name, categories.Id)
	validate := validator.New()
	validate.Struct(req)
	helper.PanicIfError(err)

	return categories
}

func (repository *CategoriesRepositoryImpl) Delete(ctx context.Context, tx *sql.Tx, categories entity.Categories) {
	SQL := "delete from categories where id = ?"
	_, err := tx.ExecContext(ctx, SQL, categories.Id)
	helper.PanicIfError(err)
}

func (repository *CategoriesRepositoryImpl) FindById(ctx context.Context, tx *sql.Tx, categoriesId int) (entity.Categories, error) {
	SQL := "select id, name from categories where id = ?"
	rows, err := tx.QueryContext(ctx, SQL, categoriesId)
	helper.PanicIfError(err)
	defer rows.Close()

	categories := entity.Categories{}
	if rows.Next() {
		err := rows.Scan(&categories.Id, &categories.Name)
		helper.PanicIfError(err)
		return categories, nil
	} else {
		return categories, errors.New("Categories is not found")
	}

}

func (repository *CategoriesRepositoryImpl) FindAll(ctx context.Context, tx *sql.Tx) []entity.Categories {
	SQL := "select id , name from categories"
	rows, err := tx.QueryContext(ctx, SQL)
	helper.PanicIfError(err)
	defer rows.Close()

	var categories []entity.Categories
	for rows.Next() {
		category := entity.Categories{}
		errScan := rows.Scan(&category.Id, &category.Name)
		helper.PanicIfError(errScan)
		categories = append(categories, category)
	}
	return categories
}

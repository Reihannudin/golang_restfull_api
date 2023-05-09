package service

import (
	"context"
	"database/sql"
	"github.com/go-playground/validator"
	"golang_rest_api_learn/helper"
	"golang_rest_api_learn/model/entity"
	"golang_rest_api_learn/model/exception"
	"golang_rest_api_learn/model/web"
	"golang_rest_api_learn/repository"
)

type CategoriesServiceImpl struct {
	CategoriesRepository repository.CategoriesRepository
	DB                   *sql.DB
	Validate             *validator.Validate
}

func NewCategoriesService(categoriesRepository repository.CategoriesRepository, DB *sql.DB, validate *validator.Validate) *CategoriesServiceImpl {
	return &CategoriesServiceImpl{
		CategoriesRepository: categoriesRepository,
		DB:                   DB,
		Validate:             validate,
	}
}

func (service *CategoriesServiceImpl) Create(ctx context.Context, request web.CategoriesCreateRequest) web.CategoriesResponse {

	err := service.Validate.Struct(request)
	helper.PanicIfError(err)

	tx, err := service.DB.Begin()
	helper.PanicIfError(err)

	defer helper.CommitOrRollback(tx)

	categories := entity.Categories{
		Name: request.Name,
	}

	category := service.CategoriesRepository.Save(ctx, tx, categories)
	return helper.ToCategoriesResponse(category)
}

func (service *CategoriesServiceImpl) Update(ctx context.Context, request web.CategoriesUpdateRequest) web.CategoriesResponse {

	err := service.Validate.Struct(request)
	helper.PanicIfError(err)

	tx, err := service.DB.Begin()
	helper.PanicIfError(err)
	defer helper.CommitOrRollback(tx)

	category, err := service.CategoriesRepository.FindById(ctx, tx, request.Id)
	//helper.PanicIfError(err)
	if err != nil {
		panic(exception.NewNotFound(err.Error()))
	}

	category.Name = request.Name

	category = service.CategoriesRepository.Update(ctx, tx, category)

	return helper.ToCategoriesResponse(category)

}

func (service *CategoriesServiceImpl) Delete(ctx context.Context, categoriesId int) {

	tx, err := service.DB.Begin()
	helper.PanicIfError(err)
	defer helper.CommitOrRollback(tx)

	category, err := service.CategoriesRepository.FindById(ctx, tx, categoriesId)
	if err != nil {
		panic(exception.NewNotFound(err.Error()))
	}

	service.CategoriesRepository.Delete(ctx, tx, category)

}

func (service *CategoriesServiceImpl) FindById(ctx context.Context, categoriesId int) web.CategoriesResponse {
	tx, err := service.DB.Begin()
	helper.PanicIfError(err)
	defer helper.CommitOrRollback(tx)

	category, err := service.CategoriesRepository.FindById(ctx, tx, categoriesId)
	if err != nil {
		panic(exception.NewNotFound(err.Error()))
	}

	return helper.ToCategoriesResponse(category)

}

func (service *CategoriesServiceImpl) FindAll(ctx context.Context) []web.CategoriesResponse {
	tx, err := service.DB.Begin()
	helper.PanicIfError(err)
	defer helper.CommitOrRollback(tx)

	categories := service.CategoriesRepository.FindAll(ctx, tx)

	return helper.ToCategoriesResponses(categories)
}

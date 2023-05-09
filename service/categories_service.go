package service

import (
	"context"
	"golang_rest_api_learn/model/web"
)

type CategoriesService interface {
	Create(ctx context.Context, request web.CategoriesCreateRequest) web.CategoriesResponse
	Update(ctx context.Context, request web.CategoriesUpdateRequest) web.CategoriesResponse
	Delete(ctx context.Context, categoriesId int)
	FindById(ctx context.Context, categoriesId int) web.CategoriesResponse
	FindAll(ctx context.Context) []web.CategoriesResponse
}

package helper

import (
	"golang_rest_api_learn/model/entity"
	"golang_rest_api_learn/model/web"
)

func ToCategoriesResponse(categories entity.Categories) web.CategoriesResponse {
	return web.CategoriesResponse{
		Id:   categories.Id,
		Name: categories.Name,
	}
}

func ToCategoriesResponses(categories []entity.Categories) []web.CategoriesResponse {
	var categoriesResponse []web.CategoriesResponse
	for _, category := range categories {
		categoriesResponse = append(categoriesResponse, ToCategoriesResponse(category))
	}
	return categoriesResponse
}

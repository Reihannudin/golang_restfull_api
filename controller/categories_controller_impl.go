package controller

import (
	"github.com/julienschmidt/httprouter"
	"golang_rest_api_learn/helper"
	"golang_rest_api_learn/model/web"
	"golang_rest_api_learn/service"
	"net/http"
	"strconv"
)

type CategoriesControllerImpl struct {
	CategoriesService service.CategoriesService
}

func NewCategoriesController(categoriesService service.CategoriesService) *CategoriesControllerImpl {
	return &CategoriesControllerImpl{
		CategoriesService: categoriesService,
	}
}

func (controller *CategoriesControllerImpl) Create(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
	categoriesCreateRequest := web.CategoriesCreateRequest{}
	helper.ReadFromRequestBody(request, &categoriesCreateRequest)

	categoriesResponse := controller.CategoriesService.Create(request.Context(), categoriesCreateRequest)
	webResponse := web.WebResponse{
		Code:   200,
		Status: "OK",
		Data:   categoriesResponse,
	}

	helper.WriteToResponseBody(writer, webResponse)

}

func (controller *CategoriesControllerImpl) Update(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {

	categoriesUpdateRequest := web.CategoriesUpdateRequest{}
	helper.ReadFromRequestBody(request, &categoriesUpdateRequest)

	categoryId := params.ByName("categoryId")
	id, err := strconv.Atoi(categoryId)
	helper.PanicIfError(err)

	categoriesUpdateRequest.Id = id

	categoriesResponse := controller.CategoriesService.Update(request.Context(), categoriesUpdateRequest)
	webResponse := web.WebResponse{
		Code:   200,
		Status: "OK",
		Data:   categoriesResponse,
	}

	helper.WriteToResponseBody(writer, webResponse)

}

func (controller *CategoriesControllerImpl) Delete(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
	categoryId := params.ByName("categoryId")
	id, err := strconv.Atoi(categoryId)
	helper.PanicIfError(err)

	controller.CategoriesService.Delete(request.Context(), id)
	webResponse := web.WebResponse{
		Code:   200,
		Status: "OK",
	}

	helper.WriteToResponseBody(writer, webResponse)

}

func (controller *CategoriesControllerImpl) FindById(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
	categoryId := params.ByName("categoryId")
	id, err := strconv.Atoi(categoryId)
	helper.PanicIfError(err)

	categoriesResponse := controller.CategoriesService.FindById(request.Context(), id)
	webResponse := web.WebResponse{
		Code:   200,
		Status: "OK",
		Data:   categoriesResponse,
	}

	helper.WriteToResponseBody(writer, webResponse)
}

func (controller *CategoriesControllerImpl) FindAll(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
	categoriesResponse := controller.CategoriesService.FindAll(request.Context())
	webResponse := web.WebResponse{
		Code:   200,
		Status: "OK",
		Data:   categoriesResponse,
	}

	helper.WriteToResponseBody(writer, webResponse)

}

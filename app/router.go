package app

import (
	"github.com/julienschmidt/httprouter"
	"golang_rest_api_learn/controller"
	"golang_rest_api_learn/model/exception"
)

func NewRouter(categoryController controller.CategoriesController) *httprouter.Router {
	router := httprouter.New()

	router.GET("/api/categories", categoryController.FindAll)
	router.GET("/api/categories/:categoryId", categoryController.FindById)
	router.POST("/api/categories", categoryController.Create)
	router.PUT("/api/categories/:categoryId", categoryController.Update)
	router.DELETE("/api/categories/:categoryId", categoryController.Delete)

	router.PanicHandler = exception.ErrorHandler

	return router
}

//go:build wireinject
// +build wireinject

package main

import (
	"github.com/go-playground/validator"
	"github.com/google/wire"
	"github.com/julienschmidt/httprouter"
	"golang_rest_api_learn/app"
	"golang_rest_api_learn/controller"
	"golang_rest_api_learn/middleware"
	"golang_rest_api_learn/repository"
	"golang_rest_api_learn/service"
	"net/http"
)

var categorySet = wire.NewSet(
	repository.NewCategoriesRespository,
	wire.Bind(new(repository.CategoriesRepository), new(*repository.CategoriesRepositoryImpl)),
	service.NewCategoriesService,
	wire.Bind(new(service.CategoriesService), new(*service.CategoriesServiceImpl)),
	controller.NewCategoriesController,
	wire.Bind(new(controller.CategoriesController), new(*controller.CategoriesControllerImpl)),
)

func InitializedServer() *http.Server {
	wire.Build(
		app.NewDB,
		validator.New,
		categorySet,
		app.NewRouter,
		wire.Bind(new(http.Handler), new(*httprouter.Router)),
		middleware.NewAuthMiddleware,
		NewServer,
	)
	return nil
}

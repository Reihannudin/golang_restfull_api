package middleware

import (
	"golang_rest_api_learn/helper"
	"golang_rest_api_learn/model/web"
	"net/http"
)

type AuthMiddleware struct {
	handler http.Handler
}

func NewAuthMiddleware(handler http.Handler) *AuthMiddleware {
	return &AuthMiddleware{
		handler: handler,
	}
}

func (middleware *AuthMiddleware) ServeHTTP(writer http.ResponseWriter, request *http.Request) {
	if "12345678" == request.Header.Get("X-API-AUTH-Key") {
		//	ok
		middleware.handler.ServeHTTP(writer, request)
	} else {
		//	error
		writer.Header().Set("Content-Type", "application/json")
		writer.WriteHeader(http.StatusUnauthorized)

		webResponse := web.WebResponse{
			Code:   http.StatusUnauthorized,
			Status: "UNAUTHORIZED",
		}

		helper.WriteToResponseBody(writer, webResponse)
	}
}

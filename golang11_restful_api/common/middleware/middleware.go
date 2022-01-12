package middleware

import (
	"golang11_restful_api/common/handler"
	"golang11_restful_api/common/response"
	"net/http"
)

type AuthMiddleware struct {
	http.Handler
}

func NewAuthMiddleware(handler http.Handler) *AuthMiddleware {
	return &AuthMiddleware{Handler: handler}
}

func (middleware *AuthMiddleware) ServeHTTP(writer http.ResponseWriter, request *http.Request) {
	if "secret" == request.Header.Get("X-API-Key") {
		middleware.Handler.ServeHTTP(writer, request)
	} else {
		writer.Header().Set("Content-Type", "application/json")
		writer.WriteHeader(http.StatusUnauthorized)

		res := response.PublicResponse{
			Code:   http.StatusUnauthorized,
			Status: "Unauthorized",
		}

		handler.SendResponseBody(writer, &res, http.StatusUnauthorized)
	}
}

package router

import (
	"github.com/julienschmidt/httprouter"
	"golang11_dependency_injection/adapter"
	"golang11_dependency_injection/common/exception"
)

func NewRouter(controller adapter.CategoryController) *httprouter.Router {
	router := httprouter.New()

	router.POST("/api/categories", controller.Create)
	router.GET("/api/categories", controller.FindAll)
	router.GET("/api/categories/:categoryId", controller.FindById)
	router.PUT("/api/categories/:categoryId", controller.Update)
	router.DELETE("/api/categories/:categoryId", controller.Delete)

	router.PanicHandler = exception.ErrorHandler

	return router
}

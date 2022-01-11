package main

import (
	"github.com/go-playground/validator"
	_ "github.com/go-sql-driver/mysql"
	"github.com/julienschmidt/httprouter"
	"golang11_restful_api/adapter"
	"golang11_restful_api/common"
	"golang11_restful_api/infrastructure/database"
	"golang11_restful_api/infrastructure/persistence/repository"
	"golang11_restful_api/usecase"
	"net/http"
)

func main() {
	validate := validator.New()
	db := database.NewDB()

	categoryRepository := repository.NewCategoryRepositoryImpl()
	categoryUsecase := usecase.NewCategoryInteractor(categoryRepository, db, validate)
	categoryController := adapter.NewCategoryHandler(categoryUsecase)

	router := httprouter.New()

	router.POST("/api/categories", categoryController.Create)
	router.GET("/api/categories", categoryController.FindAll)
	router.GET("/api/categories/:categoryId", categoryController.FindById)
	router.PUT("/api/categories/:categoryId", categoryController.Update)
	router.DELETE("/api/categories/:categoryId", categoryController.Delete)

	server := http.Server{
		Addr:    "localhost:3000",
		Handler: router,
	}

	err := server.ListenAndServe()
	common.PanicIfError(err)
}

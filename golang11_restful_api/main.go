package main

import (
	"github.com/go-playground/validator"
	_ "github.com/go-sql-driver/mysql"
	"golang11_restful_api/adapter"
	error2 "golang11_restful_api/common/error"
	"golang11_restful_api/common/middleware"
	"golang11_restful_api/infrastructure/database"
	"golang11_restful_api/infrastructure/persistence/repository"
	router2 "golang11_restful_api/infrastructure/router"
	"golang11_restful_api/usecase"
	"net/http"
)

func main() {
	validate := validator.New()
	db := database.NewDB()

	categoryRepository := repository.NewCategoryRepositoryImpl()
	categoryUsecase := usecase.NewCategoryInteractor(categoryRepository, db, validate)
	categoryController := adapter.NewCategoryHandler(categoryUsecase)

	router := router2.NewRouter(categoryController)

	server := http.Server{
		Addr:    "localhost:3000",
		Handler: middleware.NewAuthMiddleware(router),
	}

	err := server.ListenAndServe()
	error2.PanicIfError(err)
}

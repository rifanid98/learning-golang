package main

import (
	"github.com/go-playground/validator"
	_ "github.com/go-sql-driver/mysql"
	"golang12_dependency_injection/restfulapi/adapter"
	error2 "golang12_dependency_injection/restfulapi/common/error"
	"golang12_dependency_injection/restfulapi/common/middleware"
	"golang12_dependency_injection/restfulapi/infrastructure/database"
	"golang12_dependency_injection/restfulapi/infrastructure/persistence/repository"
	router2 "golang12_dependency_injection/restfulapi/infrastructure/router"
	"golang12_dependency_injection/restfulapi/usecase"
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

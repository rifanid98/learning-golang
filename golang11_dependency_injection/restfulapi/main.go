package main

import (
	"github.com/go-playground/validator"
	_ "github.com/go-sql-driver/mysql"
	"golang11_dependency_injection/adapter"
	error2 "golang11_dependency_injection/common/error"
	"golang11_dependency_injection/common/middleware"
	"golang11_dependency_injection/infrastructure/database"
	"golang11_dependency_injection/infrastructure/persistence/repository"
	router2 "golang11_dependency_injection/infrastructure/router"
	"golang11_dependency_injection/usecase"
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

//go:build wireinject
// +build wireinject

package di

import (
	"github.com/go-playground/validator"
	"github.com/google/wire"
	"github.com/julienschmidt/httprouter"
	"golang12_dependency_injection/restfulapi/adapter"
	"golang12_dependency_injection/restfulapi/common/middleware"
	repository2 "golang12_dependency_injection/restfulapi/domain/repository"
	"golang12_dependency_injection/restfulapi/infrastructure/database"
	"golang12_dependency_injection/restfulapi/infrastructure/persistence/repository"
	"golang12_dependency_injection/restfulapi/infrastructure/router"
	"golang12_dependency_injection/restfulapi/infrastructure/server"
	"golang12_dependency_injection/restfulapi/usecase"
	"net/http"
)

var categorySet = wire.NewSet(
	repository.NewCategoryRepositoryImpl,
	wire.Bind(new(repository2.CategoryRepository), new(*repository.CategoryRepositoryImpl)),
	usecase.NewCategoryInteractor,
	wire.Bind(new(usecase.CategoryUsecase), new(*usecase.CategoryInteractor)),
	adapter.NewCategoryHandler,
	wire.Bind(new(adapter.CategoryController), new(*adapter.CategoryHandler)),
)

func InitializedServer() *http.Server {
	wire.Build(
		database.NewDB,
		validator.New,
		categorySet,
		router.NewRouter,
		wire.Bind(new(http.Handler), new(*httprouter.Router)),
		middleware.NewAuthMiddleware,
		server.NewServer,
	)
	return nil
}

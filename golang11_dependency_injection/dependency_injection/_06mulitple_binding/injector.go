//go:build wireinject
// +build wireinject

package _06mulitple_binding

import "github.com/google/wire"

func InitializedDatabaseRepository() *DatabaseRepository {
	wire.Build(NewDatabasePostgreSQL, NewDatabaseMongoDB, NewDatabaseRepository)
	return nil
}

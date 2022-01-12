//go:build wireinject
// +build wireinject

package _02injector

import (
	"github.com/google/wire"
)

func InitializedService() *SimpleService {
	wire.Build(NewSimpleRepository, NewSimpleService)
	return nil
}

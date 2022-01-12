//go:build wireinject
// +build wireinject

package _03dependency_injection

import (
	"github.com/google/wire"
)

func InitializedService() (*SimpleService, error) {
	wire.Build(NewSimpleRepository, NewSimpleService)
	return nil, nil
}

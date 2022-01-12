//go:build wireinject
// +build wireinject

package _12cleanup_function

import (
	"github.com/google/wire"
)

func InitializedConnection(name string) (*Connection, func()) {
	wire.Build(NewConnection, NewFile)
	return nil, nil
}

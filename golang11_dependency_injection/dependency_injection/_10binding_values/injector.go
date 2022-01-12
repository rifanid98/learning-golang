//go:build wireinject
// +build wireinject

package _10binding_values

import (
	"github.com/google/wire"
	"io"
	"os"
)

var fooBarValueSet = wire.NewSet(
	wire.Value(&Foo{}),
	wire.Value(&Bar{}),
)

func InitializedFooBarUsingValue() *FooBar {
	wire.Build(fooBarValueSet, wire.Struct(new(FooBar), "*"))
	return nil
}

func InitializedReader() io.Reader {
	wire.Build(wire.InterfaceValue(new(io.Reader), os.Stdin))
	return nil
}

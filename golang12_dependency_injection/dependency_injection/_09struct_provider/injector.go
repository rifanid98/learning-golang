//go:build wireinject
// +build wireinject

package _09struct_provider

import "github.com/google/wire"

var fooBarSet = wire.NewSet(NewFoo, NewBar)

func InitializedFooBar() *FooBar {
	wire.Build(
		fooBarSet,
		wire.Struct(new(FooBar), "Foo", "Bar"), // * for all fields injection
	)
	return nil
}

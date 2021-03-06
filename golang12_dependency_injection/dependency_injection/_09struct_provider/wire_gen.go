// Code generated by Wire. DO NOT EDIT.

//go:generate go run github.com/google/wire/cmd/wire
//go:build !wireinject
// +build !wireinject

package _09struct_provider

import (
	"github.com/google/wire"
)

// Injectors from injector.go:

func InitializedFooBar() *FooBar {
	foo := NewFoo()
	bar := NewBar()
	fooBar := &FooBar{
		Foo: foo,
		Bar: bar,
	}
	return fooBar
}

// injector.go:

var fooBarSet = wire.NewSet(NewFoo, NewBar)

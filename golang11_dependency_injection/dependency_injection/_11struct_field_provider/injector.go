//go:build wireinject
// +build wireinject

package _11struct_field_provider

import (
	"github.com/google/wire"
)

func InitializeConfiguration() *Configuration {
	wire.Build(NewApplication, wire.FieldsOf(new(*Application), "Configuration"))
	return nil
}

// Code generated by Wire. DO NOT EDIT.

//go:generate go run github.com/google/wire/cmd/wire
//go:build !wireinject
// +build !wireinject

package _11struct_field_provider

// Injectors from injector.go:

func InitializeConfiguration() *Configuration {
	application := NewApplication()
	configuration := application.Configuration
	return configuration
}

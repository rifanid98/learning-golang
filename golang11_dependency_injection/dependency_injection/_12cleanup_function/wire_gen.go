// Code generated by Wire. DO NOT EDIT.

//go:generate go run github.com/google/wire/cmd/wire
//go:build !wireinject
// +build !wireinject

package _12cleanup_function

// Injectors from injector.go:

func InitializedConnection(name string) (*Connection, func()) {
	file, cleanup := NewFile(name)
	connection, cleanup2 := NewConnection(file)
	return connection, func() {
		cleanup2()
		cleanup()
	}
}

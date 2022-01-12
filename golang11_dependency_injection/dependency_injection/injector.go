//go:build wireinject
// +build wireinject

package dependency_injection

import (
	"github.com/google/wire"
	"io"
	"os"
)

/**
Injector

- Setelah kita membuat Provider untuk nanti kita gunakan, selanjutnya kita perlu membuat Injector
- Injector sendiri adalah sebuah function constructor, namun isinya berupa konfigurasi yang kita beritahukan ke Google
  Wire
- Injector ini sendiri sebenarnya tidak akan digunakan oleh kode program kita, Injector ini adalah function yang akan
  digunakan oleh Google Wire untuk melakukan auto generate kode Dependency Injection
- Khusus ketika membuat Injector, pada file nya kita perlu tambahkan komentar penanda :
*/

func InitializedService(isError bool) (*SimpleService, error) {
	wire.Build(NewSimpleRepository, NewSimpleService)
	return nil, nil
}

func InitializedDatabaseRepository() *DatabaseRepository {
	wire.Build(NewDatabasePostgreSQL, NewDatabaseMongoDB, NewDatabaseRepository)
	return nil
}

var fooSet = wire.NewSet(NewFooRepository, NewFooService)
var barSet = wire.NewSet(NewBarRepository, NewBarService)

func InitializedFooBarService() *FooBarService {
	wire.Build(fooSet, barSet, NewFooBarService)
	return nil
}

var helloSet = wire.NewSet(
	NewSayHelloImpl,
	wire.Bind(new(SayHello), new(*SayHelloImpl)),
)

func InitializedHelloService() *HelloService {
	wire.Build(helloSet, NewHelloService)
	return nil
}

var fooBarSet = wire.NewSet(NewFoo, NewBar)

func InitializedFooBar() *FooBar {
	wire.Build(
		fooBarSet,
		wire.Struct(new(FooBar), "Foo", "Bar"), // * for all fields injection
	)
	return nil
}

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

//go:build wireinject
// +build wireinject

package simple

import (
	"github.com/google/wire"
	"io"
	"os"
)

func InitializedService() *SimpleService {
	wire.Build(NewSimpleRepository, NewSimpleService)
	return nil
}

func InitializedError(isError bool) (*ErrorService, error) {
	wire.Build(NewErrorRepository, NewErrorService)
	return nil, nil
}

func InitializedMultiple() *DatabaseRepository {
	wire.Build(
		NewDatabaseMySQL,
		NewDatabaseMongoDB,
		NewDatabaseRepository,
	)
	return nil
}

var fooSet = wire.NewSet(NewFooRepository, NewFooService)
var barSet = wire.NewSet(NewBarRepository, NewBarService)

func InitializedProviderSet() *FooBarService {
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

//func InitializedWrongBindingInterface() *HelloService{
//	wire.Build(NewHelloService, NewSayHelloImpl())
//	return nil
//}

var FibonaciSet = wire.NewSet(
	NewFii,
	NewBoo,
	NewNaci,
)

func InitializedStructProvider() *Fibonaci {
	wire.Build(
		FibonaciSet,
		wire.Struct(new(Fibonaci), "*"),
	)
	return nil
}

var fooValue = &Foo{}
var barValue = &Bar{}

func InitializedFooBarUsingValue() *FooBar {
	wire.Build(wire.Value(fooValue), wire.Value(barValue), wire.Struct(new(FooBar), "*"))
	return nil
}

func InitializedReader() io.Reader {
	wire.Build(wire.InterfaceValue(new(io.Reader), os.Stdin))
	return nil
}

func InitializedConfiguration() *Configuration {
	wire.Build(NewApplication, wire.FieldsOf(new(*Application), "Configuration"))
	return nil
}

func InitializedConnection(name string) (*Connection, func()) {
	wire.Build(NewConnection, NewFile)
	return nil, nil
}

//go:build wireinject
// +build wireinject

package main

import (
	"github.com/google/wire"
	"go-clean-arch-ddd/account/infra/in/iris"
	"go-clean-arch-ddd/account/infra/out"
	"go-clean-arch-ddd/account/usecase/interface/in/getUser"
	"go-clean-arch-ddd/account/usecase/interface/in/register"
	"go-clean-arch-ddd/account/usecase/interface/in/rename"
	outinterface "go-clean-arch-ddd/account/usecase/interface/out"
	"go-clean-arch-ddd/account/usecase/service"
)

func newAccountUsecase() iris.WireParams {
	wire.Build(
		wire.NewSet(
			out.NewInMemoryUserRepository,
			wire.Bind(new(outinterface.UserRepository), new(*out.InMemoryUserRepository)),
		),
		wire.NewSet(
			service.NewUserRegisterService,
			wire.Bind(new(register.UseCase), new(*service.UserRegisterService)),
		),
		wire.NewSet(
			service.NewUserRenameService,
			wire.Bind(new(rename.UseCase), new(*service.UserRenameService)),
		),
		wire.NewSet(
			service.NewGetUserUseCase,
			wire.Bind(new(getuser.UseCase), new(*service.GetUserUseCase)),
		),
		wire.Struct(new(iris.WireParams), "*"),
	)
	return iris.WireParams{}
}

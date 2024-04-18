// Code generated by Wire. DO NOT EDIT.

//go:generate go run -mod=mod github.com/google/wire/cmd/wire
//go:build !wireinject
// +build !wireinject

package main

import (
	"go-clean-arch-ddd/account/infra/in/iris"
	"go-clean-arch-ddd/account/infra/out"
	"go-clean-arch-ddd/account/usecase/service"
)

// Injectors from wire.go:

func newAccountUsecase() iris.WireParams {
	inMemoryUserRepository := out.NewInMemoryUserRepository()
	userRegisterService := service.NewUserRegisterService(inMemoryUserRepository)
	userRenameService := service.NewUserRenameService(inMemoryUserRepository)
	getUserUseCase := service.NewGetUserUseCase(inMemoryUserRepository)
	wireParams := iris.WireParams{
		UserRegisterUseCase:       userRegisterService,
		UserChangePasswordUseCase: userRenameService,
		GetUserUseCase:            getUserUseCase,
	}
	return wireParams
}

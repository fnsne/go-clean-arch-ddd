package main

import (
	"github.com/kataras/iris/v12"
	irisInfra "go-clean-arch-ddd/account/infra/in/iris"
	"go-clean-arch-ddd/account/infra/out"
	"go-clean-arch-ddd/account/usecase/service"
)

/*
	We are using google wire to inject all dependencies
*/

func main() {
	application := iris.Default()

	userRepository := out.NewInMemoryUserRepository()
	userRegisterService := service.NewUserRegisterService(userRepository)

	//todo change to use DI framework to inject all use cases
	irisInfra.BindUseCases(application, userRegisterService)

	_ = application.Run(iris.Addr(":8080"))
}

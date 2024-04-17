package main

import (
	"context"
	"github.com/kataras/iris/v12"
	irisHandler "go-clean-arch-ddd/account/infra/in/iris"
	portout "go-clean-arch-ddd/account/infra/out"
	getuser "go-clean-arch-ddd/account/usecase/interface/in/getUser"
	"go-clean-arch-ddd/account/usecase/interface/in/register"
	chagnepassword "go-clean-arch-ddd/account/usecase/interface/in/rename"
	portin "go-clean-arch-ddd/account/usecase/interface/out"
	"go-clean-arch-ddd/account/usecase/service"
	"go.uber.org/fx"
)

/*
	We are using uber fx to inject all dependencies
*/

func main() {
	app := fx.New(
		fx.Provide(fx.Annotate(portout.NewInMemoryUserRepository, fx.As(new(portin.UserRepository)))),
		fx.Provide(fx.Annotate(service.NewUserRegisterService, fx.As(new(register.UseCase)))),
		fx.Provide(fx.Annotate(service.NewUserRenameService, fx.As(new(chagnepassword.UseCase)))),
		fx.Provide(fx.Annotate(service.NewGetUserUseCase, fx.As(new(getuser.UseCase)))),
		fx.Provide(iris.Default),
		fx.Invoke(newFxServer),
	)
	err := app.Start(context.Background())
	if err != nil {
		panic(err)
	}
}

func newFxServer(lifecycle fx.Lifecycle, app *iris.Application, params irisHandler.FxParams) {
	lifecycle.Append(fx.Hook{
		OnStart: func(context.Context) error {
			irisHandler.BindUseCasesWithFx(app, params)
			err := app.Run(iris.Addr(":8080"))
			return err
		},
	})
}

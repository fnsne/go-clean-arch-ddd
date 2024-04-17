package main

import (
	"context"
	"github.com/kataras/iris/v12"
	irisHandler "go-clean-arch-ddd/account/infra/in/iris"
	portout "go-clean-arch-ddd/account/infra/out"
	"go-clean-arch-ddd/account/usecase/interface/in/register"
	portin "go-clean-arch-ddd/account/usecase/interface/out"
	"go-clean-arch-ddd/account/usecase/service"
	"go.uber.org/fx"
)

/*
	We are using uber fx to inject all dependencies
*/

func main() {
	app := fx.New(
		fx.Provide(
			fx.Annotate(
				portout.NewInMemoryUserRepository,
				fx.As(new(portin.UserRepository)),
			),
		),
		fx.Provide(
			fx.Annotate(
				service.NewUserRegisterService,
				fx.As(new(register.UserRegisterUseCase)),
			),
		),
		fx.Provide(iris.Default),
		fx.Invoke(newFxServer),
	)
	err := app.Start(context.Background())
	if err != nil {
		panic(err)
	}
}

func newFxServer(lifecycle fx.Lifecycle, app *iris.Application, userRegisterUseCase register.UserRegisterUseCase) {
	lifecycle.Append(fx.Hook{
		OnStart: func(context.Context) error {
			irisHandler.BindUseCases(app, userRegisterUseCase)
			err := app.Run(iris.Addr(":8080"))
			return err
		},
	})
}

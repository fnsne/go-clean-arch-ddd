package iris

import (
	"github.com/kataras/iris/v12"
	chagnepassword "go-clean-arch-ddd/account/usecase/interface/in/changePassword"
	"go-clean-arch-ddd/account/usecase/interface/in/register"
	"go.uber.org/fx"
)

//below is for the use of uber fx DI
//================================================

type FxParams struct {
	fx.In
	UserRegisterUseCase       register.UseCase
	UserChangePasswordUseCase chagnepassword.UseCase
}

// BindUseCasesWithFx is a function to bind all use cases with uber fx
func BindUseCasesWithFx(app *iris.Application, p FxParams) {
	api := app.Party("/account")
	handlers := &Handlers{
		registerUseCase:       p.UserRegisterUseCase,
		changePasswordUseCase: p.UserChangePasswordUseCase,
	}

	api.Post("/register", handlers.UserRegister)
}

//================================================
//above is for the use of uber fx DI

type UserRegisterInput struct {
	Email    string
	Password string
	Name     string
}

type Handlers struct {
	registerUseCase       register.UseCase
	changePasswordUseCase chagnepassword.UseCase
}

func BindUseCases(app *iris.Application, registerUseCase register.UseCase, passwordUseCase chagnepassword.UseCase) {
	api := app.Party("/account")
	handlers := &Handlers{
		registerUseCase:       registerUseCase,
		changePasswordUseCase: passwordUseCase,
	}

	api.Post("/register", handlers.UserRegister)
	api.Post("/changePassword", handlers.ChangePassword)
}

func (h *Handlers) UserRegister(ctx iris.Context) {
	var input UserRegisterInput
	err := ctx.ReadJSON(&input)
	if err != nil {
		ctx.StatusCode(iris.StatusBadRequest)
		return
	}
	output, err := h.registerUseCase.Execute(register.Input{
		Email:    input.Email,
		Password: input.Password,
		Name:     input.Name,
	})
	if err != nil {
		ctx.Application().Logger().Errorf("[handler] failed to register user: %v", err)
		ctx.StatusCode(iris.StatusInternalServerError)
		return
	}
	_ = ctx.JSON(iris.Map{"id": output.ID})
}

func (h *Handlers) ChangePassword(ctx iris.Context) {
	var input chagnepassword.Input
	err := ctx.ReadJSON(&input)
	if err != nil {
		ctx.StatusCode(iris.StatusBadRequest)
		return
	}
	err = h.changePasswordUseCase.Execute(input)
	if err != nil {
		ctx.Application().Logger().Errorf("[handler] failed to change password: %v", err)
		ctx.StatusCode(iris.StatusInternalServerError)
		return
	}
	_ = ctx.JSON(iris.Map{"status": "success"})
}

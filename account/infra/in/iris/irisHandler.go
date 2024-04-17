package iris

import (
	"github.com/kataras/iris/v12"
	getuser "go-clean-arch-ddd/account/usecase/interface/in/getUser"
	"go-clean-arch-ddd/account/usecase/interface/in/register"
	"go-clean-arch-ddd/account/usecase/interface/in/rename"
	"go.uber.org/fx"
)

//below is for the use of uber fx DI
//================================================

type FxParams struct {
	fx.In
	UserRegisterUseCase       register.UseCase
	UserChangePasswordUseCase rename.UseCase
	GetUserUseCase            getuser.UseCase
}

// BindUseCasesWithFx is a function to bind all use cases with uber fx
func BindUseCasesWithFx(app *iris.Application, p FxParams) {
	BindUseCases(
		app,
		p.UserRegisterUseCase,
		p.UserChangePasswordUseCase,
		p.GetUserUseCase,
	)
}

//================================================
//above is for the use of uber fx DI

type UserRegisterInput struct {
	Email    string
	Password string
	Name     string
}

type Handlers struct {
	registerUseCase register.UseCase
	reanmeUseCase   rename.UseCase
	getUserUseCase  getuser.UseCase
}

func BindUseCases(
	app *iris.Application,
	registerUseCase register.UseCase,
	passwordUseCase rename.UseCase,
	getUserUseCase getuser.UseCase,
) {
	api := app.Party("/account")
	handlers := &Handlers{
		registerUseCase: registerUseCase,
		reanmeUseCase:   passwordUseCase,
		getUserUseCase:  getUserUseCase,
	}

	api.Post("/register", handlers.UserRegister)
	api.Post("/{id}/rename", handlers.Rename)
	api.Get("/{id}", handlers.GetUser)
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

func (h *Handlers) Rename(ctx iris.Context) {
	id := ctx.Params().Get("id")

	var input rename.Input
	err := ctx.ReadJSON(&input)
	if err != nil {
		ctx.StatusCode(iris.StatusBadRequest)
		return
	}
	input.ID = id
	err = h.reanmeUseCase.Execute(input)
	if err != nil {
		ctx.Application().Logger().Errorf("[handler] failed to change password: %v", err)
		ctx.StatusCode(iris.StatusInternalServerError)
		return
	}
	_ = ctx.JSON(iris.Map{"status": "success"})
}

func (h *Handlers) GetUser(ctx iris.Context) {
	output, err := h.getUserUseCase.Execute(getuser.Input{ID: ctx.Params().Get("id")})
	if err != nil {
		ctx.Application().Logger().Errorf("[handler] failed to get user: %v", err)
		ctx.StatusCode(iris.StatusInternalServerError)
		return
	}
	_ = ctx.JSON(iris.Map{
		"id":    output.ID,
		"email": output.Email,
		"name":  output.Name,
	})
}

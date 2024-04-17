package iris

import (
	"github.com/kataras/iris/v12"
	"go-clean-arch-ddd/account/usecase/interface/in/register"
)

type UserRegisterInput struct {
	Email    string
	Password string
	Name     string
}

type Handlers struct {
	userRegisterUseCase register.UserRegisterUseCase
}

func NewHandlers(userRegisterUseCase register.UserRegisterUseCase) *Handlers {
	return &Handlers{userRegisterUseCase: userRegisterUseCase}
}

func BindUseCases(app *iris.Application, userRegisterUseCase register.UserRegisterUseCase) {
	api := app.Party("/account")
	handlers := NewHandlers(userRegisterUseCase)

	api.Post("/register", handlers.UserRegister)
}

func (h *Handlers) UserRegister(ctx iris.Context) {
	var input UserRegisterInput
	err := ctx.ReadJSON(&input)
	if err != nil {
		ctx.StatusCode(iris.StatusBadRequest)
		return
	}
	output, err := h.userRegisterUseCase.Execute(register.UserRegisterInput{
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

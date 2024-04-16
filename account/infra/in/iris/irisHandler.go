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

func NewAccountHandlers(app *iris.Application, accountServicePrefix string, userRegisterUseCase register.UserRegisterUseCase) {
	api := app.Party(accountServicePrefix)
	api.Post("/register", func(ctx iris.Context) {
		var input UserRegisterInput
		err := ctx.ReadJSON(&input)
		if err != nil {
			ctx.StatusCode(iris.StatusBadRequest)
			return
		}
		output, err := userRegisterUseCase.Execute(register.UserRegisterInput{
			Email:    input.Email,
			Password: input.Password,
			Name:     input.Name,
		})
		if err != nil {
			ctx.StatusCode(iris.StatusInternalServerError)
			return
		}
		ctx.Application().Logger().Info("User registered with ID: " + output.ID + " success")
		_ = ctx.JSON(iris.Map{"id": output.ID})
	})
}

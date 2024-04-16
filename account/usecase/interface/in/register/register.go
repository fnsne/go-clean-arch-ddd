package register

type UserRegisterUseCase interface {
	Execute(input UserRegisterInput) (UserRegisterOutput, error)
}

type UserRegisterOutput struct {
	ID string
}

type UserRegisterInput struct {
	Email    string
	Password string
	Name     string
}

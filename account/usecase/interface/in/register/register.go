package register

type UseCase interface {
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

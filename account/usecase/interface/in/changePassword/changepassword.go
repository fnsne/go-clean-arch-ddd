package chagnepassword

type UseCase interface {
	Execute(input UserChangePasswordInput) error
}

type UserChangePasswordInput struct {
	ID       string
	Password string
}

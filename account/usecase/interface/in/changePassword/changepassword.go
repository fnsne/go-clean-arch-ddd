package chagnepassword

type UseCase interface {
	Execute(input Input) error
}

type Input struct {
	ID       string
	Password string
}

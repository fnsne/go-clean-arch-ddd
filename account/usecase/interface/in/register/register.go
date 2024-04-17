package register

type UseCase interface {
	Execute(input Input) (Output, error)
}

type Output struct {
	ID string
}

type Input struct {
	Email    string
	Password string
	Name     string
}

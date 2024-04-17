package getuser

type Output struct {
	ID    string
	Email string
	Name  string
}

type Input struct {
	ID string
}
type UseCase interface {
	Execute(input Input) (Output, error)
}

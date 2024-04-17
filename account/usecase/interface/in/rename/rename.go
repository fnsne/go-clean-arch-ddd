package rename

type UseCase interface {
	Execute(input Input) error
}

type Input struct {
	ID   string
	Name string
}

package chagnepassword

type Output struct {
	ID       string
	Email    string
	Name     string
	Password string // just for demonstration, should not be included in the output
}

type Input struct {
	ID string
}
type UseCase interface {
	Execute(input Input) (Output, error)
}

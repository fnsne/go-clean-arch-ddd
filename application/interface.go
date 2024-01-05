package application

type HashTool interface {
	Hash(password string) (string, error)
}
